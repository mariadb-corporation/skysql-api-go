// Code generated by go-swagger; DO NOT EDIT.

package config

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
	"github.com/go-openapi/swag"
)

// NewGetProvisioningV1ConfigsParams creates a new GetProvisioningV1ConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProvisioningV1ConfigsParams() *GetProvisioningV1ConfigsParams {
	return &GetProvisioningV1ConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProvisioningV1ConfigsParamsWithTimeout creates a new GetProvisioningV1ConfigsParams object
// with the ability to set a timeout on a request.
func NewGetProvisioningV1ConfigsParamsWithTimeout(timeout time.Duration) *GetProvisioningV1ConfigsParams {
	return &GetProvisioningV1ConfigsParams{
		timeout: timeout,
	}
}

// NewGetProvisioningV1ConfigsParamsWithContext creates a new GetProvisioningV1ConfigsParams object
// with the ability to set a context for a request.
func NewGetProvisioningV1ConfigsParamsWithContext(ctx context.Context) *GetProvisioningV1ConfigsParams {
	return &GetProvisioningV1ConfigsParams{
		Context: ctx,
	}
}

// NewGetProvisioningV1ConfigsParamsWithHTTPClient creates a new GetProvisioningV1ConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProvisioningV1ConfigsParamsWithHTTPClient(client *http.Client) *GetProvisioningV1ConfigsParams {
	return &GetProvisioningV1ConfigsParams{
		HTTPClient: client,
	}
}

/* GetProvisioningV1ConfigsParams contains all the parameters to send to the API endpoint
   for the get provisioning v1 configs operation.

   Typically these are written to a http.Request.
*/
type GetProvisioningV1ConfigsParams struct {

	/* PageOrder.

	   page_order
	*/
	PageOrder *string

	/* PageSize.

	   page_size

	   Default: 10
	*/
	PageSize *int64

	/* PageToken.

	   page_token
	*/
	PageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get provisioning v1 configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProvisioningV1ConfigsParams) WithDefaults() *GetProvisioningV1ConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get provisioning v1 configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProvisioningV1ConfigsParams) SetDefaults() {
	var (
		pageSizeDefault = int64(10)
	)

	val := GetProvisioningV1ConfigsParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) WithTimeout(timeout time.Duration) *GetProvisioningV1ConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) WithContext(ctx context.Context) *GetProvisioningV1ConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) WithHTTPClient(client *http.Client) *GetProvisioningV1ConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageOrder adds the pageOrder to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) WithPageOrder(pageOrder *string) *GetProvisioningV1ConfigsParams {
	o.SetPageOrder(pageOrder)
	return o
}

// SetPageOrder adds the pageOrder to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) SetPageOrder(pageOrder *string) {
	o.PageOrder = pageOrder
}

// WithPageSize adds the pageSize to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) WithPageSize(pageSize *int64) *GetProvisioningV1ConfigsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) WithPageToken(pageToken *string) *GetProvisioningV1ConfigsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the get provisioning v1 configs params
func (o *GetProvisioningV1ConfigsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetProvisioningV1ConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageOrder != nil {

		// query param page_order
		var qrPageOrder string

		if o.PageOrder != nil {
			qrPageOrder = *o.PageOrder
		}
		qPageOrder := qrPageOrder
		if qPageOrder != "" {

			if err := r.SetQueryParam("page_order", qPageOrder); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param page_token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("page_token", qPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
