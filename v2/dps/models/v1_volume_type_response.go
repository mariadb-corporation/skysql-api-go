// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VolumeTypeResponse cloud storage volume type
//
// swagger:model v1.VolumeTypeResponse
type V1VolumeTypeResponse struct {

	// created on
	CreatedOn int64 `json:"created_on,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is default
	IsDefault bool `json:"is_default,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// updated on
	UpdatedOn int64 `json:"updated_on,omitempty"`
}

// Validate validates this v1 volume type response
func (m *V1VolumeTypeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 volume type response based on context it is used
func (m *V1VolumeTypeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VolumeTypeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VolumeTypeResponse) UnmarshalBinary(b []byte) error {
	var res V1VolumeTypeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
