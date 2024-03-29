// Code generated by go-swagger; DO NOT EDIT.

package update_channel_email

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

	"github.com/mariadb-corporation/skysql-api-go/v2/nos/models"
)

// NewPatchNotificationsV1ChannelsEmailsChannelIDParams creates a new PatchNotificationsV1ChannelsEmailsChannelIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchNotificationsV1ChannelsEmailsChannelIDParams() *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	return &PatchNotificationsV1ChannelsEmailsChannelIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchNotificationsV1ChannelsEmailsChannelIDParamsWithTimeout creates a new PatchNotificationsV1ChannelsEmailsChannelIDParams object
// with the ability to set a timeout on a request.
func NewPatchNotificationsV1ChannelsEmailsChannelIDParamsWithTimeout(timeout time.Duration) *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	return &PatchNotificationsV1ChannelsEmailsChannelIDParams{
		timeout: timeout,
	}
}

// NewPatchNotificationsV1ChannelsEmailsChannelIDParamsWithContext creates a new PatchNotificationsV1ChannelsEmailsChannelIDParams object
// with the ability to set a context for a request.
func NewPatchNotificationsV1ChannelsEmailsChannelIDParamsWithContext(ctx context.Context) *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	return &PatchNotificationsV1ChannelsEmailsChannelIDParams{
		Context: ctx,
	}
}

// NewPatchNotificationsV1ChannelsEmailsChannelIDParamsWithHTTPClient creates a new PatchNotificationsV1ChannelsEmailsChannelIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchNotificationsV1ChannelsEmailsChannelIDParamsWithHTTPClient(client *http.Client) *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	return &PatchNotificationsV1ChannelsEmailsChannelIDParams{
		HTTPClient: client,
	}
}

/* PatchNotificationsV1ChannelsEmailsChannelIDParams contains all the parameters to send to the API endpoint
   for the patch notifications v1 channels emails channel ID operation.

   Typically these are written to a http.Request.
*/
type PatchNotificationsV1ChannelsEmailsChannelIDParams struct {

	/* ChannelID.

	   channel id
	*/
	ChannelID string

	/* Request.

	   Request Body
	*/
	Request *models.DtoUpdateChannelEmail

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch notifications v1 channels emails channel ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) WithDefaults() *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch notifications v1 channels emails channel ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) WithTimeout(timeout time.Duration) *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) WithContext(ctx context.Context) *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) WithHTTPClient(client *http.Client) *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) WithChannelID(channelID string) *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithRequest adds the request to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) WithRequest(request *models.DtoUpdateChannelEmail) *PatchNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the patch notifications v1 channels emails channel ID params
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) SetRequest(request *models.DtoUpdateChannelEmail) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PatchNotificationsV1ChannelsEmailsChannelIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel_id
	if err := r.SetPathParam("channel_id", o.ChannelID); err != nil {
		return err
	}
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
