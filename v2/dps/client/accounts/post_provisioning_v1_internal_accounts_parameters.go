// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewPostProvisioningV1InternalAccountsParams creates a new PostProvisioningV1InternalAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProvisioningV1InternalAccountsParams() *PostProvisioningV1InternalAccountsParams {
	return &PostProvisioningV1InternalAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProvisioningV1InternalAccountsParamsWithTimeout creates a new PostProvisioningV1InternalAccountsParams object
// with the ability to set a timeout on a request.
func NewPostProvisioningV1InternalAccountsParamsWithTimeout(timeout time.Duration) *PostProvisioningV1InternalAccountsParams {
	return &PostProvisioningV1InternalAccountsParams{
		timeout: timeout,
	}
}

// NewPostProvisioningV1InternalAccountsParamsWithContext creates a new PostProvisioningV1InternalAccountsParams object
// with the ability to set a context for a request.
func NewPostProvisioningV1InternalAccountsParamsWithContext(ctx context.Context) *PostProvisioningV1InternalAccountsParams {
	return &PostProvisioningV1InternalAccountsParams{
		Context: ctx,
	}
}

// NewPostProvisioningV1InternalAccountsParamsWithHTTPClient creates a new PostProvisioningV1InternalAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProvisioningV1InternalAccountsParamsWithHTTPClient(client *http.Client) *PostProvisioningV1InternalAccountsParams {
	return &PostProvisioningV1InternalAccountsParams{
		HTTPClient: client,
	}
}

/* PostProvisioningV1InternalAccountsParams contains all the parameters to send to the API endpoint
   for the post provisioning v1 internal accounts operation.

   Typically these are written to a http.Request.
*/
type PostProvisioningV1InternalAccountsParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.V1CreateAccountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post provisioning v1 internal accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProvisioningV1InternalAccountsParams) WithDefaults() *PostProvisioningV1InternalAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post provisioning v1 internal accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProvisioningV1InternalAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post provisioning v1 internal accounts params
func (o *PostProvisioningV1InternalAccountsParams) WithTimeout(timeout time.Duration) *PostProvisioningV1InternalAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post provisioning v1 internal accounts params
func (o *PostProvisioningV1InternalAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post provisioning v1 internal accounts params
func (o *PostProvisioningV1InternalAccountsParams) WithContext(ctx context.Context) *PostProvisioningV1InternalAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post provisioning v1 internal accounts params
func (o *PostProvisioningV1InternalAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post provisioning v1 internal accounts params
func (o *PostProvisioningV1InternalAccountsParams) WithHTTPClient(client *http.Client) *PostProvisioningV1InternalAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post provisioning v1 internal accounts params
func (o *PostProvisioningV1InternalAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post provisioning v1 internal accounts params
func (o *PostProvisioningV1InternalAccountsParams) WithRequest(request *models.V1CreateAccountRequest) *PostProvisioningV1InternalAccountsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post provisioning v1 internal accounts params
func (o *PostProvisioningV1InternalAccountsParams) SetRequest(request *models.V1CreateAccountRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostProvisioningV1InternalAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}