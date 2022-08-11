// Code generated by go-swagger; DO NOT EDIT.

package delete_preference

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

// NewDeleteNotificationsV1PreferencesPreferenceIDParams creates a new DeleteNotificationsV1PreferencesPreferenceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNotificationsV1PreferencesPreferenceIDParams() *DeleteNotificationsV1PreferencesPreferenceIDParams {
	return &DeleteNotificationsV1PreferencesPreferenceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNotificationsV1PreferencesPreferenceIDParamsWithTimeout creates a new DeleteNotificationsV1PreferencesPreferenceIDParams object
// with the ability to set a timeout on a request.
func NewDeleteNotificationsV1PreferencesPreferenceIDParamsWithTimeout(timeout time.Duration) *DeleteNotificationsV1PreferencesPreferenceIDParams {
	return &DeleteNotificationsV1PreferencesPreferenceIDParams{
		timeout: timeout,
	}
}

// NewDeleteNotificationsV1PreferencesPreferenceIDParamsWithContext creates a new DeleteNotificationsV1PreferencesPreferenceIDParams object
// with the ability to set a context for a request.
func NewDeleteNotificationsV1PreferencesPreferenceIDParamsWithContext(ctx context.Context) *DeleteNotificationsV1PreferencesPreferenceIDParams {
	return &DeleteNotificationsV1PreferencesPreferenceIDParams{
		Context: ctx,
	}
}

// NewDeleteNotificationsV1PreferencesPreferenceIDParamsWithHTTPClient creates a new DeleteNotificationsV1PreferencesPreferenceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNotificationsV1PreferencesPreferenceIDParamsWithHTTPClient(client *http.Client) *DeleteNotificationsV1PreferencesPreferenceIDParams {
	return &DeleteNotificationsV1PreferencesPreferenceIDParams{
		HTTPClient: client,
	}
}

/* DeleteNotificationsV1PreferencesPreferenceIDParams contains all the parameters to send to the API endpoint
   for the delete notifications v1 preferences preference ID operation.

   Typically these are written to a http.Request.
*/
type DeleteNotificationsV1PreferencesPreferenceIDParams struct {

	/* PreferenceID.

	   preference id
	*/
	PreferenceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete notifications v1 preferences preference ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) WithDefaults() *DeleteNotificationsV1PreferencesPreferenceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete notifications v1 preferences preference ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete notifications v1 preferences preference ID params
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) WithTimeout(timeout time.Duration) *DeleteNotificationsV1PreferencesPreferenceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete notifications v1 preferences preference ID params
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete notifications v1 preferences preference ID params
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) WithContext(ctx context.Context) *DeleteNotificationsV1PreferencesPreferenceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete notifications v1 preferences preference ID params
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete notifications v1 preferences preference ID params
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) WithHTTPClient(client *http.Client) *DeleteNotificationsV1PreferencesPreferenceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete notifications v1 preferences preference ID params
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPreferenceID adds the preferenceID to the delete notifications v1 preferences preference ID params
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) WithPreferenceID(preferenceID string) *DeleteNotificationsV1PreferencesPreferenceIDParams {
	o.SetPreferenceID(preferenceID)
	return o
}

// SetPreferenceID adds the preferenceId to the delete notifications v1 preferences preference ID params
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) SetPreferenceID(preferenceID string) {
	o.PreferenceID = preferenceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNotificationsV1PreferencesPreferenceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param preference_id
	if err := r.SetPathParam("preference_id", o.PreferenceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
