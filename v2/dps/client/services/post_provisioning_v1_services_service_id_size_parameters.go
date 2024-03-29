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

	"github.com/mariadb-corporation/skysql-api-go/v2/dps/models"
)

// NewPostProvisioningV1ServicesServiceIDSizeParams creates a new PostProvisioningV1ServicesServiceIDSizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProvisioningV1ServicesServiceIDSizeParams() *PostProvisioningV1ServicesServiceIDSizeParams {
	return &PostProvisioningV1ServicesServiceIDSizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProvisioningV1ServicesServiceIDSizeParamsWithTimeout creates a new PostProvisioningV1ServicesServiceIDSizeParams object
// with the ability to set a timeout on a request.
func NewPostProvisioningV1ServicesServiceIDSizeParamsWithTimeout(timeout time.Duration) *PostProvisioningV1ServicesServiceIDSizeParams {
	return &PostProvisioningV1ServicesServiceIDSizeParams{
		timeout: timeout,
	}
}

// NewPostProvisioningV1ServicesServiceIDSizeParamsWithContext creates a new PostProvisioningV1ServicesServiceIDSizeParams object
// with the ability to set a context for a request.
func NewPostProvisioningV1ServicesServiceIDSizeParamsWithContext(ctx context.Context) *PostProvisioningV1ServicesServiceIDSizeParams {
	return &PostProvisioningV1ServicesServiceIDSizeParams{
		Context: ctx,
	}
}

// NewPostProvisioningV1ServicesServiceIDSizeParamsWithHTTPClient creates a new PostProvisioningV1ServicesServiceIDSizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProvisioningV1ServicesServiceIDSizeParamsWithHTTPClient(client *http.Client) *PostProvisioningV1ServicesServiceIDSizeParams {
	return &PostProvisioningV1ServicesServiceIDSizeParams{
		HTTPClient: client,
	}
}

/* PostProvisioningV1ServicesServiceIDSizeParams contains all the parameters to send to the API endpoint
   for the post provisioning v1 services service ID size operation.

   Typically these are written to a http.Request.
*/
type PostProvisioningV1ServicesServiceIDSizeParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.V1ServiceSizeState

	/* ServiceID.

	   Identifier for a services
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post provisioning v1 services service ID size params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProvisioningV1ServicesServiceIDSizeParams) WithDefaults() *PostProvisioningV1ServicesServiceIDSizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post provisioning v1 services service ID size params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProvisioningV1ServicesServiceIDSizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) WithTimeout(timeout time.Duration) *PostProvisioningV1ServicesServiceIDSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) WithContext(ctx context.Context) *PostProvisioningV1ServicesServiceIDSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) WithHTTPClient(client *http.Client) *PostProvisioningV1ServicesServiceIDSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) WithRequest(request *models.V1ServiceSizeState) *PostProvisioningV1ServicesServiceIDSizeParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) SetRequest(request *models.V1ServiceSizeState) {
	o.Request = request
}

// WithServiceID adds the serviceID to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) WithServiceID(serviceID string) *PostProvisioningV1ServicesServiceIDSizeParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the post provisioning v1 services service ID size params
func (o *PostProvisioningV1ServicesServiceIDSizeParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostProvisioningV1ServicesServiceIDSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param service_id
	if err := r.SetPathParam("service_id", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
