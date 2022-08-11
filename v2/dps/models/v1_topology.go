// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Topology Cluster topology valid for a particular product, e.g. Standalone or MaxScale
//
// swagger:model v1.Topology
type V1Topology struct {

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nodes
	Nodes interface{} `json:"nodes,omitempty"`

	// order
	Order int64 `json:"order,omitempty"`

	// service type
	ServiceType string `json:"service_type,omitempty"`

	// storage engine
	StorageEngine string `json:"storage_engine,omitempty"`
}

// Validate validates this v1 topology
func (m *V1Topology) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 topology based on context it is used
func (m *V1Topology) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Topology) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Topology) UnmarshalBinary(b []byte) error {
	var res V1Topology
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}