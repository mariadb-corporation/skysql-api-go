// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ServiceSizeState Nodes size for a service
//
// swagger:model v1.ServiceSizeState
type V1ServiceSizeState struct {

	// size
	// Example: sky-2x8
	Size string `json:"size,omitempty"`
}

// Validate validates this v1 service size state
func (m *V1ServiceSizeState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 service size state based on context it is used
func (m *V1ServiceSizeState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ServiceSizeState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ServiceSizeState) UnmarshalBinary(b []byte) error {
	var res V1ServiceSizeState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
