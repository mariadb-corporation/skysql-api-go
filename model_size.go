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

// Size Node size, as defined by the providers
type Size struct {
	Product string `json:"product"`
	Visibility string `json:"visibility"`
	SysModCount string `json:"sys_mod_count"`
	Active string `json:"active"`
	Cpu string `json:"cpu"`
	SysUpdatedOn string `json:"sys_updated_on"`
	SysTags string `json:"sys_tags"`
	Sequence string `json:"sequence"`
	SysId string `json:"sys_id"`
	Component string `json:"component"`
	SysUpdatedBy string `json:"sys_updated_by"`
	Tier string `json:"tier"`
	Provider string `json:"provider"`
	SysCreatedOn string `json:"sys_created_on"`
	Name string `json:"name"`
	NodePool string `json:"node_pool"`
	Value string `json:"value"`
	SysCreatedBy string `json:"sys_created_by"`
	ApiHandle string `json:"api_handle"`
	Ram string `json:"ram"`
}

// NewSize instantiates a new Size object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSize(product string, visibility string, sysModCount string, active string, cpu string, sysUpdatedOn string, sysTags string, sequence string, sysId string, component string, sysUpdatedBy string, tier string, provider string, sysCreatedOn string, name string, nodePool string, value string, sysCreatedBy string, apiHandle string, ram string) *Size {
	this := Size{}
	this.Product = product
	this.Visibility = visibility
	this.SysModCount = sysModCount
	this.Active = active
	this.Cpu = cpu
	this.SysUpdatedOn = sysUpdatedOn
	this.SysTags = sysTags
	this.Sequence = sequence
	this.SysId = sysId
	this.Component = component
	this.SysUpdatedBy = sysUpdatedBy
	this.Tier = tier
	this.Provider = provider
	this.SysCreatedOn = sysCreatedOn
	this.Name = name
	this.NodePool = nodePool
	this.Value = value
	this.SysCreatedBy = sysCreatedBy
	this.ApiHandle = apiHandle
	this.Ram = ram
	return &this
}

// NewSizeWithDefaults instantiates a new Size object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSizeWithDefaults() *Size {
	this := Size{}
	return &this
}

// GetProduct returns the Product field value
func (o *Size) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *Size) GetProductOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *Size) SetProduct(v string) {
	o.Product = v
}

// GetVisibility returns the Visibility field value
func (o *Size) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *Size) GetVisibilityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *Size) SetVisibility(v string) {
	o.Visibility = v
}

// GetSysModCount returns the SysModCount field value
func (o *Size) GetSysModCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SysModCount
}

// GetSysModCountOk returns a tuple with the SysModCount field value
// and a boolean to check if the value has been set.
func (o *Size) GetSysModCountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SysModCount, true
}

// SetSysModCount sets field value
func (o *Size) SetSysModCount(v string) {
	o.SysModCount = v
}

// GetActive returns the Active field value
func (o *Size) GetActive() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *Size) GetActiveOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *Size) SetActive(v string) {
	o.Active = v
}

// GetCpu returns the Cpu field value
func (o *Size) GetCpu() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Size) GetCpuOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *Size) SetCpu(v string) {
	o.Cpu = v
}

// GetSysUpdatedOn returns the SysUpdatedOn field value
func (o *Size) GetSysUpdatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SysUpdatedOn
}

// GetSysUpdatedOnOk returns a tuple with the SysUpdatedOn field value
// and a boolean to check if the value has been set.
func (o *Size) GetSysUpdatedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SysUpdatedOn, true
}

// SetSysUpdatedOn sets field value
func (o *Size) SetSysUpdatedOn(v string) {
	o.SysUpdatedOn = v
}

// GetSysTags returns the SysTags field value
func (o *Size) GetSysTags() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SysTags
}

// GetSysTagsOk returns a tuple with the SysTags field value
// and a boolean to check if the value has been set.
func (o *Size) GetSysTagsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SysTags, true
}

// SetSysTags sets field value
func (o *Size) SetSysTags(v string) {
	o.SysTags = v
}

// GetSequence returns the Sequence field value
func (o *Size) GetSequence() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *Size) GetSequenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *Size) SetSequence(v string) {
	o.Sequence = v
}

// GetSysId returns the SysId field value
func (o *Size) GetSysId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SysId
}

// GetSysIdOk returns a tuple with the SysId field value
// and a boolean to check if the value has been set.
func (o *Size) GetSysIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SysId, true
}

// SetSysId sets field value
func (o *Size) SetSysId(v string) {
	o.SysId = v
}

// GetComponent returns the Component field value
func (o *Size) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *Size) GetComponentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *Size) SetComponent(v string) {
	o.Component = v
}

// GetSysUpdatedBy returns the SysUpdatedBy field value
func (o *Size) GetSysUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SysUpdatedBy
}

// GetSysUpdatedByOk returns a tuple with the SysUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *Size) GetSysUpdatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SysUpdatedBy, true
}

// SetSysUpdatedBy sets field value
func (o *Size) SetSysUpdatedBy(v string) {
	o.SysUpdatedBy = v
}

// GetTier returns the Tier field value
func (o *Size) GetTier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *Size) GetTierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *Size) SetTier(v string) {
	o.Tier = v
}

// GetProvider returns the Provider field value
func (o *Size) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Size) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *Size) SetProvider(v string) {
	o.Provider = v
}

// GetSysCreatedOn returns the SysCreatedOn field value
func (o *Size) GetSysCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SysCreatedOn
}

// GetSysCreatedOnOk returns a tuple with the SysCreatedOn field value
// and a boolean to check if the value has been set.
func (o *Size) GetSysCreatedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SysCreatedOn, true
}

// SetSysCreatedOn sets field value
func (o *Size) SetSysCreatedOn(v string) {
	o.SysCreatedOn = v
}

// GetName returns the Name field value
func (o *Size) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Size) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Size) SetName(v string) {
	o.Name = v
}

// GetNodePool returns the NodePool field value
func (o *Size) GetNodePool() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodePool
}

// GetNodePoolOk returns a tuple with the NodePool field value
// and a boolean to check if the value has been set.
func (o *Size) GetNodePoolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodePool, true
}

// SetNodePool sets field value
func (o *Size) SetNodePool(v string) {
	o.NodePool = v
}

// GetValue returns the Value field value
func (o *Size) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Size) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Size) SetValue(v string) {
	o.Value = v
}

// GetSysCreatedBy returns the SysCreatedBy field value
func (o *Size) GetSysCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SysCreatedBy
}

// GetSysCreatedByOk returns a tuple with the SysCreatedBy field value
// and a boolean to check if the value has been set.
func (o *Size) GetSysCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SysCreatedBy, true
}

// SetSysCreatedBy sets field value
func (o *Size) SetSysCreatedBy(v string) {
	o.SysCreatedBy = v
}

// GetApiHandle returns the ApiHandle field value
func (o *Size) GetApiHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiHandle
}

// GetApiHandleOk returns a tuple with the ApiHandle field value
// and a boolean to check if the value has been set.
func (o *Size) GetApiHandleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiHandle, true
}

// SetApiHandle sets field value
func (o *Size) SetApiHandle(v string) {
	o.ApiHandle = v
}

// GetRam returns the Ram field value
func (o *Size) GetRam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *Size) GetRamOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *Size) SetRam(v string) {
	o.Ram = v
}

func (o Size) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["product"] = o.Product
	}
	if true {
		toSerialize["visibility"] = o.Visibility
	}
	if true {
		toSerialize["sys_mod_count"] = o.SysModCount
	}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["cpu"] = o.Cpu
	}
	if true {
		toSerialize["sys_updated_on"] = o.SysUpdatedOn
	}
	if true {
		toSerialize["sys_tags"] = o.SysTags
	}
	if true {
		toSerialize["sequence"] = o.Sequence
	}
	if true {
		toSerialize["sys_id"] = o.SysId
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["sys_updated_by"] = o.SysUpdatedBy
	}
	if true {
		toSerialize["tier"] = o.Tier
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["sys_created_on"] = o.SysCreatedOn
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["node_pool"] = o.NodePool
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["sys_created_by"] = o.SysCreatedBy
	}
	if true {
		toSerialize["api_handle"] = o.ApiHandle
	}
	if true {
		toSerialize["ram"] = o.Ram
	}
	return json.Marshal(toSerialize)
}

type NullableSize struct {
	value *Size
	isSet bool
}

func (v NullableSize) Get() *Size {
	return v.value
}

func (v *NullableSize) Set(val *Size) {
	v.value = val
	v.isSet = true
}

func (v NullableSize) IsSet() bool {
	return v.isSet
}

func (v *NullableSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSize(val *Size) *NullableSize {
	return &NullableSize{value: val, isSet: true}
}

func (v NullableSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


