// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoChannelEmail dto channel email
//
// swagger:model dto.ChannelEmail
type DtoChannelEmail struct {

	// channel name
	ChannelName string `json:"channel_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// receivers
	Receivers []*DtoEmailRecipients `json:"receivers"`

	// subscriber id
	SubscriberID string `json:"subscriber_id,omitempty"`
}

// Validate validates this dto channel email
func (m *DtoChannelEmail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReceivers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoChannelEmail) validateReceivers(formats strfmt.Registry) error {
	if swag.IsZero(m.Receivers) { // not required
		return nil
	}

	for i := 0; i < len(m.Receivers); i++ {
		if swag.IsZero(m.Receivers[i]) { // not required
			continue
		}

		if m.Receivers[i] != nil {
			if err := m.Receivers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dto channel email based on the context it is used
func (m *DtoChannelEmail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReceivers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoChannelEmail) contextValidateReceivers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Receivers); i++ {

		if m.Receivers[i] != nil {
			if err := m.Receivers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DtoChannelEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoChannelEmail) UnmarshalBinary(b []byte) error {
	var res DtoChannelEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
