# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **string** |  | 
**SysModCount** | **string** |  | 
**SysUpdatedOn** | **string** |  | 
**DisplayName** | **string** |  | 
**SysTags** | **string** |  | 
**Type** | **string** |  | 
**EnterpriseVersion** | Pointer to **string** |  | [optional] 
**ReleaseNotesUrl** | **string** |  | 
**SysId** | **string** |  | 
**SysUpdatedBy** | **string** |  | 
**Public** | **string** |  | 
**Provider** | **string** |  | 
**ReleaseDate** | **string** |  | 
**SysCreatedOn** | **string** |  | 
**ParentRelease** | **string** |  | 
**Name** | **string** |  | 
**SysCreatedBy** | **string** |  | 

## Methods

### NewVersion

`func NewVersion(product string, sysModCount string, sysUpdatedOn string, displayName string, sysTags string, type_ string, releaseNotesUrl string, sysId string, sysUpdatedBy string, public string, provider string, releaseDate string, sysCreatedOn string, parentRelease string, name string, sysCreatedBy string, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *Version) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Version) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Version) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetSysModCount

`func (o *Version) GetSysModCount() string`

GetSysModCount returns the SysModCount field if non-nil, zero value otherwise.

### GetSysModCountOk

`func (o *Version) GetSysModCountOk() (*string, bool)`

GetSysModCountOk returns a tuple with the SysModCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysModCount

`func (o *Version) SetSysModCount(v string)`

SetSysModCount sets SysModCount field to given value.


### GetSysUpdatedOn

`func (o *Version) GetSysUpdatedOn() string`

GetSysUpdatedOn returns the SysUpdatedOn field if non-nil, zero value otherwise.

### GetSysUpdatedOnOk

`func (o *Version) GetSysUpdatedOnOk() (*string, bool)`

GetSysUpdatedOnOk returns a tuple with the SysUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysUpdatedOn

`func (o *Version) SetSysUpdatedOn(v string)`

SetSysUpdatedOn sets SysUpdatedOn field to given value.


### GetDisplayName

`func (o *Version) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Version) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Version) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSysTags

`func (o *Version) GetSysTags() string`

GetSysTags returns the SysTags field if non-nil, zero value otherwise.

### GetSysTagsOk

`func (o *Version) GetSysTagsOk() (*string, bool)`

GetSysTagsOk returns a tuple with the SysTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysTags

`func (o *Version) SetSysTags(v string)`

SetSysTags sets SysTags field to given value.


### GetType

`func (o *Version) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Version) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Version) SetType(v string)`

SetType sets Type field to given value.


### GetEnterpriseVersion

`func (o *Version) GetEnterpriseVersion() string`

GetEnterpriseVersion returns the EnterpriseVersion field if non-nil, zero value otherwise.

### GetEnterpriseVersionOk

`func (o *Version) GetEnterpriseVersionOk() (*string, bool)`

GetEnterpriseVersionOk returns a tuple with the EnterpriseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseVersion

`func (o *Version) SetEnterpriseVersion(v string)`

SetEnterpriseVersion sets EnterpriseVersion field to given value.

### HasEnterpriseVersion

`func (o *Version) HasEnterpriseVersion() bool`

HasEnterpriseVersion returns a boolean if a field has been set.

### GetReleaseNotesUrl

`func (o *Version) GetReleaseNotesUrl() string`

GetReleaseNotesUrl returns the ReleaseNotesUrl field if non-nil, zero value otherwise.

### GetReleaseNotesUrlOk

`func (o *Version) GetReleaseNotesUrlOk() (*string, bool)`

GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotesUrl

`func (o *Version) SetReleaseNotesUrl(v string)`

SetReleaseNotesUrl sets ReleaseNotesUrl field to given value.


### GetSysId

`func (o *Version) GetSysId() string`

GetSysId returns the SysId field if non-nil, zero value otherwise.

### GetSysIdOk

`func (o *Version) GetSysIdOk() (*string, bool)`

GetSysIdOk returns a tuple with the SysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysId

`func (o *Version) SetSysId(v string)`

SetSysId sets SysId field to given value.


### GetSysUpdatedBy

`func (o *Version) GetSysUpdatedBy() string`

GetSysUpdatedBy returns the SysUpdatedBy field if non-nil, zero value otherwise.

### GetSysUpdatedByOk

`func (o *Version) GetSysUpdatedByOk() (*string, bool)`

GetSysUpdatedByOk returns a tuple with the SysUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysUpdatedBy

`func (o *Version) SetSysUpdatedBy(v string)`

SetSysUpdatedBy sets SysUpdatedBy field to given value.


### GetPublic

`func (o *Version) GetPublic() string`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Version) GetPublicOk() (*string, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Version) SetPublic(v string)`

SetPublic sets Public field to given value.


### GetProvider

`func (o *Version) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Version) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Version) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetReleaseDate

`func (o *Version) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Version) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Version) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetSysCreatedOn

`func (o *Version) GetSysCreatedOn() string`

GetSysCreatedOn returns the SysCreatedOn field if non-nil, zero value otherwise.

### GetSysCreatedOnOk

`func (o *Version) GetSysCreatedOnOk() (*string, bool)`

GetSysCreatedOnOk returns a tuple with the SysCreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysCreatedOn

`func (o *Version) SetSysCreatedOn(v string)`

SetSysCreatedOn sets SysCreatedOn field to given value.


### GetParentRelease

`func (o *Version) GetParentRelease() string`

GetParentRelease returns the ParentRelease field if non-nil, zero value otherwise.

### GetParentReleaseOk

`func (o *Version) GetParentReleaseOk() (*string, bool)`

GetParentReleaseOk returns a tuple with the ParentRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRelease

`func (o *Version) SetParentRelease(v string)`

SetParentRelease sets ParentRelease field to given value.


### GetName

`func (o *Version) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Version) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Version) SetName(v string)`

SetName sets Name field to given value.


### GetSysCreatedBy

`func (o *Version) GetSysCreatedBy() string`

GetSysCreatedBy returns the SysCreatedBy field if non-nil, zero value otherwise.

### GetSysCreatedByOk

`func (o *Version) GetSysCreatedByOk() (*string, bool)`

GetSysCreatedByOk returns a tuple with the SysCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysCreatedBy

`func (o *Version) SetSysCreatedBy(v string)`

SetSysCreatedBy sets SysCreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


