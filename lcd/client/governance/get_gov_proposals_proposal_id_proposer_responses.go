// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// GetGovProposalsProposalIDProposerReader is a Reader for the GetGovProposalsProposalIDProposer structure.
type GetGovProposalsProposalIDProposerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovProposalsProposalIDProposerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovProposalsProposalIDProposerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovProposalsProposalIDProposerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovProposalsProposalIDProposerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovProposalsProposalIDProposerOK creates a GetGovProposalsProposalIDProposerOK with default headers values
func NewGetGovProposalsProposalIDProposerOK() *GetGovProposalsProposalIDProposerOK {
	return &GetGovProposalsProposalIDProposerOK{}
}

/*GetGovProposalsProposalIDProposerOK handles this case with default header values.

OK
*/
type GetGovProposalsProposalIDProposerOK struct {
	Payload *models.Proposer
}

func (o *GetGovProposalsProposalIDProposerOK) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/proposer][%d] getGovProposalsProposalIdProposerOK  %+v", 200, o.Payload)
}

func (o *GetGovProposalsProposalIDProposerOK) GetPayload() *models.Proposer {
	return o.Payload
}

func (o *GetGovProposalsProposalIDProposerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Proposer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovProposalsProposalIDProposerBadRequest creates a GetGovProposalsProposalIDProposerBadRequest with default headers values
func NewGetGovProposalsProposalIDProposerBadRequest() *GetGovProposalsProposalIDProposerBadRequest {
	return &GetGovProposalsProposalIDProposerBadRequest{}
}

/*GetGovProposalsProposalIDProposerBadRequest handles this case with default header values.

Invalid proposal ID
*/
type GetGovProposalsProposalIDProposerBadRequest struct {
}

func (o *GetGovProposalsProposalIDProposerBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/proposer][%d] getGovProposalsProposalIdProposerBadRequest ", 400)
}

func (o *GetGovProposalsProposalIDProposerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDProposerInternalServerError creates a GetGovProposalsProposalIDProposerInternalServerError with default headers values
func NewGetGovProposalsProposalIDProposerInternalServerError() *GetGovProposalsProposalIDProposerInternalServerError {
	return &GetGovProposalsProposalIDProposerInternalServerError{}
}

/*GetGovProposalsProposalIDProposerInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetGovProposalsProposalIDProposerInternalServerError struct {
}

func (o *GetGovProposalsProposalIDProposerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/proposer][%d] getGovProposalsProposalIdProposerInternalServerError ", 500)
}

func (o *GetGovProposalsProposalIDProposerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
