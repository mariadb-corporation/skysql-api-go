// Code generated by go-swagger; DO NOT EDIT.

package delete_channel_email

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

// NewDeleteNotificationsV1ChannelsEmailsChannelIDParams creates a new DeleteNotificationsV1ChannelsEmailsChannelIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNotificationsV1ChannelsEmailsChannelIDParams() *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	return &DeleteNotificationsV1ChannelsEmailsChannelIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNotificationsV1ChannelsEmailsChannelIDParamsWithTimeout creates a new DeleteNotificationsV1ChannelsEmailsChannelIDParams object
// with the ability to set a timeout on a request.
func NewDeleteNotificationsV1ChannelsEmailsChannelIDParamsWithTimeout(timeout time.Duration) *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	return &DeleteNotificationsV1ChannelsEmailsChannelIDParams{
		timeout: timeout,
	}
}

// NewDeleteNotificationsV1ChannelsEmailsChannelIDParamsWithContext creates a new DeleteNotificationsV1ChannelsEmailsChannelIDParams object
// with the ability to set a context for a request.
func NewDeleteNotificationsV1ChannelsEmailsChannelIDParamsWithContext(ctx context.Context) *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	return &DeleteNotificationsV1ChannelsEmailsChannelIDParams{
		Context: ctx,
	}
}

// NewDeleteNotificationsV1ChannelsEmailsChannelIDParamsWithHTTPClient creates a new DeleteNotificationsV1ChannelsEmailsChannelIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNotificationsV1ChannelsEmailsChannelIDParamsWithHTTPClient(client *http.Client) *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	return &DeleteNotificationsV1ChannelsEmailsChannelIDParams{
		HTTPClient: client,
	}
}

/* DeleteNotificationsV1ChannelsEmailsChannelIDParams contains all the parameters to send to the API endpoint
   for the delete notifications v1 channels emails channel ID operation.

   Typically these are written to a http.Request.
*/
type DeleteNotificationsV1ChannelsEmailsChannelIDParams struct {

	/* ChannelID.

	   channel id
	*/
	ChannelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete notifications v1 channels emails channel ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) WithDefaults() *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete notifications v1 channels emails channel ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete notifications v1 channels emails channel ID params
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) WithTimeout(timeout time.Duration) *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete notifications v1 channels emails channel ID params
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete notifications v1 channels emails channel ID params
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) WithContext(ctx context.Context) *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete notifications v1 channels emails channel ID params
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete notifications v1 channels emails channel ID params
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) WithHTTPClient(client *http.Client) *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete notifications v1 channels emails channel ID params
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the delete notifications v1 channels emails channel ID params
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) WithChannelID(channelID string) *DeleteNotificationsV1ChannelsEmailsChannelIDParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the delete notifications v1 channels emails channel ID params
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel_id
	if err := r.SetPathParam("channel_id", o.ChannelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
