// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ServiceResponseInternal Service object
//
// swagger:model v1.ServiceResponseInternal
type V1ServiceResponseInternal struct {

	// applications
	Applications interface{} `json:"applications,omitempty"`

	// availability zone
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// cert
	Cert interface{} `json:"cert,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// created on
	CreatedOn int64 `json:"created_on,omitempty"`

	// custom image versions
	CustomImageVersions interface{} `json:"custom_image_versions,omitempty"`

	// endpoints
	Endpoints []*V1Endpoint `json:"endpoints"`

	// fqdn
	Fqdn string `json:"fqdn,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ips
	Ips interface{} `json:"ips,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nodes
	Nodes int64 `json:"nodes,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// proxy options
	ProxyOptions interface{} `json:"proxy_options,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// service number
	ServiceNumber string `json:"service_number,omitempty"`

	// size
	Size string `json:"size,omitempty"`

	// ssl enabled
	SslEnabled bool `json:"ssl_enabled,omitempty"`

	// status
	// Enum: [ready failed pending stopped retired]
	Status string `json:"status,omitempty"`

	// storage volume
	StorageVolume *V1StorageVolume `json:"storage_volume,omitempty"`

	// tier
	Tier string `json:"tier,omitempty"`

	// topology
	Topology string `json:"topology,omitempty"`

	// updated by
	UpdatedBy string `json:"updated_by,omitempty"`

	// updated on
	UpdatedOn int64 `json:"updated_on,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 service response internal
func (m *V1ServiceResponseInternal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ServiceResponseInternal) validateEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Endpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.Endpoints); i++ {
		if swag.IsZero(m.Endpoints[i]) { // not required
			continue
		}

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var v1ServiceResponseInternalTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ready","failed","pending","stopped","retired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ServiceResponseInternalTypeStatusPropEnum = append(v1ServiceResponseInternalTypeStatusPropEnum, v)
	}
}

const (

	// V1ServiceResponseInternalStatusReady captures enum value "ready"
	V1ServiceResponseInternalStatusReady string = "ready"

	// V1ServiceResponseInternalStatusFailed captures enum value "failed"
	V1ServiceResponseInternalStatusFailed string = "failed"

	// V1ServiceResponseInternalStatusPending captures enum value "pending"
	V1ServiceResponseInternalStatusPending string = "pending"

	// V1ServiceResponseInternalStatusStopped captures enum value "stopped"
	V1ServiceResponseInternalStatusStopped string = "stopped"

	// V1ServiceResponseInternalStatusRetired captures enum value "retired"
	V1ServiceResponseInternalStatusRetired string = "retired"
)

// prop value enum
func (m *V1ServiceResponseInternal) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1ServiceResponseInternalTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1ServiceResponseInternal) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *V1ServiceResponseInternal) validateStorageVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageVolume) { // not required
		return nil
	}

	if m.StorageVolume != nil {
		if err := m.StorageVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 service response internal based on the context it is used
func (m *V1ServiceResponseInternal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ServiceResponseInternal) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Endpoints); i++ {

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ServiceResponseInternal) contextValidateStorageVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageVolume != nil {
		if err := m.StorageVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ServiceResponseInternal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ServiceResponseInternal) UnmarshalBinary(b []byte) error {
	var res V1ServiceResponseInternal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}