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

// GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrReader is a Reader for the GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddr structure.
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK creates a GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK() *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK {
	return &GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK{}
}

/*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK handles this case with default header values.

OK
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK struct {
	Payload *models.Delegation
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK  %+v", 200, o.Payload)
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK) GetPayload() *models.Delegation {
	return o.Payload
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Delegation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest creates a GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest() *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest {
	return &GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest{}
}

/*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest handles this case with default header values.

Invalid delegator address or validator address
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest struct {
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest ", 400)
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError creates a GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError with default headers values
func NewGetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError() *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError {
	return &GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError{}
}

/*GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError struct {
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/delegations/{validatorAddr}][%d] getStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError ", 500)
}

func (o *GetStakingDelegatorsDelegatorAddrDelegationsValidatorAddrInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
