# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** |  | [optional] 
**OperationalStatus** | **string** |  | 
**EnablePamAuthentication** | **string** |  | 
**SysUpdatedOn** | **string** |  | 
**Number** | **string** |  | 
**InstanceState** | **string** |  | 
**ReadOnlyPort** | **string** |  | 
**ReadWritePort** | **string** |  | 
**Id** | **string** |  | 
**ReleaseVersion** | **string** |  | 
**GlAccount** | **string** |  | 
**SysCreatedBy** | **string** |  | 
**SslCertificate** | **string** |  | 
**ColumnstoreBucket** | **string** |  | 
**Topology** | **string** |  | 
**OwnedBy** | **string** |  | 
**Proxy** | **string** |  | 
**Size** | **string** |  | 
**DnsDomain** | **string** |  | 
**TxStorage** | **string** |  | 
**SslExpiresOn** | **string** |  | 
**ReplMasterHostExt** | **string** |  | 
**MaxscaleConfig** | **string** |  | 
**VolumeIops** | **string** |  | 
**Attributes** | **string** |  | 
**ReplicationStatus** | Pointer to **string** |  | [optional] 
**SkipSync** | **string** |  | 
**ReplicationType** | Pointer to **string** |  | [optional] 
**ReplMaster** | **string** |  | 
**SysUpdatedBy** | **string** |  | 
**BulkdataPort2** | **string** |  | 
**SysCreatedOn** | **string** |  | 
**BulkdataPort1** | **string** |  | 
**ActiveReplicas** | **string** |  | 
**Fqdn** | **string** |  | 
**SslSerial** | **string** |  | 
**InstallStatus** | **string** |  | 
**Name** | **string** |  | 
**Region** | Pointer to **string** |  | [optional] 
**ReplRegion** | Pointer to **string** |  | [optional] 
**CustomConfig** | **string** |  | 
**SysId** | **string** |  | 
**Provider** | **string** |  | 
**MacAddress** | **string** |  | 
**Replicas** | **string** |  | 
**SysModCount** | **string** |  | 
**Monitor** | **string** |  | 
**IpAddress** | **string** |  | 
**MaxscaleProxy** | **string** |  | 
**FaultCount** | **string** |  | 

## Methods

### NewDatabase

`func NewDatabase(operationalStatus string, enablePamAuthentication string, sysUpdatedOn string, number string, instanceState string, readOnlyPort string, readWritePort string, id string, releaseVersion string, glAccount string, sysCreatedBy string, sslCertificate string, columnstoreBucket string, topology string, ownedBy string, proxy string, size string, dnsDomain string, txStorage string, sslExpiresOn string, replMasterHostExt string, maxscaleConfig string, volumeIops string, attributes string, skipSync string, replMaster string, sysUpdatedBy string, bulkdataPort2 string, sysCreatedOn string, bulkdataPort1 string, activeReplicas string, fqdn string, sslSerial string, installStatus string, name string, customConfig string, sysId string, provider string, macAddress string, replicas string, sysModCount string, monitor string, ipAddress string, maxscaleProxy string, faultCount string, ) *Database`

NewDatabase instantiates a new Database object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *Database) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Database) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Database) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Database) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *Database) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *Database) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *Database) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.


### GetEnablePamAuthentication

`func (o *Database) GetEnablePamAuthentication() string`

GetEnablePamAuthentication returns the EnablePamAuthentication field if non-nil, zero value otherwise.

### GetEnablePamAuthenticationOk

`func (o *Database) GetEnablePamAuthenticationOk() (*string, bool)`

GetEnablePamAuthenticationOk returns a tuple with the EnablePamAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePamAuthentication

`func (o *Database) SetEnablePamAuthentication(v string)`

SetEnablePamAuthentication sets EnablePamAuthentication field to given value.


### GetSysUpdatedOn

`func (o *Database) GetSysUpdatedOn() string`

GetSysUpdatedOn returns the SysUpdatedOn field if non-nil, zero value otherwise.

### GetSysUpdatedOnOk

`func (o *Database) GetSysUpdatedOnOk() (*string, bool)`

GetSysUpdatedOnOk returns a tuple with the SysUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysUpdatedOn

`func (o *Database) SetSysUpdatedOn(v string)`

SetSysUpdatedOn sets SysUpdatedOn field to given value.


### GetNumber

`func (o *Database) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Database) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Database) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetInstanceState

`func (o *Database) GetInstanceState() string`

GetInstanceState returns the InstanceState field if non-nil, zero value otherwise.

### GetInstanceStateOk

`func (o *Database) GetInstanceStateOk() (*string, bool)`

GetInstanceStateOk returns a tuple with the InstanceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceState

`func (o *Database) SetInstanceState(v string)`

SetInstanceState sets InstanceState field to given value.


### GetReadOnlyPort

`func (o *Database) GetReadOnlyPort() string`

GetReadOnlyPort returns the ReadOnlyPort field if non-nil, zero value otherwise.

### GetReadOnlyPortOk

`func (o *Database) GetReadOnlyPortOk() (*string, bool)`

GetReadOnlyPortOk returns a tuple with the ReadOnlyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyPort

`func (o *Database) SetReadOnlyPort(v string)`

SetReadOnlyPort sets ReadOnlyPort field to given value.


### GetReadWritePort

`func (o *Database) GetReadWritePort() string`

GetReadWritePort returns the ReadWritePort field if non-nil, zero value otherwise.

### GetReadWritePortOk

`func (o *Database) GetReadWritePortOk() (*string, bool)`

GetReadWritePortOk returns a tuple with the ReadWritePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWritePort

`func (o *Database) SetReadWritePort(v string)`

SetReadWritePort sets ReadWritePort field to given value.


### GetId

`func (o *Database) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Database) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Database) SetId(v string)`

SetId sets Id field to given value.


### GetReleaseVersion

`func (o *Database) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *Database) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *Database) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.


### GetGlAccount

`func (o *Database) GetGlAccount() string`

GetGlAccount returns the GlAccount field if non-nil, zero value otherwise.

### GetGlAccountOk

`func (o *Database) GetGlAccountOk() (*string, bool)`

GetGlAccountOk returns a tuple with the GlAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlAccount

`func (o *Database) SetGlAccount(v string)`

SetGlAccount sets GlAccount field to given value.


### GetSysCreatedBy

`func (o *Database) GetSysCreatedBy() string`

GetSysCreatedBy returns the SysCreatedBy field if non-nil, zero value otherwise.

### GetSysCreatedByOk

`func (o *Database) GetSysCreatedByOk() (*string, bool)`

GetSysCreatedByOk returns a tuple with the SysCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysCreatedBy

`func (o *Database) SetSysCreatedBy(v string)`

SetSysCreatedBy sets SysCreatedBy field to given value.


### GetSslCertificate

`func (o *Database) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *Database) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *Database) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.


### GetColumnstoreBucket

`func (o *Database) GetColumnstoreBucket() string`

GetColumnstoreBucket returns the ColumnstoreBucket field if non-nil, zero value otherwise.

### GetColumnstoreBucketOk

`func (o *Database) GetColumnstoreBucketOk() (*string, bool)`

GetColumnstoreBucketOk returns a tuple with the ColumnstoreBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnstoreBucket

`func (o *Database) SetColumnstoreBucket(v string)`

SetColumnstoreBucket sets ColumnstoreBucket field to given value.


### GetTopology

`func (o *Database) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *Database) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *Database) SetTopology(v string)`

SetTopology sets Topology field to given value.


### GetOwnedBy

`func (o *Database) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *Database) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *Database) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetProxy

`func (o *Database) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Database) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Database) SetProxy(v string)`

SetProxy sets Proxy field to given value.


### GetSize

`func (o *Database) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Database) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Database) SetSize(v string)`

SetSize sets Size field to given value.


### GetDnsDomain

`func (o *Database) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *Database) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *Database) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.


### GetTxStorage

`func (o *Database) GetTxStorage() string`

GetTxStorage returns the TxStorage field if non-nil, zero value otherwise.

### GetTxStorageOk

`func (o *Database) GetTxStorageOk() (*string, bool)`

GetTxStorageOk returns a tuple with the TxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStorage

`func (o *Database) SetTxStorage(v string)`

SetTxStorage sets TxStorage field to given value.


### GetSslExpiresOn

`func (o *Database) GetSslExpiresOn() string`

GetSslExpiresOn returns the SslExpiresOn field if non-nil, zero value otherwise.

### GetSslExpiresOnOk

`func (o *Database) GetSslExpiresOnOk() (*string, bool)`

GetSslExpiresOnOk returns a tuple with the SslExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslExpiresOn

`func (o *Database) SetSslExpiresOn(v string)`

SetSslExpiresOn sets SslExpiresOn field to given value.


### GetReplMasterHostExt

`func (o *Database) GetReplMasterHostExt() string`

GetReplMasterHostExt returns the ReplMasterHostExt field if non-nil, zero value otherwise.

### GetReplMasterHostExtOk

`func (o *Database) GetReplMasterHostExtOk() (*string, bool)`

GetReplMasterHostExtOk returns a tuple with the ReplMasterHostExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplMasterHostExt

`func (o *Database) SetReplMasterHostExt(v string)`

SetReplMasterHostExt sets ReplMasterHostExt field to given value.


### GetMaxscaleConfig

`func (o *Database) GetMaxscaleConfig() string`

GetMaxscaleConfig returns the MaxscaleConfig field if non-nil, zero value otherwise.

### GetMaxscaleConfigOk

`func (o *Database) GetMaxscaleConfigOk() (*string, bool)`

GetMaxscaleConfigOk returns a tuple with the MaxscaleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxscaleConfig

`func (o *Database) SetMaxscaleConfig(v string)`

SetMaxscaleConfig sets MaxscaleConfig field to given value.


### GetVolumeIops

`func (o *Database) GetVolumeIops() string`

GetVolumeIops returns the VolumeIops field if non-nil, zero value otherwise.

### GetVolumeIopsOk

`func (o *Database) GetVolumeIopsOk() (*string, bool)`

GetVolumeIopsOk returns a tuple with the VolumeIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIops

`func (o *Database) SetVolumeIops(v string)`

SetVolumeIops sets VolumeIops field to given value.


### GetAttributes

`func (o *Database) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Database) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Database) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.


### GetReplicationStatus

`func (o *Database) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *Database) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *Database) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *Database) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetSkipSync

`func (o *Database) GetSkipSync() string`

GetSkipSync returns the SkipSync field if non-nil, zero value otherwise.

### GetSkipSyncOk

`func (o *Database) GetSkipSyncOk() (*string, bool)`

GetSkipSyncOk returns a tuple with the SkipSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSync

`func (o *Database) SetSkipSync(v string)`

SetSkipSync sets SkipSync field to given value.


### GetReplicationType

`func (o *Database) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *Database) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *Database) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.

### HasReplicationType

`func (o *Database) HasReplicationType() bool`

HasReplicationType returns a boolean if a field has been set.

### GetReplMaster

`func (o *Database) GetReplMaster() string`

GetReplMaster returns the ReplMaster field if non-nil, zero value otherwise.

### GetReplMasterOk

`func (o *Database) GetReplMasterOk() (*string, bool)`

GetReplMasterOk returns a tuple with the ReplMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplMaster

`func (o *Database) SetReplMaster(v string)`

SetReplMaster sets ReplMaster field to given value.


### GetSysUpdatedBy

`func (o *Database) GetSysUpdatedBy() string`

GetSysUpdatedBy returns the SysUpdatedBy field if non-nil, zero value otherwise.

### GetSysUpdatedByOk

`func (o *Database) GetSysUpdatedByOk() (*string, bool)`

GetSysUpdatedByOk returns a tuple with the SysUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysUpdatedBy

`func (o *Database) SetSysUpdatedBy(v string)`

SetSysUpdatedBy sets SysUpdatedBy field to given value.


### GetBulkdataPort2

`func (o *Database) GetBulkdataPort2() string`

GetBulkdataPort2 returns the BulkdataPort2 field if non-nil, zero value otherwise.

### GetBulkdataPort2Ok

`func (o *Database) GetBulkdataPort2Ok() (*string, bool)`

GetBulkdataPort2Ok returns a tuple with the BulkdataPort2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkdataPort2

`func (o *Database) SetBulkdataPort2(v string)`

SetBulkdataPort2 sets BulkdataPort2 field to given value.


### GetSysCreatedOn

`func (o *Database) GetSysCreatedOn() string`

GetSysCreatedOn returns the SysCreatedOn field if non-nil, zero value otherwise.

### GetSysCreatedOnOk

`func (o *Database) GetSysCreatedOnOk() (*string, bool)`

GetSysCreatedOnOk returns a tuple with the SysCreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysCreatedOn

`func (o *Database) SetSysCreatedOn(v string)`

SetSysCreatedOn sets SysCreatedOn field to given value.


### GetBulkdataPort1

`func (o *Database) GetBulkdataPort1() string`

GetBulkdataPort1 returns the BulkdataPort1 field if non-nil, zero value otherwise.

### GetBulkdataPort1Ok

`func (o *Database) GetBulkdataPort1Ok() (*string, bool)`

GetBulkdataPort1Ok returns a tuple with the BulkdataPort1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkdataPort1

`func (o *Database) SetBulkdataPort1(v string)`

SetBulkdataPort1 sets BulkdataPort1 field to given value.


### GetActiveReplicas

`func (o *Database) GetActiveReplicas() string`

GetActiveReplicas returns the ActiveReplicas field if non-nil, zero value otherwise.

### GetActiveReplicasOk

`func (o *Database) GetActiveReplicasOk() (*string, bool)`

GetActiveReplicasOk returns a tuple with the ActiveReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveReplicas

`func (o *Database) SetActiveReplicas(v string)`

SetActiveReplicas sets ActiveReplicas field to given value.


### GetFqdn

`func (o *Database) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Database) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Database) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetSslSerial

`func (o *Database) GetSslSerial() string`

GetSslSerial returns the SslSerial field if non-nil, zero value otherwise.

### GetSslSerialOk

`func (o *Database) GetSslSerialOk() (*string, bool)`

GetSslSerialOk returns a tuple with the SslSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSerial

`func (o *Database) SetSslSerial(v string)`

SetSslSerial sets SslSerial field to given value.


### GetInstallStatus

`func (o *Database) GetInstallStatus() string`

GetInstallStatus returns the InstallStatus field if non-nil, zero value otherwise.

### GetInstallStatusOk

`func (o *Database) GetInstallStatusOk() (*string, bool)`

GetInstallStatusOk returns a tuple with the InstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStatus

`func (o *Database) SetInstallStatus(v string)`

SetInstallStatus sets InstallStatus field to given value.


### GetName

`func (o *Database) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Database) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Database) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *Database) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Database) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Database) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Database) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReplRegion

`func (o *Database) GetReplRegion() string`

GetReplRegion returns the ReplRegion field if non-nil, zero value otherwise.

### GetReplRegionOk

`func (o *Database) GetReplRegionOk() (*string, bool)`

GetReplRegionOk returns a tuple with the ReplRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplRegion

`func (o *Database) SetReplRegion(v string)`

SetReplRegion sets ReplRegion field to given value.

### HasReplRegion

`func (o *Database) HasReplRegion() bool`

HasReplRegion returns a boolean if a field has been set.

### GetCustomConfig

`func (o *Database) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *Database) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *Database) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.


### GetSysId

`func (o *Database) GetSysId() string`

GetSysId returns the SysId field if non-nil, zero value otherwise.

### GetSysIdOk

`func (o *Database) GetSysIdOk() (*string, bool)`

GetSysIdOk returns a tuple with the SysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysId

`func (o *Database) SetSysId(v string)`

SetSysId sets SysId field to given value.


### GetProvider

`func (o *Database) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Database) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Database) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetMacAddress

`func (o *Database) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Database) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Database) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetReplicas

`func (o *Database) GetReplicas() string`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Database) GetReplicasOk() (*string, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Database) SetReplicas(v string)`

SetReplicas sets Replicas field to given value.


### GetSysModCount

`func (o *Database) GetSysModCount() string`

GetSysModCount returns the SysModCount field if non-nil, zero value otherwise.

### GetSysModCountOk

`func (o *Database) GetSysModCountOk() (*string, bool)`

GetSysModCountOk returns a tuple with the SysModCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysModCount

`func (o *Database) SetSysModCount(v string)`

SetSysModCount sets SysModCount field to given value.


### GetMonitor

`func (o *Database) GetMonitor() string`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *Database) GetMonitorOk() (*string, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *Database) SetMonitor(v string)`

SetMonitor sets Monitor field to given value.


### GetIpAddress

`func (o *Database) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Database) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Database) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetMaxscaleProxy

`func (o *Database) GetMaxscaleProxy() string`

GetMaxscaleProxy returns the MaxscaleProxy field if non-nil, zero value otherwise.

### GetMaxscaleProxyOk

`func (o *Database) GetMaxscaleProxyOk() (*string, bool)`

GetMaxscaleProxyOk returns a tuple with the MaxscaleProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxscaleProxy

`func (o *Database) SetMaxscaleProxy(v string)`

SetMaxscaleProxy sets MaxscaleProxy field to given value.


### GetFaultCount

`func (o *Database) GetFaultCount() string`

GetFaultCount returns the FaultCount field if non-nil, zero value otherwise.

### GetFaultCountOk

`func (o *Database) GetFaultCountOk() (*string, bool)`

GetFaultCountOk returns a tuple with the FaultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultCount

`func (o *Database) SetFaultCount(v string)`

SetFaultCount sets FaultCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


