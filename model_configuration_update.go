/*
SkySQL API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package skysql

import (
	"encoding/json"
)

// ConfigurationUpdate Request body to update a configuration
type ConfigurationUpdate struct {
	Name *string `json:"name,omitempty"`
	ConfigJson *string `json:"config_json,omitempty"`
}

// NewConfigurationUpdate instantiates a new ConfigurationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationUpdate() *ConfigurationUpdate {
	this := ConfigurationUpdate{}
	var configJson string = "{}"
	this.ConfigJson = &configJson
	return &this
}

// NewConfigurationUpdateWithDefaults instantiates a new ConfigurationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationUpdateWithDefaults() *ConfigurationUpdate {
	this := ConfigurationUpdate{}
	var configJson string = "{}"
	this.ConfigJson = &configJson
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigurationUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigurationUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigurationUpdate) SetName(v string) {
	o.Name = &v
}

// GetConfigJson returns the ConfigJson field value if set, zero value otherwise.
func (o *ConfigurationUpdate) GetConfigJson() string {
	if o == nil || o.ConfigJson == nil {
		var ret string
		return ret
	}
	return *o.ConfigJson
}

// GetConfigJsonOk returns a tuple with the ConfigJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationUpdate) GetConfigJsonOk() (*string, bool) {
	if o == nil || o.ConfigJson == nil {
		return nil, false
	}
	return o.ConfigJson, true
}

// HasConfigJson returns a boolean if a field has been set.
func (o *ConfigurationUpdate) HasConfigJson() bool {
	if o != nil && o.ConfigJson != nil {
		return true
	}

	return false
}

// SetConfigJson gets a reference to the given string and assigns it to the ConfigJson field.
func (o *ConfigurationUpdate) SetConfigJson(v string) {
	o.ConfigJson = &v
}

func (o ConfigurationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ConfigJson != nil {
		toSerialize["config_json"] = o.ConfigJson
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationUpdate struct {
	value *ConfigurationUpdate
	isSet bool
}

func (v NullableConfigurationUpdate) Get() *ConfigurationUpdate {
	return v.value
}

func (v *NullableConfigurationUpdate) Set(val *ConfigurationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationUpdate(val *ConfigurationUpdate) *NullableConfigurationUpdate {
	return &NullableConfigurationUpdate{value: val, isSet: true}
}

func (v NullableConfigurationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


