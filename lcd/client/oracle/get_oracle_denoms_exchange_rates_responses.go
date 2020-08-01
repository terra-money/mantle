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

// GetOracleDenomsExchangeRatesReader is a Reader for the GetOracleDenomsExchangeRates structure.
type GetOracleDenomsExchangeRatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleDenomsExchangeRatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOracleDenomsExchangeRatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOracleDenomsExchangeRatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOracleDenomsExchangeRatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOracleDenomsExchangeRatesOK creates a GetOracleDenomsExchangeRatesOK with default headers values
func NewGetOracleDenomsExchangeRatesOK() *GetOracleDenomsExchangeRatesOK {
	return &GetOracleDenomsExchangeRatesOK{}
}

/*GetOracleDenomsExchangeRatesOK handles this case with default header values.

OK
*/
type GetOracleDenomsExchangeRatesOK struct {
	Payload []*models.DecCoin
}

func (o *GetOracleDenomsExchangeRatesOK) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/exchange_rates][%d] getOracleDenomsExchangeRatesOK  %+v", 200, o.Payload)
}

func (o *GetOracleDenomsExchangeRatesOK) GetPayload() []*models.DecCoin {
	return o.Payload
}

func (o *GetOracleDenomsExchangeRatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleDenomsExchangeRatesBadRequest creates a GetOracleDenomsExchangeRatesBadRequest with default headers values
func NewGetOracleDenomsExchangeRatesBadRequest() *GetOracleDenomsExchangeRatesBadRequest {
	return &GetOracleDenomsExchangeRatesBadRequest{}
}

/*GetOracleDenomsExchangeRatesBadRequest handles this case with default header values.

Bad Request
*/
type GetOracleDenomsExchangeRatesBadRequest struct {
}

func (o *GetOracleDenomsExchangeRatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/exchange_rates][%d] getOracleDenomsExchangeRatesBadRequest ", 400)
}

func (o *GetOracleDenomsExchangeRatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOracleDenomsExchangeRatesInternalServerError creates a GetOracleDenomsExchangeRatesInternalServerError with default headers values
func NewGetOracleDenomsExchangeRatesInternalServerError() *GetOracleDenomsExchangeRatesInternalServerError {
	return &GetOracleDenomsExchangeRatesInternalServerError{}
}

/*GetOracleDenomsExchangeRatesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOracleDenomsExchangeRatesInternalServerError struct {
}

func (o *GetOracleDenomsExchangeRatesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/exchange_rates][%d] getOracleDenomsExchangeRatesInternalServerError ", 500)
}

func (o *GetOracleDenomsExchangeRatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
