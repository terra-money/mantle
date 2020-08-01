// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// GetOracleVotersValidatorPrevotesReader is a Reader for the GetOracleVotersValidatorPrevotes structure.
type GetOracleVotersValidatorPrevotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleVotersValidatorPrevotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOracleVotersValidatorPrevotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOracleVotersValidatorPrevotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOracleVotersValidatorPrevotesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOracleVotersValidatorPrevotesOK creates a GetOracleVotersValidatorPrevotesOK with default headers values
func NewGetOracleVotersValidatorPrevotesOK() *GetOracleVotersValidatorPrevotesOK {
	return &GetOracleVotersValidatorPrevotesOK{}
}

/*GetOracleVotersValidatorPrevotesOK handles this case with default header values.

OK
*/
type GetOracleVotersValidatorPrevotesOK struct {
	Payload []*models.ExchangeRatePrevote
}

func (o *GetOracleVotersValidatorPrevotesOK) Error() string {
	return fmt.Sprintf("[GET /oracle/voters/{validator}/prevotes][%d] getOracleVotersValidatorPrevotesOK  %+v", 200, o.Payload)
}

func (o *GetOracleVotersValidatorPrevotesOK) GetPayload() []*models.ExchangeRatePrevote {
	return o.Payload
}

func (o *GetOracleVotersValidatorPrevotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleVotersValidatorPrevotesBadRequest creates a GetOracleVotersValidatorPrevotesBadRequest with default headers values
func NewGetOracleVotersValidatorPrevotesBadRequest() *GetOracleVotersValidatorPrevotesBadRequest {
	return &GetOracleVotersValidatorPrevotesBadRequest{}
}

/*GetOracleVotersValidatorPrevotesBadRequest handles this case with default header values.

Bad Request
*/
type GetOracleVotersValidatorPrevotesBadRequest struct {
}

func (o *GetOracleVotersValidatorPrevotesBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracle/voters/{validator}/prevotes][%d] getOracleVotersValidatorPrevotesBadRequest ", 400)
}

func (o *GetOracleVotersValidatorPrevotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOracleVotersValidatorPrevotesInternalServerError creates a GetOracleVotersValidatorPrevotesInternalServerError with default headers values
func NewGetOracleVotersValidatorPrevotesInternalServerError() *GetOracleVotersValidatorPrevotesInternalServerError {
	return &GetOracleVotersValidatorPrevotesInternalServerError{}
}

/*GetOracleVotersValidatorPrevotesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOracleVotersValidatorPrevotesInternalServerError struct {
}

func (o *GetOracleVotersValidatorPrevotesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oracle/voters/{validator}/prevotes][%d] getOracleVotersValidatorPrevotesInternalServerError ", 500)
}

func (o *GetOracleVotersValidatorPrevotesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
