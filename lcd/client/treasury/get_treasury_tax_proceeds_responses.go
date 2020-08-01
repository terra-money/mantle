// Code generated by go-swagger; DO NOT EDIT.

package treasury

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// GetTreasuryTaxProceedsReader is a Reader for the GetTreasuryTaxProceeds structure.
type GetTreasuryTaxProceedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTreasuryTaxProceedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTreasuryTaxProceedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetTreasuryTaxProceedsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTreasuryTaxProceedsOK creates a GetTreasuryTaxProceedsOK with default headers values
func NewGetTreasuryTaxProceedsOK() *GetTreasuryTaxProceedsOK {
	return &GetTreasuryTaxProceedsOK{}
}

/*GetTreasuryTaxProceedsOK handles this case with default header values.

OK
*/
type GetTreasuryTaxProceedsOK struct {
	Payload []*models.Coin
}

func (o *GetTreasuryTaxProceedsOK) Error() string {
	return fmt.Sprintf("[GET /treasury/tax_proceeds][%d] getTreasuryTaxProceedsOK  %+v", 200, o.Payload)
}

func (o *GetTreasuryTaxProceedsOK) GetPayload() []*models.Coin {
	return o.Payload
}

func (o *GetTreasuryTaxProceedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTreasuryTaxProceedsInternalServerError creates a GetTreasuryTaxProceedsInternalServerError with default headers values
func NewGetTreasuryTaxProceedsInternalServerError() *GetTreasuryTaxProceedsInternalServerError {
	return &GetTreasuryTaxProceedsInternalServerError{}
}

/*GetTreasuryTaxProceedsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetTreasuryTaxProceedsInternalServerError struct {
}

func (o *GetTreasuryTaxProceedsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /treasury/tax_proceeds][%d] getTreasuryTaxProceedsInternalServerError ", 500)
}

func (o *GetTreasuryTaxProceedsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
