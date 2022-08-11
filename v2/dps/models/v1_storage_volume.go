// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1StorageVolume v1 storage volume
//
// swagger:model v1.StorageVolume
type V1StorageVolume struct {

	// iops
	Iops int64 `json:"iops,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// volume type
	VolumeType string `json:"volume_type,omitempty"`
}

// Validate validates this v1 storage volume
func (m *V1StorageVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 storage volume based on context it is used
func (m *V1StorageVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1StorageVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1StorageVolume) UnmarshalBinary(b []byte) error {
	var res V1StorageVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
