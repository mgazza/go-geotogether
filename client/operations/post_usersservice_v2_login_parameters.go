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

// NewPostUsersserviceV2LoginParams creates a new PostUsersserviceV2LoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUsersserviceV2LoginParams() *PostUsersserviceV2LoginParams {
	return &PostUsersserviceV2LoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsersserviceV2LoginParamsWithTimeout creates a new PostUsersserviceV2LoginParams object
// with the ability to set a timeout on a request.
func NewPostUsersserviceV2LoginParamsWithTimeout(timeout time.Duration) *PostUsersserviceV2LoginParams {
	return &PostUsersserviceV2LoginParams{
		timeout: timeout,
	}
}

// NewPostUsersserviceV2LoginParamsWithContext creates a new PostUsersserviceV2LoginParams object
// with the ability to set a context for a request.
func NewPostUsersserviceV2LoginParamsWithContext(ctx context.Context) *PostUsersserviceV2LoginParams {
	return &PostUsersserviceV2LoginParams{
		Context: ctx,
	}
}

// NewPostUsersserviceV2LoginParamsWithHTTPClient creates a new PostUsersserviceV2LoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUsersserviceV2LoginParamsWithHTTPClient(client *http.Client) *PostUsersserviceV2LoginParams {
	return &PostUsersserviceV2LoginParams{
		HTTPClient: client,
	}
}

/*
PostUsersserviceV2LoginParams contains all the parameters to send to the API endpoint

	for the post usersservice v2 login operation.

	Typically these are written to a http.Request.
*/
type PostUsersserviceV2LoginParams struct {

	// Body.
	Body PostUsersserviceV2LoginBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post usersservice v2 login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUsersserviceV2LoginParams) WithDefaults() *PostUsersserviceV2LoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post usersservice v2 login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUsersserviceV2LoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post usersservice v2 login params
func (o *PostUsersserviceV2LoginParams) WithTimeout(timeout time.Duration) *PostUsersserviceV2LoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post usersservice v2 login params
func (o *PostUsersserviceV2LoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post usersservice v2 login params
func (o *PostUsersserviceV2LoginParams) WithContext(ctx context.Context) *PostUsersserviceV2LoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post usersservice v2 login params
func (o *PostUsersserviceV2LoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post usersservice v2 login params
func (o *PostUsersserviceV2LoginParams) WithHTTPClient(client *http.Client) *PostUsersserviceV2LoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post usersservice v2 login params
func (o *PostUsersserviceV2LoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post usersservice v2 login params
func (o *PostUsersserviceV2LoginParams) WithBody(body PostUsersserviceV2LoginBody) *PostUsersserviceV2LoginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post usersservice v2 login params
func (o *PostUsersserviceV2LoginParams) SetBody(body PostUsersserviceV2LoginBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsersserviceV2LoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
