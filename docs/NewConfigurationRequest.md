# NewConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Topology** | **string** |  | 
**ConfigJson** | Pointer to **string** |  | [optional] [default to "{}"]

## Methods

### NewNewConfigurationRequest

`func NewNewConfigurationRequest(name string, topology string, ) *NewConfigurationRequest`

NewNewConfigurationRequest instantiates a new NewConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewConfigurationRequestWithDefaults

`func NewNewConfigurationRequestWithDefaults() *NewConfigurationRequest`

NewNewConfigurationRequestWithDefaults instantiates a new NewConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewConfigurationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewConfigurationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewConfigurationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTopology

`func (o *NewConfigurationRequest) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *NewConfigurationRequest) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *NewConfigurationRequest) SetTopology(v string)`

SetTopology sets Topology field to given value.


### GetConfigJson

`func (o *NewConfigurationRequest) GetConfigJson() string`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *NewConfigurationRequest) GetConfigJsonOk() (*string, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *NewConfigurationRequest) SetConfigJson(v string)`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *NewConfigurationRequest) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


