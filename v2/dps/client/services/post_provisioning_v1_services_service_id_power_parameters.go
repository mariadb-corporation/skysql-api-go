// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewPostProvisioningV1ServicesServiceIDPowerParams creates a new PostProvisioningV1ServicesServiceIDPowerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProvisioningV1ServicesServiceIDPowerParams() *PostProvisioningV1ServicesServiceIDPowerParams {
	return &PostProvisioningV1ServicesServiceIDPowerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProvisioningV1ServicesServiceIDPowerParamsWithTimeout creates a new PostProvisioningV1ServicesServiceIDPowerParams object
// with the ability to set a timeout on a request.
func NewPostProvisioningV1ServicesServiceIDPowerParamsWithTimeout(timeout time.Duration) *PostProvisioningV1ServicesServiceIDPowerParams {
	return &PostProvisioningV1ServicesServiceIDPowerParams{
		timeout: timeout,
	}
}

// NewPostProvisioningV1ServicesServiceIDPowerParamsWithContext creates a new PostProvisioningV1ServicesServiceIDPowerParams object
// with the ability to set a context for a request.
func NewPostProvisioningV1ServicesServiceIDPowerParamsWithContext(ctx context.Context) *PostProvisioningV1ServicesServiceIDPowerParams {
	return &PostProvisioningV1ServicesServiceIDPowerParams{
		Context: ctx,
	}
}

// NewPostProvisioningV1ServicesServiceIDPowerParamsWithHTTPClient creates a new PostProvisioningV1ServicesServiceIDPowerParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProvisioningV1ServicesServiceIDPowerParamsWithHTTPClient(client *http.Client) *PostProvisioningV1ServicesServiceIDPowerParams {
	return &PostProvisioningV1ServicesServiceIDPowerParams{
		HTTPClient: client,
	}
}

/* PostProvisioningV1ServicesServiceIDPowerParams contains all the parameters to send to the API endpoint
   for the post provisioning v1 services service ID power operation.

   Typically these are written to a http.Request.
*/
type PostProvisioningV1ServicesServiceIDPowerParams struct {

	/* ServiceID.

	   Service ID
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post provisioning v1 services service ID power params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProvisioningV1ServicesServiceIDPowerParams) WithDefaults() *PostProvisioningV1ServicesServiceIDPowerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post provisioning v1 services service ID power params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProvisioningV1ServicesServiceIDPowerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post provisioning v1 services service ID power params
func (o *PostProvisioningV1ServicesServiceIDPowerParams) WithTimeout(timeout time.Duration) *PostProvisioningV1ServicesServiceIDPowerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post provisioning v1 services service ID power params
func (o *PostProvisioningV1ServicesServiceIDPowerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post provisioning v1 services service ID power params
func (o *PostProvisioningV1ServicesServiceIDPowerParams) WithContext(ctx context.Context) *PostProvisioningV1ServicesServiceIDPowerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post provisioning v1 services service ID power params
func (o *PostProvisioningV1ServicesServiceIDPowerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post provisioning v1 services service ID power params
func (o *PostProvisioningV1ServicesServiceIDPowerParams) WithHTTPClient(client *http.Client) *PostProvisioningV1ServicesServiceIDPowerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post provisioning v1 services service ID power params
func (o *PostProvisioningV1ServicesServiceIDPowerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the post provisioning v1 services service ID power params
func (o *PostProvisioningV1ServicesServiceIDPowerParams) WithServiceID(serviceID string) *PostProvisioningV1ServicesServiceIDPowerParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the post provisioning v1 services service ID power params
func (o *PostProvisioningV1ServicesServiceIDPowerParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostProvisioningV1ServicesServiceIDPowerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param service_id
	if err := r.SetPathParam("service_id", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
