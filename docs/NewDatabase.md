# NewDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseVersion** | **string** |  | 
**Topology** | **string** |  | 
**Size** | **string** |  | 
**TxStorage** | **string** |  | 
**MaxscaleConfig** | **string** |  | 
**Name** | **string** |  | 
**Region** | **string** |  | 
**ReplRegion** | **string** |  | 
**Provider** | **string** |  | 
**Replicas** | **string** |  | 
**Monitor** | **string** |  | 
**MaxscaleProxy** | **string** |  | 

## Methods

### NewNewDatabase

`func NewNewDatabase(releaseVersion string, topology string, size string, txStorage string, maxscaleConfig string, name string, region string, replRegion string, provider string, replicas string, monitor string, maxscaleProxy string, ) *NewDatabase`

NewNewDatabase instantiates a new NewDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDatabaseWithDefaults

`func NewNewDatabaseWithDefaults() *NewDatabase`

NewNewDatabaseWithDefaults instantiates a new NewDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseVersion

`func (o *NewDatabase) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *NewDatabase) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *NewDatabase) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.


### GetTopology

`func (o *NewDatabase) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *NewDatabase) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *NewDatabase) SetTopology(v string)`

SetTopology sets Topology field to given value.


### GetSize

`func (o *NewDatabase) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NewDatabase) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NewDatabase) SetSize(v string)`

SetSize sets Size field to given value.


### GetTxStorage

`func (o *NewDatabase) GetTxStorage() string`

GetTxStorage returns the TxStorage field if non-nil, zero value otherwise.

### GetTxStorageOk

`func (o *NewDatabase) GetTxStorageOk() (*string, bool)`

GetTxStorageOk returns a tuple with the TxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStorage

`func (o *NewDatabase) SetTxStorage(v string)`

SetTxStorage sets TxStorage field to given value.


### GetMaxscaleConfig

`func (o *NewDatabase) GetMaxscaleConfig() string`

GetMaxscaleConfig returns the MaxscaleConfig field if non-nil, zero value otherwise.

### GetMaxscaleConfigOk

`func (o *NewDatabase) GetMaxscaleConfigOk() (*string, bool)`

GetMaxscaleConfigOk returns a tuple with the MaxscaleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxscaleConfig

`func (o *NewDatabase) SetMaxscaleConfig(v string)`

SetMaxscaleConfig sets MaxscaleConfig field to given value.


### GetName

`func (o *NewDatabase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewDatabase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewDatabase) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *NewDatabase) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NewDatabase) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NewDatabase) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetReplRegion

`func (o *NewDatabase) GetReplRegion() string`

GetReplRegion returns the ReplRegion field if non-nil, zero value otherwise.

### GetReplRegionOk

`func (o *NewDatabase) GetReplRegionOk() (*string, bool)`

GetReplRegionOk returns a tuple with the ReplRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplRegion

`func (o *NewDatabase) SetReplRegion(v string)`

SetReplRegion sets ReplRegion field to given value.


### GetProvider

`func (o *NewDatabase) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NewDatabase) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NewDatabase) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetReplicas

`func (o *NewDatabase) GetReplicas() string`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *NewDatabase) GetReplicasOk() (*string, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *NewDatabase) SetReplicas(v string)`

SetReplicas sets Replicas field to given value.


### GetMonitor

`func (o *NewDatabase) GetMonitor() string`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *NewDatabase) GetMonitorOk() (*string, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *NewDatabase) SetMonitor(v string)`

SetMonitor sets Monitor field to given value.


### GetMaxscaleProxy

`func (o *NewDatabase) GetMaxscaleProxy() string`

GetMaxscaleProxy returns the MaxscaleProxy field if non-nil, zero value otherwise.

### GetMaxscaleProxyOk

`func (o *NewDatabase) GetMaxscaleProxyOk() (*string, bool)`

GetMaxscaleProxyOk returns a tuple with the MaxscaleProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxscaleProxy

`func (o *NewDatabase) SetMaxscaleProxy(v string)`

SetMaxscaleProxy sets MaxscaleProxy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


