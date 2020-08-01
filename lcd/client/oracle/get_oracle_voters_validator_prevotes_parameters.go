// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetOracleVotersValidatorPrevotesParams creates a new GetOracleVotersValidatorPrevotesParams object
// with the default values initialized.
func NewGetOracleVotersValidatorPrevotesParams() *GetOracleVotersValidatorPrevotesParams {
	var ()
	return &GetOracleVotersValidatorPrevotesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOracleVotersValidatorPrevotesParamsWithTimeout creates a new GetOracleVotersValidatorPrevotesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOracleVotersValidatorPrevotesParamsWithTimeout(timeout time.Duration) *GetOracleVotersValidatorPrevotesParams {
	var ()
	return &GetOracleVotersValidatorPrevotesParams{

		timeout: timeout,
	}
}

// NewGetOracleVotersValidatorPrevotesParamsWithContext creates a new GetOracleVotersValidatorPrevotesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOracleVotersValidatorPrevotesParamsWithContext(ctx context.Context) *GetOracleVotersValidatorPrevotesParams {
	var ()
	return &GetOracleVotersValidatorPrevotesParams{

		Context: ctx,
	}
}

// NewGetOracleVotersValidatorPrevotesParamsWithHTTPClient creates a new GetOracleVotersValidatorPrevotesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOracleVotersValidatorPrevotesParamsWithHTTPClient(client *http.Client) *GetOracleVotersValidatorPrevotesParams {
	var ()
	return &GetOracleVotersValidatorPrevotesParams{
		HTTPClient: client,
	}
}

/*GetOracleVotersValidatorPrevotesParams contains all the parameters to send to the API endpoint
for the get oracle voters validator prevotes operation typically these are written to a http.Request
*/
type GetOracleVotersValidatorPrevotesParams struct {

	/*Validator*/
	Validator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get oracle voters validator prevotes params
func (o *GetOracleVotersValidatorPrevotesParams) WithTimeout(timeout time.Duration) *GetOracleVotersValidatorPrevotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oracle voters validator prevotes params
func (o *GetOracleVotersValidatorPrevotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oracle voters validator prevotes params
func (o *GetOracleVotersValidatorPrevotesParams) WithContext(ctx context.Context) *GetOracleVotersValidatorPrevotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oracle voters validator prevotes params
func (o *GetOracleVotersValidatorPrevotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oracle voters validator prevotes params
func (o *GetOracleVotersValidatorPrevotesParams) WithHTTPClient(client *http.Client) *GetOracleVotersValidatorPrevotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oracle voters validator prevotes params
func (o *GetOracleVotersValidatorPrevotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidator adds the validator to the get oracle voters validator prevotes params
func (o *GetOracleVotersValidatorPrevotesParams) WithValidator(validator string) *GetOracleVotersValidatorPrevotesParams {
	o.SetValidator(validator)
	return o
}

// SetValidator adds the validator to the get oracle voters validator prevotes params
func (o *GetOracleVotersValidatorPrevotesParams) SetValidator(validator string) {
	o.Validator = validator
}

// WriteToRequest writes these params to a swagger request
func (o *GetOracleVotersValidatorPrevotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validator
	if err := r.SetPathParam("validator", o.Validator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
