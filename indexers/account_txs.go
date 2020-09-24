package indexers

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/terra-project/core/x/bank"
	"github.com/terra-project/core/x/gov"
	"github.com/terra-project/core/x/market"
	"github.com/terra-project/core/x/oracle"
	"github.com/terra-project/core/x/staking"
	"github.com/terra-project/core/x/wasm"
	"github.com/terra-project/mantle/types"
	"reflect"
)

func RegisterAccountTxs(register types.Register) {
	register(
		IndexAccountTx,
		reflect.TypeOf((*AccountTxs)(nil)),
	)
}

type AccountTx struct {
	Account   string `model:"index"`
	MsgType   string `model:"index"`
	Tx        mantleTx
}

type AccountTxs []AccountTx

type request struct {
	Txs       Txs
	BaseState struct {
		Height int64
		Txs    []types.LazyTx
	}
}

func IndexAccountTx(query types.Query, commit types.Commit) error {
	req := request{}
	if err := query(&req, nil); err != nil {
		return err
	}

	accountTxs := AccountTxs{}

	for txIndex, tx := range req.BaseState.Txs {
		txdoc := tx.Decode()

		for _, msg := range txdoc.GetMsgs() {
			var relatedAddresses []string
			if relatedAddresses = getAddressFromMsg(msg); len(relatedAddresses) == 0 {
				continue
			}

			for _, addr := range relatedAddresses {
				accountTxs = append(accountTxs, AccountTx{
					Account:   addr,
					MsgType:   fmt.Sprintf("%s/%s", msg.Route(), msg.Type()),
					Tx:        req.Txs[txIndex],
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
