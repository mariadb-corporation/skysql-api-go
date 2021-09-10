# QuotaProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Limit** | **int32** |  | 
**Remaining** | Pointer to **int32** |  | [optional] 
**Used** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewQuotaProgress

`func NewQuotaProgress(name string, description string, limit int32, ) *QuotaProgress`

NewQuotaProgress instantiates a new QuotaProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaProgressWithDefaults

`func NewQuotaProgressWithDefaults() *QuotaProgress`

NewQuotaProgressWithDefaults instantiates a new QuotaProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuotaProgress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuotaProgress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuotaProgress) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *QuotaProgress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotaProgress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotaProgress) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLimit

`func (o *QuotaProgress) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuotaProgress) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuotaProgress) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetRemaining

`func (o *QuotaProgress) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *QuotaProgress) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *QuotaProgress) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *QuotaProgress) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetUsed

`func (o *QuotaProgress) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *QuotaProgress) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *QuotaProgress) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *QuotaProgress) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetRegion

`func (o *QuotaProgress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *QuotaProgress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *QuotaProgress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *QuotaProgress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetProvider

`func (o *QuotaProgress) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *QuotaProgress) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *QuotaProgress) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *QuotaProgress) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


