// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOracleDenomsActivesReader is a Reader for the GetOracleDenomsActives structure.
type GetOracleDenomsActivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleDenomsActivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOracleDenomsActivesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOracleDenomsActivesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOracleDenomsActivesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOracleDenomsActivesOK creates a GetOracleDenomsActivesOK with default headers values
func NewGetOracleDenomsActivesOK() *GetOracleDenomsActivesOK {
	return &GetOracleDenomsActivesOK{}
}

/*GetOracleDenomsActivesOK handles this case with default header values.

OK
*/
type GetOracleDenomsActivesOK struct {
	Payload []string
}

func (o *GetOracleDenomsActivesOK) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/actives][%d] getOracleDenomsActivesOK  %+v", 200, o.Payload)
}

func (o *GetOracleDenomsActivesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetOracleDenomsActivesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleDenomsActivesBadRequest creates a GetOracleDenomsActivesBadRequest with default headers values
func NewGetOracleDenomsActivesBadRequest() *GetOracleDenomsActivesBadRequest {
	return &GetOracleDenomsActivesBadRequest{}
}

/*GetOracleDenomsActivesBadRequest handles this case with default header values.

Bad Request
*/
type GetOracleDenomsActivesBadRequest struct {
}

func (o *GetOracleDenomsActivesBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/actives][%d] getOracleDenomsActivesBadRequest ", 400)
}

func (o *GetOracleDenomsActivesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOracleDenomsActivesInternalServerError creates a GetOracleDenomsActivesInternalServerError with default headers values
func NewGetOracleDenomsActivesInternalServerError() *GetOracleDenomsActivesInternalServerError {
	return &GetOracleDenomsActivesInternalServerError{}
}

/*GetOracleDenomsActivesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOracleDenomsActivesInternalServerError struct {
}

func (o *GetOracleDenomsActivesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/actives][%d] getOracleDenomsActivesInternalServerError ", 500)
}

func (o *GetOracleDenomsActivesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
