// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrReader is a Reader for the GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddr structure.
type GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK creates a GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK with default headers values
func NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK() *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK {
	return &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK{}
}

/*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK handles this case with default header values.

OK
*/
type GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK struct {
	Payload *models.UnbondingDelegation
}

func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/unbonding_delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK  %+v", 200, o.Payload)
}

func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK) GetPayload() *models.UnbondingDelegation {
	return o.Payload
}

func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnbondingDelegation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest creates a GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest with default headers values
func NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest() *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest {
	return &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest{}
}

/*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest handles this case with default header values.

Invalid delegator address or validator address
*/
type GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest struct {
}

func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/unbonding_delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest ", 400)
}

func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError creates a GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError with default headers values
func NewGetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError() *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError {
	return &GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError{}
}

/*GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError struct {
}

func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/unbonding_delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError ", 500)
}

func (o *GetStakingDelegatorsDelegatorAddrUnbondingDelegationsValidatorAddrInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
