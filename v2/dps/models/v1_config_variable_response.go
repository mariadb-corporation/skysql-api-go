// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ConfigVariableResponse ConfigVariableResponse
//
// swagger:model v1.ConfigVariableResponse
type V1ConfigVariableResponse struct {

	// keywords
	Keywords []string `json:"keywords"`

	// name
	Name string `json:"name,omitempty"`

	// requires restart
	RequiresRestart bool `json:"requires_restart,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// values
	Values []string `json:"values"`
}

// Validate validates this v1 config variable response
func (m *V1ConfigVariableResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 config variable response based on context it is used
func (m *V1ConfigVariableResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ConfigVariableResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ConfigVariableResponse) UnmarshalBinary(b []byte) error {
	var res V1ConfigVariableResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
