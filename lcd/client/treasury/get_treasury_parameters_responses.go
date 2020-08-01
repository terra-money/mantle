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

// GetTreasuryParametersReader is a Reader for the GetTreasuryParameters structure.
type GetTreasuryParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTreasuryParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTreasuryParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTreasuryParametersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTreasuryParametersOK creates a GetTreasuryParametersOK with default headers values
func NewGetTreasuryParametersOK() *GetTreasuryParametersOK {
	return &GetTreasuryParametersOK{}
}

/*GetTreasuryParametersOK handles this case with default header values.

OK
*/
type GetTreasuryParametersOK struct {
	Payload *models.TreasuryParams
}

func (o *GetTreasuryParametersOK) Error() string {
	return fmt.Sprintf("[GET /treasury/parameters][%d] getTreasuryParametersOK  %+v", 200, o.Payload)
}

func (o *GetTreasuryParametersOK) GetPayload() *models.TreasuryParams {
	return o.Payload
}

func (o *GetTreasuryParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TreasuryParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTreasuryParametersNotFound creates a GetTreasuryParametersNotFound with default headers values
func NewGetTreasuryParametersNotFound() *GetTreasuryParametersNotFound {
	return &GetTreasuryParametersNotFound{}
}

/*GetTreasuryParametersNotFound handles this case with default header values.

Not Found
*/
type GetTreasuryParametersNotFound struct {
}

func (o *GetTreasuryParametersNotFound) Error() string {
	return fmt.Sprintf("[GET /treasury/parameters][%d] getTreasuryParametersNotFound ", 404)
}

func (o *GetTreasuryParametersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
