// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetProvisioningV1InternalAccountsOrgIDProvider(params *GetProvisioningV1InternalAccountsOrgIDProviderParams, opts ...ClientOption) (*GetProvisioningV1InternalAccountsOrgIDProviderOK, error)

	PatchProvisioningV1InternalAccountsOrgIDProvider(params *PatchProvisioningV1InternalAccountsOrgIDProviderParams, opts ...ClientOption) (*PatchProvisioningV1InternalAccountsOrgIDProviderNoContent, error)

	PostProvisioningV1InternalAccounts(params *PostProvisioningV1InternalAccountsParams, opts ...ClientOption) (*PostProvisioningV1InternalAccountsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetProvisioningV1InternalAccountsOrgIDProvider gets account by org id and provider

  Retrieve an account by org_id and provider
*/
func (a *Client) GetProvisioningV1InternalAccountsOrgIDProvider(params *GetProvisioningV1InternalAccountsOrgIDProviderParams, opts ...ClientOption) (*GetProvisioningV1InternalAccountsOrgIDProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProvisioningV1InternalAccountsOrgIDProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProvisioningV1InternalAccountsOrgIDProvider",
		Method:             "GET",
		PathPattern:        "/provisioning/v1/internal/accounts/{org_id}/{provider}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProvisioningV1InternalAccountsOrgIDProviderReader{formats: a.formats},
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
	success, ok := result.(*GetProvisioningV1InternalAccountsOrgIDProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProvisioningV1InternalAccountsOrgIDProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchProvisioningV1InternalAccountsOrgIDProvider updates account

  Update the account
*/
func (a *Client) PatchProvisioningV1InternalAccountsOrgIDProvider(params *PatchProvisioningV1InternalAccountsOrgIDProviderParams, opts ...ClientOption) (*PatchProvisioningV1InternalAccountsOrgIDProviderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchProvisioningV1InternalAccountsOrgIDProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchProvisioningV1InternalAccountsOrgIDProvider",
		Method:             "PATCH",
		PathPattern:        "/provisioning/v1/internal/accounts/{org_id}/{provider}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchProvisioningV1InternalAccountsOrgIDProviderReader{formats: a.formats},
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
	success, ok := result.(*PatchProvisioningV1InternalAccountsOrgIDProviderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchProvisioningV1InternalAccountsOrgIDProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostProvisioningV1InternalAccounts creates a new account

  Create a new account
*/
func (a *Client) PostProvisioningV1InternalAccounts(params *PostProvisioningV1InternalAccountsParams, opts ...ClientOption) (*PostProvisioningV1InternalAccountsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProvisioningV1InternalAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostProvisioningV1InternalAccounts",
		Method:             "POST",
		PathPattern:        "/provisioning/v1/internal/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostProvisioningV1InternalAccountsReader{formats: a.formats},
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
	success, ok := result.(*PostProvisioningV1InternalAccountsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostProvisioningV1InternalAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
