package account_txs

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/terra-project/core/x/bank"
	"github.com/terra-project/core/x/gov"
	"github.com/terra-project/core/x/market"
	"github.com/terra-project/core/x/oracle"
	"github.com/terra-project/core/x/staking"
	"github.com/terra-project/core/x/wasm"
	"github.com/terra-project/mantle-official/indexers/tx_infos"
	"github.com/terra-project/mantle-official/indexers/txs"
	"github.com/terra-project/mantle-official/utils"
	"github.com/terra-project/mantle/types"
	"reflect"
)

type AccountTx struct {
	Account           string `model:"index"`
	MsgType           string `model:"index"`
	TxInfo            tx_infos.TxInfo
	Tx                txs.MantleTx
}

type AccountTxs []AccountTx

type request struct {
	Txs       txs.Txs
	TxInfos   tx_infos.TxInfos
	BaseState struct {
		Height int64
		Txs    []types.Tx
	}
}

func RegisterAccountTxs(register types.Register) {
	register(
		IndexAccountTx,
		reflect.TypeOf((*AccountTxs)(nil)),
	)
}

func IndexAccountTx(query types.Query, commit types.Commit) error {
	req := request{}
	if err := query(&req, nil); err != nil {
		return err
	}

	accountTxs := AccountTxs{}

	for txIndex, tx := range req.BaseState.Txs {
		txdoc, txdocErr := types.TxDecoder(tx)
		if txdocErr != nil {
			return txdocErr
		}
		if !req.TxInfos[txIndex].Success {
			continue
		}

		for _, msg := range txdoc.GetMsgs() {
			var relatedAddresses []string
			if relatedAddresses = getAddressFromMsg(msg); len(relatedAddresses) == 0 {
				continue
			}

			// handle cw20 transfers too
			if msgExecute, isMsgExecute := msg.(wasm.MsgExecuteContract); isMsgExecute {
				cw20Recipient, cw20Err := getCW20TransferRecipient(msgExecute)
				if cw20Err != nil {
					// unmarshal error cannot be handled
					// skip
					continue
				}
				if cw20Recipient != "" {
					relatedAddresses = append(relatedAddresses, cw20Recipient)
				}
			}

			for _, addr := range relatedAddresses {
				accountTxs = append(accountTxs, AccountTx{
					Account: addr,
					MsgType: utils.MsgRouteAndTypeToString(msg.Route(), msg.Type()),
					TxInfo:  req.TxInfos[txIndex],
					Tx:      req.Txs[txIndex],
				})
			}
		}
	}

	if commitErr := commit(accountTxs); commitErr != nil {
		return commitErr
	}

	return nil
}

func getAddressFromMsg(msg sdk.Msg) []string {
	switch m := msg.(type) {
	case bank.MsgSend:
		return []string{
			m.FromAddress.String(),
			m.ToAddress.String(),
		}

	case bank.MsgMultiSend:
		addrs := make([]string, 0)
		for _, input := range m.Inputs {
			addrs = append(addrs, input.Address.String())
		}

		for _, output := range m.Outputs {
			addrs = append(addrs, output.Address.String())
		}

		return addrs

	case staking.MsgDelegate:
		return []string{
			m.DelegatorAddress.String(),
		}
	case staking.MsgCreateValidator:
		return []string{
			m.DelegatorAddress.String(),
		}
	case staking.MsgBeginRedelegate:
		return []string{
			m.DelegatorAddress.String(),
		}
	case staking.MsgUndelegate:
		return []string{
			m.DelegatorAddress.String(),
		}
	case distribution.MsgSetWithdrawAddress:
		return []string{
			m.DelegatorAddress.String(),
		}
	case distribution.MsgWithdrawDelegatorReward:
		return []string{
			m.DelegatorAddress.String(),
		}

	case distribution.MsgWithdrawValidatorCommission:
		return []string{
			sdk.AccAddress(m.ValidatorAddress).String(),
		}

	case market.MsgSwap:
		return []string{
			m.Trader.String(),
		}

	case oracle.MsgExchangeRateVote:
		return []string{
			m.Feeder.String(),
		}
	case oracle.MsgExchangeRatePrevote:
		return []string{
			m.Feeder.String(),
		}

	case gov.MsgDeposit:
		return []string{
			m.Depositor.String(),
		}
	case gov.MsgVote:
		return []string{
			m.Voter.String(),
		}

	case gov.MsgSubmitProposal:
		return []string{
			m.Proposer.String(),
		}

	case wasm.MsgStoreCode:
		return []string{
			m.Sender.String(),
		}

	case wasm.MsgInstantiateContract:
		return []string{
			m.Owner.String(),
		}

	case wasm.MsgExecuteContract:
		return []string{
			m.Sender.String(),
			m.Contract.String(),
		}

	case wasm.MsgUpdateContractOwner:
		return []string{
			m.Owner.String(),
			m.Contract.String(),
		}

	case wasm.MsgMigrateContract:
		return []string{
			m.Owner.String(),
			m.Contract.String(),
		}

	default:
		return nil
	}

	return nil
}

func getCW20TransferRecipient(msg wasm.MsgExecuteContract) (string, error) {
	executeMsgString := msg.ExecuteMsg.Bytes()
	msgExecuteContract := new(struct {
		Transfer struct {
			Amount    string `json:"amount"`
			Recipient string `json:"recipient"`
		} `json:"transfer"`
	})

	if unmarshalErr := json.Unmarshal(executeMsgString, msgExecuteContract); unmarshalErr != nil {
		return "", unmarshalErr
	}

	var isValidCW20Transfer = !utils.IsZero(msgExecuteContract) && !utils.IsZero(msgExecuteContract.Transfer) && !utils.IsZero(msgExecuteContract.Transfer.Recipient) && !utils.IsZero(msgExecuteContract.Transfer.Amount)

	// skip if this is not cw20 transfer
	if !isValidCW20Transfer {
		return "", nil
	}

	return msgExecuteContract.Transfer.Recipient, nil
}
