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

// NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams creates a new DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams() *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	return &DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistParamsWithTimeout creates a new DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams object
// with the ability to set a timeout on a request.
func NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistParamsWithTimeout(timeout time.Duration) *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	return &DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams{
		timeout: timeout,
	}
}

// NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistParamsWithContext creates a new DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams object
// with the ability to set a context for a request.
func NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistParamsWithContext(ctx context.Context) *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	return &DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams{
		Context: ctx,
	}
}

// NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistParamsWithHTTPClient creates a new DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistParamsWithHTTPClient(client *http.Client) *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	return &DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams{
		HTTPClient: client,
	}
}

/* DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams contains all the parameters to send to the API endpoint
   for the delete provisioning v1 services service ID security allowlist operation.

   Typically these are written to a http.Request.
*/
type DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.V1DeleteAllowedAddressRequest

	/* ServiceID.

	   Service ID
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete provisioning v1 services service ID security allowlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) WithDefaults() *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete provisioning v1 services service ID security allowlist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) WithTimeout(timeout time.Duration) *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) WithContext(ctx context.Context) *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) WithHTTPClient(client *http.Client) *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) WithRequest(request *models.V1DeleteAllowedAddressRequest) *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) SetRequest(request *models.V1DeleteAllowedAddressRequest) {
	o.Request = request
}

// WithServiceID adds the serviceID to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) WithServiceID(serviceID string) *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the delete provisioning v1 services service ID security allowlist params
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
