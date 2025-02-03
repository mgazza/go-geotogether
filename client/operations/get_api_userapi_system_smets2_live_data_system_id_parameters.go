// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetAPIUserapiSystemSmets2LiveDataSystemIDParams creates a new GetAPIUserapiSystemSmets2LiveDataSystemIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIUserapiSystemSmets2LiveDataSystemIDParams() *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	return &GetAPIUserapiSystemSmets2LiveDataSystemIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIUserapiSystemSmets2LiveDataSystemIDParamsWithTimeout creates a new GetAPIUserapiSystemSmets2LiveDataSystemIDParams object
// with the ability to set a timeout on a request.
func NewGetAPIUserapiSystemSmets2LiveDataSystemIDParamsWithTimeout(timeout time.Duration) *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	return &GetAPIUserapiSystemSmets2LiveDataSystemIDParams{
		timeout: timeout,
	}
}

// NewGetAPIUserapiSystemSmets2LiveDataSystemIDParamsWithContext creates a new GetAPIUserapiSystemSmets2LiveDataSystemIDParams object
// with the ability to set a context for a request.
func NewGetAPIUserapiSystemSmets2LiveDataSystemIDParamsWithContext(ctx context.Context) *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	return &GetAPIUserapiSystemSmets2LiveDataSystemIDParams{
		Context: ctx,
	}
}

// NewGetAPIUserapiSystemSmets2LiveDataSystemIDParamsWithHTTPClient creates a new GetAPIUserapiSystemSmets2LiveDataSystemIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIUserapiSystemSmets2LiveDataSystemIDParamsWithHTTPClient(client *http.Client) *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	return &GetAPIUserapiSystemSmets2LiveDataSystemIDParams{
		HTTPClient: client,
	}
}

/*
GetAPIUserapiSystemSmets2LiveDataSystemIDParams contains all the parameters to send to the API endpoint

	for the get API userapi system smets2 live data system ID operation.

	Typically these are written to a http.Request.
*/
type GetAPIUserapiSystemSmets2LiveDataSystemIDParams struct {

	/* Authorization.

	   Bearer token received from login
	*/
	Authorization string

	/* SystemID.

	   Unique system identifier
	*/
	SystemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API userapi system smets2 live data system ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) WithDefaults() *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API userapi system smets2 live data system ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) WithTimeout(timeout time.Duration) *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) WithContext(ctx context.Context) *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) WithHTTPClient(client *http.Client) *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) WithAuthorization(authorization string) *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSystemID adds the systemID to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) WithSystemID(systemID string) *GetAPIUserapiSystemSmets2LiveDataSystemIDParams {
	o.SetSystemID(systemID)
	return o
}

// SetSystemID adds the systemId to the get API userapi system smets2 live data system ID params
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) SetSystemID(systemID string) {
	o.SystemID = systemID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIUserapiSystemSmets2LiveDataSystemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param systemId
	if err := r.SetPathParam("systemId", o.SystemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
