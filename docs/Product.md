# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortDescription** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**SysId** | **string** |  | 
**SysUpdatedBy** | **string** |  | 
**SysUpdatedOn** | **string** |  | 
**SysCreatedOn** | **string** |  | 
**DisplayName** | **string** |  | 
**Name** | **string** |  | 
**SysCreatedBy** | **string** |  | 
**SysTags** | **string** |  | 
**Active** | **string** |  | 
**Order** | **string** |  | 
**SysModCount** | **string** |  | 
**ActiveTopologies** | **string** |  | 
**DefaultTopology** | **string** |  | 

## Methods

### NewProduct

`func NewProduct(shortDescription string, sysId string, sysUpdatedBy string, sysUpdatedOn string, sysCreatedOn string, displayName string, name string, sysCreatedBy string, sysTags string, active string, order string, sysModCount string, activeTopologies string, defaultTopology string, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortDescription

`func (o *Product) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Product) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Product) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.


### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSysId

`func (o *Product) GetSysId() string`

GetSysId returns the SysId field if non-nil, zero value otherwise.

### GetSysIdOk

`func (o *Product) GetSysIdOk() (*string, bool)`

GetSysIdOk returns a tuple with the SysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysId

`func (o *Product) SetSysId(v string)`

SetSysId sets SysId field to given value.


### GetSysUpdatedBy

`func (o *Product) GetSysUpdatedBy() string`

GetSysUpdatedBy returns the SysUpdatedBy field if non-nil, zero value otherwise.

### GetSysUpdatedByOk

`func (o *Product) GetSysUpdatedByOk() (*string, bool)`

GetSysUpdatedByOk returns a tuple with the SysUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysUpdatedBy

`func (o *Product) SetSysUpdatedBy(v string)`

SetSysUpdatedBy sets SysUpdatedBy field to given value.


### GetSysUpdatedOn

`func (o *Product) GetSysUpdatedOn() string`

GetSysUpdatedOn returns the SysUpdatedOn field if non-nil, zero value otherwise.

### GetSysUpdatedOnOk

`func (o *Product) GetSysUpdatedOnOk() (*string, bool)`

GetSysUpdatedOnOk returns a tuple with the SysUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysUpdatedOn

`func (o *Product) SetSysUpdatedOn(v string)`

SetSysUpdatedOn sets SysUpdatedOn field to given value.


### GetSysCreatedOn

`func (o *Product) GetSysCreatedOn() string`

GetSysCreatedOn returns the SysCreatedOn field if non-nil, zero value otherwise.

### GetSysCreatedOnOk

`func (o *Product) GetSysCreatedOnOk() (*string, bool)`

GetSysCreatedOnOk returns a tuple with the SysCreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysCreatedOn

`func (o *Product) SetSysCreatedOn(v string)`

SetSysCreatedOn sets SysCreatedOn field to given value.


### GetDisplayName

`func (o *Product) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Product) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Product) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetSysCreatedBy

`func (o *Product) GetSysCreatedBy() string`

GetSysCreatedBy returns the SysCreatedBy field if non-nil, zero value otherwise.

### GetSysCreatedByOk

`func (o *Product) GetSysCreatedByOk() (*string, bool)`

GetSysCreatedByOk returns a tuple with the SysCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysCreatedBy

`func (o *Product) SetSysCreatedBy(v string)`

SetSysCreatedBy sets SysCreatedBy field to given value.


### GetSysTags

`func (o *Product) GetSysTags() string`

GetSysTags returns the SysTags field if non-nil, zero value otherwise.

### GetSysTagsOk

`func (o *Product) GetSysTagsOk() (*string, bool)`

GetSysTagsOk returns a tuple with the SysTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysTags

`func (o *Product) SetSysTags(v string)`

SetSysTags sets SysTags field to given value.


### GetActive

`func (o *Product) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Product) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Product) SetActive(v string)`

SetActive sets Active field to given value.


### GetOrder

`func (o *Product) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Product) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Product) SetOrder(v string)`

SetOrder sets Order field to given value.


### GetSysModCount

`func (o *Product) GetSysModCount() string`

GetSysModCount returns the SysModCount field if non-nil, zero value otherwise.

### GetSysModCountOk

`func (o *Product) GetSysModCountOk() (*string, bool)`

GetSysModCountOk returns a tuple with the SysModCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysModCount

`func (o *Product) SetSysModCount(v string)`

SetSysModCount sets SysModCount field to given value.


### GetActiveTopologies

`func (o *Product) GetActiveTopologies() string`

GetActiveTopologies returns the ActiveTopologies field if non-nil, zero value otherwise.

### GetActiveTopologiesOk

`func (o *Product) GetActiveTopologiesOk() (*string, bool)`

GetActiveTopologiesOk returns a tuple with the ActiveTopologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTopologies

`func (o *Product) SetActiveTopologies(v string)`

SetActiveTopologies sets ActiveTopologies field to given value.


### GetDefaultTopology

`func (o *Product) GetDefaultTopology() string`

GetDefaultTopology returns the DefaultTopology field if non-nil, zero value otherwise.

### GetDefaultTopologyOk

`func (o *Product) GetDefaultTopologyOk() (*string, bool)`

GetDefaultTopologyOk returns a tuple with the DefaultTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTopology

`func (o *Product) SetDefaultTopology(v string)`

SetDefaultTopology sets DefaultTopology field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


