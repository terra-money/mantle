// Code generated by go-swagger; DO NOT EDIT.

package staking

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

// NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams creates a new GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams object
// with the default values initialized.
func NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams() *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	var ()
	return &GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParamsWithTimeout creates a new GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParamsWithTimeout(timeout time.Duration) *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	var ()
	return &GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams{

		timeout: timeout,
	}
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParamsWithContext creates a new GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParamsWithContext(ctx context.Context) *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	var ()
	return &GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams{

		Context: ctx,
	}
}

// NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParamsWithHTTPClient creates a new GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParamsWithHTTPClient(client *http.Client) *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	var ()
	return &GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams{
		HTTPClient: client,
	}
}

/*GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams contains all the parameters to send to the API endpoint
for the get staking delegators delegator addr validators validator addr operation typically these are written to a http.Request
*/
type GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams struct {

	/*DelegatorAddr
	  Bech32 AccAddress of Delegator

	*/
	DelegatorAddr string
	/*ValidatorAddr
	  Bech32 ValAddress of Delegator

	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) WithTimeout(timeout time.Duration) *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) WithContext(ctx context.Context) *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) WithHTTPClient(client *http.Client) *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) WithDelegatorAddr(delegatorAddr string) *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WithValidatorAddr adds the validatorAddr to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) WithValidatorAddr(validatorAddr string) *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get staking delegators delegator addr validators validator addr params
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetStakingDelegatorsDelegatorAddrValidatorsValidatorAddrParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegatorAddr
	if err := r.SetPathParam("delegatorAddr", o.DelegatorAddr); err != nil {
		return err
	}

	// path param validatorAddr
	if err := r.SetPathParam("validatorAddr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
