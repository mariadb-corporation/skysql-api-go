// Package skysql provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package skysql

const (
	HTTPBearerScopes = "HTTPBearer.Scopes"
)

// IP Address that has been added to the database network allowlist
type AllowlistIPAddress struct {
	Comment   *string `json:"comment,omitempty"`
	Database  string  `json:"database"`
	IpAddress string  `json:"ip_address"`
}

// GET Configuration Response
type ConfigurationResp struct {
	ConfigurationVersions *[]ConfigurationVersionResp `json:"configuration_versions,omitempty"`
	Name                  string                      `json:"name"`
	Number                string                      `json:"number"`
	Public                string                      `json:"public"`
	SysId                 string                      `json:"sys_id"`
	Topology              string                      `json:"topology"`
}

// GET Configuration Response nested configuration version
type ConfigurationVersionResp struct {
	ConfigJson     map[string]interface{} `json:"config_json"`
	CurrentVersion string                 `json:"current_version"`
	Version        string                 `json:"version"`
}

// Request body to create a new MariaDB Configuration
type CreateConfigurationRequest struct {
	ConfigJson *map[string]interface{} `json:"config_json,omitempty"`
	Name       string                  `json:"name"`
	Topology   string                  `json:"topology"`
}

// Update Configuration Response
type CreateConfigurationResp struct {
	// GET Configuration Response nested configuration version
	ConfigurationVersion ConfigurationVersionResp `json:"configuration_version"`
	Name                 string                   `json:"name"`
	Number               string                   `json:"number"`
	Public               string                   `json:"public"`
	SysId                string                   `json:"sys_id"`
	Topology             string                   `json:"topology"`
}

// MariaDB cluster deployed by SkySQL
type Database struct {
	ActiveReplicas          string  `json:"active_replicas"`
	Attributes              string  `json:"attributes"`
	BulkdataPort1           string  `json:"bulkdata_port_1"`
	BulkdataPort2           string  `json:"bulkdata_port_2"`
	Cluster                 *string `json:"cluster,omitempty"`
	ColumnstoreBucket       string  `json:"columnstore_bucket"`
	CustomConfig            string  `json:"custom_config"`
	DnsDomain               string  `json:"dns_domain"`
	EnablePamAuthentication string  `json:"enable_pam_authentication"`
	FaultCount              string  `json:"fault_count"`
	Fqdn                    string  `json:"fqdn"`
	GlAccount               string  `json:"gl_account"`
	Id                      string  `json:"id"`
	InstallStatus           string  `json:"install_status"`
	InstanceState           string  `json:"instance_state"`
	IpAddress               string  `json:"ip_address"`
	MacAddress              string  `json:"mac_address"`
	MaxscaleConfig          string  `json:"maxscale_config"`
	MaxscaleProxy           string  `json:"maxscale_proxy"`
	Monitor                 string  `json:"monitor"`
	Name                    string  `json:"name"`
	Number                  string  `json:"number"`
	OperationalStatus       string  `json:"operational_status"`
	OwnedBy                 string  `json:"owned_by"`
	Provider                string  `json:"provider"`
	Proxy                   string  `json:"proxy"`
	ReadOnlyPort            string  `json:"read_only_port"`
	ReadWritePort           string  `json:"read_write_port"`
	Region                  *string `json:"region,omitempty"`
	ReleaseVersion          string  `json:"release_version"`
	ReplMaster              string  `json:"repl_master"`
	ReplMasterHostExt       string  `json:"repl_master_host_ext"`
	ReplRegion              *string `json:"repl_region,omitempty"`
	Replicas                string  `json:"replicas"`
	ReplicationStatus       *string `json:"replication_status,omitempty"`
	ReplicationType         *string `json:"replication_type,omitempty"`
	Size                    string  `json:"size"`
	SkipSync                string  `json:"skip_sync"`
	SslCertificate          string  `json:"ssl_certificate"`
	SslExpiresOn            string  `json:"ssl_expires_on"`
	SslSerial               string  `json:"ssl_serial"`
	SysCreatedBy            string  `json:"sys_created_by"`
	SysCreatedOn            string  `json:"sys_created_on"`
	SysId                   string  `json:"sys_id"`
	SysModCount             string  `json:"sys_mod_count"`
	SysUpdatedBy            string  `json:"sys_updated_by"`
	SysUpdatedOn            string  `json:"sys_updated_on"`
	Topology                string  `json:"topology"`
	TxStorage               string  `json:"tx_storage"`
	VolumeIops              *string `json:"volume_iops,omitempty"`
	VolumeType              *string `json:"volume_type,omitempty"`
}

// Actions that can be taken on a Database in a Task
type DatabaseActions interface{}

// Response body for a database status
type DatabaseStatus struct {
	Status string `json:"status"`
}

// Request body to update a database
type DatabaseStatusUpdate struct {
	// Actions that can be taken on a Database in a Task
	Action DatabaseActions `json:"action"`
}

// Request body to update a database - currently limited to name only
type DatabaseUpdate struct {
	Name string `json:"name"`
}

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// IP Address representation in SkySQL
type IPAddress struct {
	Comment   *string `json:"comment,omitempty"`
	IpAddress string  `json:"ip_address"`
}

// Generic message body containing error details for failed requests
type Message struct {
	Details string `json:"details"`
}

// Request body to create a new MariaDB cluster deployed by SkySQL
type NewDatabase struct {
	MaxscaleConfig string `json:"maxscale_config"`
	MaxscaleProxy  string `json:"maxscale_proxy"`
	Monitor        string `json:"monitor"`
	Name           string `json:"name"`
	Provider       string `json:"provider"`
	Region         string `json:"region"`
	ReleaseVersion string `json:"release_version"`
	ReplRegion     string `json:"repl_region"`
	Replicas       string `json:"replicas"`
	Size           string `json:"size"`
	Topology       string `json:"topology"`
	TxStorage      string `json:"tx_storage"`
	VolumeIops     string `json:"volume_iops"`
	VolumeType     string `json:"volume_type"`
}

// A database product, e.g. Xpand
type Product struct {
	Active           string  `json:"active"`
	ActiveTopologies string  `json:"active_topologies"`
	DefaultTopology  string  `json:"default_topology"`
	Description      *string `json:"description,omitempty"`
	DisplayName      string  `json:"display_name"`
	Name             string  `json:"name"`
	Order            string  `json:"order"`
	ShortDescription string  `json:"short_description"`
	SysCreatedBy     string  `json:"sys_created_by"`
	SysCreatedOn     string  `json:"sys_created_on"`
	SysId            string  `json:"sys_id"`
	SysModCount      string  `json:"sys_mod_count"`
	SysTags          string  `json:"sys_tags"`
	SysUpdatedBy     string  `json:"sys_updated_by"`
	SysUpdatedOn     string  `json:"sys_updated_on"`
}

// Cloud provider, e.g. AWS or GCP
type Provider struct {
	Active       string `json:"active"`
	IconImage    string `json:"icon_image"`
	LogoImage    string `json:"logo_image"`
	Name         string `json:"name"`
	Products     string `json:"products"`
	SysCreatedBy string `json:"sys_created_by"`
	SysCreatedOn string `json:"sys_created_on"`
	SysId        string `json:"sys_id"`
	SysModCount  string `json:"sys_mod_count"`
	SysTags      string `json:"sys_tags"`
	SysUpdatedBy string `json:"sys_updated_by"`
	SysUpdatedOn string `json:"sys_updated_on"`
	Topologies   string `json:"topologies"`
	Value        string `json:"value"`
}

// A quota progress response
type QuotaProgress struct {
	Description string  `json:"description"`
	Limit       int     `json:"limit"`
	Name        string  `json:"name"`
	Provider    *string `json:"provider,omitempty"`
	Region      *string `json:"region,omitempty"`
	Remaining   *int    `json:"remaining,omitempty"`
	Used        *int    `json:"used,omitempty"`
}

// Geographic region, as defined by the providers
type Region struct {
	Active       string  `json:"active"`
	ApiHandle    string  `json:"api_handle"`
	Default      string  `json:"default"`
	Location     string  `json:"location"`
	Name         string  `json:"name"`
	NodeType     *string `json:"node_type,omitempty"`
	Provider     string  `json:"provider"`
	Region       string  `json:"region"`
	SysCreatedBy string  `json:"sys_created_by"`
	SysCreatedOn string  `json:"sys_created_on"`
	SysId        string  `json:"sys_id"`
	SysModCount  string  `json:"sys_mod_count"`
	SysTags      string  `json:"sys_tags"`
	SysUpdatedBy string  `json:"sys_updated_by"`
	SysUpdatedOn string  `json:"sys_updated_on"`
}

// Node size, as defined by the providers
type Size struct {
	Active       string `json:"active"`
	ApiHandle    string `json:"api_handle"`
	Component    string `json:"component"`
	Cpu          string `json:"cpu"`
	Name         string `json:"name"`
	NodePool     string `json:"node_pool"`
	Product      string `json:"product"`
	Provider     string `json:"provider"`
	Ram          string `json:"ram"`
	Sequence     string `json:"sequence"`
	SysCreatedBy string `json:"sys_created_by"`
	SysCreatedOn string `json:"sys_created_on"`
	SysId        string `json:"sys_id"`
	SysModCount  string `json:"sys_mod_count"`
	SysTags      string `json:"sys_tags"`
	SysUpdatedBy string `json:"sys_updated_by"`
	SysUpdatedOn string `json:"sys_updated_on"`
	Tier         string `json:"tier"`
	Value        string `json:"value"`
	Visibility   string `json:"visibility"`
}

// Availability tier, e.g. dedicated or shared tenancy
type Tier struct {
	Active       string `json:"active"`
	Name         string `json:"name"`
	Paid         string `json:"paid"`
	SysCreatedBy string `json:"sys_created_by"`
	SysCreatedOn string `json:"sys_created_on"`
	SysId        string `json:"sys_id"`
	SysModCount  string `json:"sys_mod_count"`
	SysTags      string `json:"sys_tags"`
	SysUpdatedBy string `json:"sys_updated_by"`
	SysUpdatedOn string `json:"sys_updated_on"`
}

// Cluster topology valid for a particular product, e.g. Standalone or MaxScale
type Topology struct {
	Active                 string `json:"active"`
	ApiHandle              string `json:"api_handle"`
	DisplayName            string `json:"display_name"`
	HasClustrixFrontend    string `json:"has_clustrix_frontend"`
	HasReplicas            string `json:"has_replicas"`
	MariadbCronbackEnabled string `json:"mariadb_cronback_enabled"`
	MaxscaleSupported      string `json:"maxscale_supported"`
	Name                   string `json:"name"`
	OperatorReplicaMinimum string `json:"operator_replica_minimum"`
	Paid                   string `json:"paid"`
	Product                string `json:"product"`
	ReplicaLabel           string `json:"replica_label"`
	ReplicaOptions         string `json:"replica_options"`
	SysCreatedBy           string `json:"sys_created_by"`
	SysCreatedOn           string `json:"sys_created_on"`
	SysId                  string `json:"sys_id"`
	SysModCount            string `json:"sys_mod_count"`
	SysTags                string `json:"sys_tags"`
	SysUpdatedBy           string `json:"sys_updated_by"`
	SysUpdatedOn           string `json:"sys_updated_on"`
	Value                  string `json:"value"`
}

// Request body to update a configuration
type UpdateConfigurationRequest struct {
	ConfigJson *map[string]interface{} `json:"config_json,omitempty"`
	Name       *string                 `json:"name,omitempty"`
}

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

// Database version, e.g. 10.4 or 10.5
type Version struct {
	DisplayName       string  `json:"display_name"`
	EnterpriseVersion *string `json:"enterprise_version,omitempty"`
	Name              string  `json:"name"`
	ParentRelease     string  `json:"parent_release"`
	Product           string  `json:"product"`
	Provider          string  `json:"provider"`
	Public            string  `json:"public"`
	ReleaseDate       string  `json:"release_date"`
	ReleaseNotesUrl   string  `json:"release_notes_url"`
	SysCreatedBy      string  `json:"sys_created_by"`
	SysCreatedOn      string  `json:"sys_created_on"`
	SysId             string  `json:"sys_id"`
	SysModCount       string  `json:"sys_mod_count"`
	SysTags           string  `json:"sys_tags"`
	SysUpdatedBy      string  `json:"sys_updated_by"`
	SysUpdatedOn      string  `json:"sys_updated_on"`
	Type              string  `json:"type"`
}

// ListConfigurationsParams defines parameters for ListConfigurations.
type ListConfigurationsParams struct {
	Limit *int `json:"limit,omitempty"`
}

// CreateConfigurationJSONBody defines parameters for CreateConfiguration.
type CreateConfigurationJSONBody CreateConfigurationRequest

// UpdateConfigurationJSONBody defines parameters for UpdateConfiguration.
type UpdateConfigurationJSONBody UpdateConfigurationRequest

// ListDatabasesParams defines parameters for ListDatabases.
type ListDatabasesParams struct {
	Limit *int `json:"limit,omitempty"`
}

// CreateDatabaseJSONBody defines parameters for CreateDatabase.
type CreateDatabaseJSONBody NewDatabase

// UpdateDatabaseJSONBody defines parameters for UpdateDatabase.
type UpdateDatabaseJSONBody DatabaseUpdate

// RemoveAllowedAddressParams defines parameters for RemoveAllowedAddress.
type RemoveAllowedAddressParams struct {
	Address *string `json:"address,omitempty"`
}

// ListAllowedAddressesParams defines parameters for ListAllowedAddresses.
type ListAllowedAddressesParams struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// AddAllowedAddressJSONBody defines parameters for AddAllowedAddress.
type AddAllowedAddressJSONBody IPAddress

// UpdateStatusJSONBody defines parameters for UpdateStatus.
type UpdateStatusJSONBody DatabaseStatusUpdate

// ReadProductsParams defines parameters for ReadProducts.
type ReadProductsParams struct {
	Limit *int `json:"limit,omitempty"`
}

// ReadProvidersParams defines parameters for ReadProviders.
type ReadProvidersParams struct {
	Limit *int `json:"limit,omitempty"`
}

// ReadRegionsParams defines parameters for ReadRegions.
type ReadRegionsParams struct {
	Provider string `json:"provider"`
	Limit    *int   `json:"limit,omitempty"`
}

// ReadSizesParams defines parameters for ReadSizes.
type ReadSizesParams struct {
	Product  string `json:"product"`
	Provider string `json:"provider"`
	Limit    *int   `json:"limit,omitempty"`
}

// ReadTiersParams defines parameters for ReadTiers.
type ReadTiersParams struct {
	Limit *int `json:"limit,omitempty"`
}

// ReadTopologiesParams defines parameters for ReadTopologies.
type ReadTopologiesParams struct {
	Product string `json:"product"`
	Limit   *int   `json:"limit,omitempty"`
}

// ReadVersionsParams defines parameters for ReadVersions.
type ReadVersionsParams struct {
	Limit *int `json:"limit,omitempty"`
}

// CreateConfigurationJSONRequestBody defines body for CreateConfiguration for application/json ContentType.
type CreateConfigurationJSONRequestBody CreateConfigurationJSONBody

// UpdateConfigurationJSONRequestBody defines body for UpdateConfiguration for application/json ContentType.
type UpdateConfigurationJSONRequestBody UpdateConfigurationJSONBody

// CreateDatabaseJSONRequestBody defines body for CreateDatabase for application/json ContentType.
type CreateDatabaseJSONRequestBody CreateDatabaseJSONBody

// UpdateDatabaseJSONRequestBody defines body for UpdateDatabase for application/json ContentType.
type UpdateDatabaseJSONRequestBody UpdateDatabaseJSONBody

// AddAllowedAddressJSONRequestBody defines body for AddAllowedAddress for application/json ContentType.
type AddAllowedAddressJSONRequestBody AddAllowedAddressJSONBody

// UpdateStatusJSONRequestBody defines body for UpdateStatus for application/json ContentType.
type UpdateStatusJSONRequestBody UpdateStatusJSONBody

