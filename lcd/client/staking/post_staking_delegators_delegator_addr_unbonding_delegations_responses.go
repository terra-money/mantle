// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-project/mantle/lcd/models"
)

// PostStakingDelegatorsDelegatorAddrUnbondingDelegationsReader is a Reader for the PostStakingDelegatorsDelegatorAddrUnbondingDelegations structure.
type PostStakingDelegatorsDelegatorAddrUnbondingDelegationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK creates a PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK with default headers values
func NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK() *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK {
	return &PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK{}
}

/*PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK handles this case with default header values.

OK
*/
type PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK struct {
	Payload *models.BroadcastTxCommitResult
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/unbonding_delegations][%d] postStakingDelegatorsDelegatorAddrUnbondingDelegationsOK  %+v", 200, o.Payload)
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK) GetPayload() *models.BroadcastTxCommitResult {
	return o.Payload
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastTxCommitResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest creates a PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest with default headers values
func NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest() *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest {
	return &PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest{}
}

/*PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest handles this case with default header values.

Invalid delegator address or unbonding delegation request body
*/
type PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest struct {
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/unbonding_delegations][%d] postStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest ", 400)
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized creates a PostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized with default headers values
func NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized() *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized {
	return &PostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized{}
}

/*PostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized handles this case with default header values.

Key password is wrong
*/
type PostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized struct {
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/unbonding_delegations][%d] postStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized ", 401)
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError creates a PostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError with default headers values
func NewPostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError() *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError {
	return &PostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError{}
}

/*PostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError struct {
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/unbonding_delegations][%d] postStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError ", 500)
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody post staking delegators delegator addr unbonding delegations body
swagger:model PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody
*/
type PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`

	// delegator address
	DelegatorAddress models.Address `json:"delegator_address,omitempty"`

	// validator address
	ValidatorAddress models.ValidatorAddress `json:"validator_address,omitempty"`
}

// Validate validates this post staking delegators delegator addr unbonding delegations body
func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDelegatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValidatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody) validateDelegatorAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.DelegatorAddress) { // not required
		return nil
	}

	if err := o.DelegatorAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("delegation" + "." + "delegator_address")
		}
		return err
	}

	return nil
}

func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody) validateValidatorAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.ValidatorAddress) { // not required
		return nil
	}

	if err := o.ValidatorAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("delegation" + "." + "validator_address")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody) UnmarshalBinary(b []byte) error {
	var res PostStakingDelegatorsDelegatorAddrUnbondingDelegationsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
