# ConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** |  | 
**Topology** | **string** |  | 
**SysModCount** | **string** |  | 
**Active** | **string** |  | 
**Project** | **string** |  | 
**OwnedBy** | **string** |  | 
**SysUpdatedOn** | **string** |  | 
**SysTags** | **string** |  | 
**SysId** | **string** |  | 
**SysUpdatedBy** | **string** |  | 
**Public** | **string** |  | 
**SysCreatedOn** | **string** |  | 
**Name** | **string** |  | 
**SysCreatedBy** | **string** |  | 
**ConfigurationVersions** | [**[]ConfigurationVersion**](ConfigurationVersion.md) |  | 

## Methods

### NewConfigurationResponse

`func NewConfigurationResponse(number string, topology string, sysModCount string, active string, project string, ownedBy string, sysUpdatedOn string, sysTags string, sysId string, sysUpdatedBy string, public string, sysCreatedOn string, name string, sysCreatedBy string, configurationVersions []ConfigurationVersion, ) *ConfigurationResponse`

NewConfigurationResponse instantiates a new ConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationResponseWithDefaults

`func NewConfigurationResponseWithDefaults() *ConfigurationResponse`

NewConfigurationResponseWithDefaults instantiates a new ConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ConfigurationResponse) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ConfigurationResponse) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ConfigurationResponse) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetTopology

`func (o *ConfigurationResponse) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *ConfigurationResponse) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *ConfigurationResponse) SetTopology(v string)`

SetTopology sets Topology field to given value.


### GetSysModCount

`func (o *ConfigurationResponse) GetSysModCount() string`

GetSysModCount returns the SysModCount field if non-nil, zero value otherwise.

### GetSysModCountOk

`func (o *ConfigurationResponse) GetSysModCountOk() (*string, bool)`

GetSysModCountOk returns a tuple with the SysModCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysModCount

`func (o *ConfigurationResponse) SetSysModCount(v string)`

SetSysModCount sets SysModCount field to given value.


### GetActive

`func (o *ConfigurationResponse) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConfigurationResponse) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConfigurationResponse) SetActive(v string)`

SetActive sets Active field to given value.


### GetProject

`func (o *ConfigurationResponse) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ConfigurationResponse) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ConfigurationResponse) SetProject(v string)`

SetProject sets Project field to given value.


### GetOwnedBy

`func (o *ConfigurationResponse) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ConfigurationResponse) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ConfigurationResponse) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetSysUpdatedOn

`func (o *ConfigurationResponse) GetSysUpdatedOn() string`

GetSysUpdatedOn returns the SysUpdatedOn field if non-nil, zero value otherwise.

### GetSysUpdatedOnOk

`func (o *ConfigurationResponse) GetSysUpdatedOnOk() (*string, bool)`

GetSysUpdatedOnOk returns a tuple with the SysUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysUpdatedOn

`func (o *ConfigurationResponse) SetSysUpdatedOn(v string)`

SetSysUpdatedOn sets SysUpdatedOn field to given value.


### GetSysTags

`func (o *ConfigurationResponse) GetSysTags() string`

GetSysTags returns the SysTags field if non-nil, zero value otherwise.

### GetSysTagsOk

`func (o *ConfigurationResponse) GetSysTagsOk() (*string, bool)`

GetSysTagsOk returns a tuple with the SysTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysTags

`func (o *ConfigurationResponse) SetSysTags(v string)`

SetSysTags sets SysTags field to given value.


### GetSysId

`func (o *ConfigurationResponse) GetSysId() string`

GetSysId returns the SysId field if non-nil, zero value otherwise.

### GetSysIdOk

`func (o *ConfigurationResponse) GetSysIdOk() (*string, bool)`

GetSysIdOk returns a tuple with the SysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysId

`func (o *ConfigurationResponse) SetSysId(v string)`

SetSysId sets SysId field to given value.


### GetSysUpdatedBy

`func (o *ConfigurationResponse) GetSysUpdatedBy() string`

GetSysUpdatedBy returns the SysUpdatedBy field if non-nil, zero value otherwise.

### GetSysUpdatedByOk

`func (o *ConfigurationResponse) GetSysUpdatedByOk() (*string, bool)`

GetSysUpdatedByOk returns a tuple with the SysUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysUpdatedBy

`func (o *ConfigurationResponse) SetSysUpdatedBy(v string)`

SetSysUpdatedBy sets SysUpdatedBy field to given value.


### GetPublic

`func (o *ConfigurationResponse) GetPublic() string`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ConfigurationResponse) GetPublicOk() (*string, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ConfigurationResponse) SetPublic(v string)`

SetPublic sets Public field to given value.


### GetSysCreatedOn

`func (o *ConfigurationResponse) GetSysCreatedOn() string`

GetSysCreatedOn returns the SysCreatedOn field if non-nil, zero value otherwise.

### GetSysCreatedOnOk

`func (o *ConfigurationResponse) GetSysCreatedOnOk() (*string, bool)`

GetSysCreatedOnOk returns a tuple with the SysCreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysCreatedOn

`func (o *ConfigurationResponse) SetSysCreatedOn(v string)`

SetSysCreatedOn sets SysCreatedOn field to given value.


### GetName

`func (o *ConfigurationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSysCreatedBy

`func (o *ConfigurationResponse) GetSysCreatedBy() string`

GetSysCreatedBy returns the SysCreatedBy field if non-nil, zero value otherwise.

### GetSysCreatedByOk

`func (o *ConfigurationResponse) GetSysCreatedByOk() (*string, bool)`

GetSysCreatedByOk returns a tuple with the SysCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysCreatedBy

`func (o *ConfigurationResponse) SetSysCreatedBy(v string)`

SetSysCreatedBy sets SysCreatedBy field to given value.


### GetConfigurationVersions

`func (o *ConfigurationResponse) GetConfigurationVersions() []ConfigurationVersion`

GetConfigurationVersions returns the ConfigurationVersions field if non-nil, zero value otherwise.

### GetConfigurationVersionsOk

`func (o *ConfigurationResponse) GetConfigurationVersionsOk() (*[]ConfigurationVersion, bool)`

GetConfigurationVersionsOk returns a tuple with the ConfigurationVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationVersions

`func (o *ConfigurationResponse) SetConfigurationVersions(v []ConfigurationVersion)`

SetConfigurationVersions sets ConfigurationVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


