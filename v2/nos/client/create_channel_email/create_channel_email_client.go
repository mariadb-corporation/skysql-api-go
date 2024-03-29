// Code generated by go-swagger; DO NOT EDIT.

package create_channel_email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new create channel email API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for create channel email API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PostNotificationsV1SubscribersSubscriberIDChannelsEmails(params *PostNotificationsV1SubscribersSubscriberIDChannelsEmailsParams, opts ...ClientOption) (*PostNotificationsV1SubscribersSubscriberIDChannelsEmailsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PostNotificationsV1SubscribersSubscriberIDChannelsEmails creates email notification channel for a subscriber
*/
func (a *Client) PostNotificationsV1SubscribersSubscriberIDChannelsEmails(params *PostNotificationsV1SubscribersSubscriberIDChannelsEmailsParams, opts ...ClientOption) (*PostNotificationsV1SubscribersSubscriberIDChannelsEmailsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNotificationsV1SubscribersSubscriberIDChannelsEmailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostNotificationsV1SubscribersSubscriberIDChannelsEmails",
		Method:             "POST",
		PathPattern:        "/notifications/v1/subscribers/{subscriber_id}/channels/emails",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostNotificationsV1SubscribersSubscriberIDChannelsEmailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNotificationsV1SubscribersSubscriberIDChannelsEmailsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostNotificationsV1SubscribersSubscriberIDChannelsEmails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
