# ConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ConfigJson** | Pointer to **string** |  | [optional] [default to "{}"]

## Methods

### NewConfigurationUpdate

`func NewConfigurationUpdate() *ConfigurationUpdate`

NewConfigurationUpdate instantiates a new ConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationUpdateWithDefaults

`func NewConfigurationUpdateWithDefaults() *ConfigurationUpdate`

NewConfigurationUpdateWithDefaults instantiates a new ConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigurationUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigurationUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigJson

`func (o *ConfigurationUpdate) GetConfigJson() string`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *ConfigurationUpdate) GetConfigJsonOk() (*string, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *ConfigurationUpdate) SetConfigJson(v string)`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *ConfigurationUpdate) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


