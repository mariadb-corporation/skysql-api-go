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

// NewPostProvisioningV1ServicesServiceIDStorageSizeParams creates a new PostProvisioningV1ServicesServiceIDStorageSizeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProvisioningV1ServicesServiceIDStorageSizeParams() *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	return &PostProvisioningV1ServicesServiceIDStorageSizeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProvisioningV1ServicesServiceIDStorageSizeParamsWithTimeout creates a new PostProvisioningV1ServicesServiceIDStorageSizeParams object
// with the ability to set a timeout on a request.
func NewPostProvisioningV1ServicesServiceIDStorageSizeParamsWithTimeout(timeout time.Duration) *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	return &PostProvisioningV1ServicesServiceIDStorageSizeParams{
		timeout: timeout,
	}
}

// NewPostProvisioningV1ServicesServiceIDStorageSizeParamsWithContext creates a new PostProvisioningV1ServicesServiceIDStorageSizeParams object
// with the ability to set a context for a request.
func NewPostProvisioningV1ServicesServiceIDStorageSizeParamsWithContext(ctx context.Context) *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	return &PostProvisioningV1ServicesServiceIDStorageSizeParams{
		Context: ctx,
	}
}

// NewPostProvisioningV1ServicesServiceIDStorageSizeParamsWithHTTPClient creates a new PostProvisioningV1ServicesServiceIDStorageSizeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProvisioningV1ServicesServiceIDStorageSizeParamsWithHTTPClient(client *http.Client) *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	return &PostProvisioningV1ServicesServiceIDStorageSizeParams{
		HTTPClient: client,
	}
}

/* PostProvisioningV1ServicesServiceIDStorageSizeParams contains all the parameters to send to the API endpoint
   for the post provisioning v1 services service ID storage size operation.

   Typically these are written to a http.Request.
*/
type PostProvisioningV1ServicesServiceIDStorageSizeParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.V1SetStorageSizeRequest

	/* ServiceID.

	   Identifier for a services
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post provisioning v1 services service ID storage size params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) WithDefaults() *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post provisioning v1 services service ID storage size params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) WithTimeout(timeout time.Duration) *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) WithContext(ctx context.Context) *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) WithHTTPClient(client *http.Client) *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) WithRequest(request *models.V1SetStorageSizeRequest) *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) SetRequest(request *models.V1SetStorageSizeRequest) {
	o.Request = request
}

// WithServiceID adds the serviceID to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) WithServiceID(serviceID string) *PostProvisioningV1ServicesServiceIDStorageSizeParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the post provisioning v1 services service ID storage size params
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostProvisioningV1ServicesServiceIDStorageSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
