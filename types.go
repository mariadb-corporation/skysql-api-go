// Package skysql provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package skysql

const (
	HTTPBearerScopes = "HTTPBearer.Scopes"
)

// Defines values for AllowlistStatuses.
const (
	AllowlistStatusesEnforcing AllowlistStatuses = "Enforcing"

	AllowlistStatusesUpdating AllowlistStatuses = "Updating"
)

// Defines values for ServiceInSslTls.
const (
	ServiceInSslTlsDisabled ServiceInSslTls = "Disabled"

	ServiceInSslTlsEmpty ServiceInSslTls = ""

	ServiceInSslTlsEnabled ServiceInSslTls = "Enabled"
)

// Defines values for ServiceInTier.
const (
	ServiceInTierFoundation ServiceInTier = "Foundation"

	ServiceInTierPower ServiceInTier = "Power"
)

// Defines values for ServiceInTopology.
const (
	ServiceInTopologyDistributedTransactions ServiceInTopology = "Distributed Transactions"

	ServiceInTopologyMultiNodeAnalytics ServiceInTopology = "Multi-Node Analytics"

	ServiceInTopologyReplicatedTransactions ServiceInTopology = "Replicated Transactions"

	ServiceInTopologySingleNodeAnalytics ServiceInTopology = "Single Node Analytics"

	ServiceInTopologySingleNodeTransactions ServiceInTopology = "Single Node Transactions"
)

// Defines values for ServiceOutSslTls.
const (
	ServiceOutSslTlsDisabled ServiceOutSslTls = "Disabled"

	ServiceOutSslTlsEmpty ServiceOutSslTls = ""

	ServiceOutSslTlsEnabled ServiceOutSslTls = "Enabled"
)

// Defines values for SnowProviders.
const (
	SnowProvidersAWS SnowProviders = "AWS"

	SnowProvidersAmazon SnowProviders = "Amazon"

	SnowProvidersAmazonAWS SnowProviders = "Amazon AWS"

	SnowProvidersAws SnowProviders = "aws"

	SnowProvidersGCP SnowProviders = "GCP"

	SnowProvidersGcp SnowProviders = "gcp"

	SnowProvidersGoogle SnowProviders = "Google"

	SnowProvidersGoogleCloud SnowProviders = "Google Cloud"
)

// IP Address that has been added to the services network allowlist
type AllowlistIPAddress struct {
	Comment   *string `json:"comment,omitempty"`
	IpAddress string  `json:"ip_address"`
	ServiceId string  `json:"service_id"`
}

// Status of the service allowlist to indicate if address changes are provisioning
type AllowlistStatus struct {
	// Possible statuses for the allowlist
	Status AllowlistStatuses `json:"status"`
}

// Possible statuses for the allowlist
type AllowlistStatuses string

// GET Configuration Response
type ConfigurationResp struct {
	ConfigurationVersions *[]ConfigurationVersionResp `json:"configuration_versions,omitempty"`
	Name                  string                      `json:"name"`
	Number                string                      `json:"number"`
	Public                string                      `json:"public"`
	Topology              string                      `json:"topology"`
}

// Configuration Response Base Model
type ConfigurationResponseBaseModel struct {
	Name     string `json:"name"`
	Number   string `json:"number"`
	Public   string `json:"public"`
	Topology string `json:"topology"`
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
	Topology             string                   `json:"topology"`
}

// A credential issued for initial connection to a database,
// intended to be immediately replaced by the user
type DefaultCredentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
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
	Detail string `json:"detail"`
}

// Cloud provider, e.g. AWS or GCP
type Provider struct {
	Active       *string `json:"active,omitempty"`
	CreatedOn    *string `json:"created_on,omitempty"`
	IconImage    *string `json:"icon_image,omitempty"`
	LogoImage    *string `json:"logo_image,omitempty"`
	ModCount     *string `json:"mod_count,omitempty"`
	Name         string  `json:"name"`
	ServiceTypes *string `json:"service_types,omitempty"`
	Topologies   *string `json:"topologies,omitempty"`
	UpdatedOn    string  `json:"updated_on"`
	Value        string  `json:"value"`
}

// A quota progress response
type QuotaProgress struct {
	Description *string `json:"description,omitempty"`
	Limit       int     `json:"limit"`
	Name        string  `json:"name"`
	Provider    *string `json:"provider,omitempty"`
	Region      *string `json:"region,omitempty"`
	Remaining   *int    `json:"remaining,omitempty"`
	Used        *int    `json:"used,omitempty"`
}

// Geographic region, as defined by the providers
type Region struct {
	CreatedOn *string `json:"created_on,omitempty"`
	Default   *string `json:"default,omitempty"`
	Name      string  `json:"name"`
	NodeType  *string `json:"node_type,omitempty"`
	Provider  string  `json:"provider"`
	Region    string  `json:"region"`
}

// Actions that can be taken on a services in a Task
type ServiceActions interface{}

// Request body to create a new MariaDB services deployed by SkySQL
type ServiceIn struct {
	MaxscaleConfig *string `json:"maxscale_config,omitempty"`
	MaxscaleProxy  *string `json:"maxscale_proxy,omitempty"`
	Monitor        *string `json:"monitor,omitempty"`
	Name           string  `json:"name"`

	// Providers configured in snow
	Provider       SnowProviders     `json:"provider"`
	Region         string            `json:"region"`
	ReleaseVersion string            `json:"release_version"`
	Replicas       string            `json:"replicas"`
	Size           string            `json:"size"`
	SslTls         *ServiceInSslTls  `json:"ssl_tls,omitempty"`
	Tier           ServiceInTier     `json:"tier"`
	Topology       ServiceInTopology `json:"topology"`
	TxStorage      string            `json:"tx_storage"`
	VolumeIops     *string           `json:"volume_iops,omitempty"`
}

// ServiceInSslTls defines model for ServiceIn.SslTls.
type ServiceInSslTls string

// ServiceInTier defines model for ServiceIn.Tier.
type ServiceInTier string

// ServiceInTopology defines model for ServiceIn.Topology.
type ServiceInTopology string

// Base class for Service responses
type ServiceOut struct {
	CreatedOn      *string           `json:"created_on,omitempty"`
	CustomConfig   *string           `json:"custom_config,omitempty"`
	Fqdn           *string           `json:"fqdn,omitempty"`
	Id             *string           `json:"id,omitempty"`
	InstallStatus  *string           `json:"install_status,omitempty"`
	IpAddress      *string           `json:"ip_address,omitempty"`
	MaxscaleConfig *string           `json:"maxscale_config,omitempty"`
	MaxscaleProxy  *string           `json:"maxscale_proxy,omitempty"`
	Monitor        *string           `json:"monitor,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Number         *string           `json:"number,omitempty"`
	OwnedBy        *string           `json:"owned_by,omitempty"`
	Provider       *string           `json:"provider,omitempty"`
	ReadOnlyPort   *string           `json:"read_only_port,omitempty"`
	ReadWritePort  *string           `json:"read_write_port,omitempty"`
	Region         *string           `json:"region,omitempty"`
	ReleaseVersion *string           `json:"release_version,omitempty"`
	Replicas       *string           `json:"replicas,omitempty"`
	Size           *string           `json:"size,omitempty"`
	SslTls         *ServiceOutSslTls `json:"ssl_tls,omitempty"`
	Tier           *string           `json:"tier,omitempty"`
	Topology       *string           `json:"topology,omitempty"`
	TxStorage      *string           `json:"tx_storage,omitempty"`
	UpdatedOn      *string           `json:"updated_on,omitempty"`
	VolumeIops     *string           `json:"volume_iops,omitempty"`
}

// ServiceOutSslTls defines model for ServiceOut.SslTls.
type ServiceOutSslTls string

// Response body for a services status
type ServiceStatus struct {
	Status string `json:"status"`
}

// Request body to update a service
type ServiceStatusUpdate struct {
	// Actions that can be taken on a services in a Task
	Action ServiceActions `json:"action"`
}

// A database service type, e.g. Distributed Transactions
type ServiceType struct {
	ActiveTopologies *string `json:"active_topologies,omitempty"`
	CreatedOn        *string `json:"created_on,omitempty"`
	DefaultTopology  *string `json:"default_topology,omitempty"`
	Description      *string `json:"description,omitempty"`
	ModCount         *string `json:"mod_count,omitempty"`
	Name             *string `json:"name,omitempty"`
	Order            *string `json:"order,omitempty"`
	ShortDescription *string `json:"short_description,omitempty"`
	UpdatedOn        *string `json:"updated_on,omitempty"`
}

// Request body to update a services - currently limited to name only
type ServiceUpdate struct {
	Name      *string `json:"name,omitempty"`
	TxStorage *string `json:"tx_storage,omitempty"`
}

// Node size, as defined by the providers
type Size struct {
	Cpu         *string `json:"cpu,omitempty"`
	CreatedOn   *string `json:"created_on,omitempty"`
	Name        string  `json:"name"`
	Provider    string  `json:"provider"`
	Ram         *string `json:"ram,omitempty"`
	ServiceType string  `json:"service_type"`
	Tier        string  `json:"tier"`
	Visibility  string  `json:"visibility"`
}

// Providers configured in snow
type SnowProviders string

// Availability tier, e.g. Power, Foundation, etc.
type Tier struct {
	Name *string `json:"name,omitempty"`
	Paid *string `json:"paid,omitempty"`
}

// Cluster topology valid for a particular product, e.g. Standalone or MaxScale
type Topology struct {
	Active            *string `json:"active,omitempty"`
	MaxscaleSupported *string `json:"maxscale_supported,omitempty"`
	Name              *string `json:"name,omitempty"`
	ReplicaLabel      *string `json:"replica_label,omitempty"`
	ReplicaOptions    *string `json:"replica_options,omitempty"`
	ServiceType       string  `json:"service_type"`
}

// Message details containing unmet expectations from failed specs
type UnmetExpectationDetail struct {
	// Message body with details regarding failed specifications
	Detail UnmetExpectationMessage `json:"detail"`
}

// Message body with details regarding failed specifications
type UnmetExpectationMessage struct {
	Description       string  `json:"description"`
	UnmetExpectations *string `json:"unmet_expectations,omitempty"`
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
	CreatedOn         *string `json:"created_on,omitempty"`
	EnterpriseVersion *string `json:"enterprise_version,omitempty"`
	ModCount          *string `json:"mod_count,omitempty"`
	Name              *string `json:"name,omitempty"`
	ParentRelease     *string `json:"parent_release,omitempty"`
	Provider          *string `json:"provider,omitempty"`
	Public            *string `json:"public,omitempty"`
	ReleaseDate       *string `json:"release_date,omitempty"`
	ReleaseNotesUrl   *string `json:"release_notes_url,omitempty"`
	ServiceType       *string `json:"service_type,omitempty"`
	Type              *string `json:"type,omitempty"`
	UpdatedOn         *string `json:"updated_on,omitempty"`
}

// ListConfigurationsParams defines parameters for ListConfigurations.
type ListConfigurationsParams struct {
	Limit *int `json:"limit,omitempty"`
}

// CreateConfigurationJSONBody defines parameters for CreateConfiguration.
type CreateConfigurationJSONBody CreateConfigurationRequest

// UpdateConfigurationJSONBody defines parameters for UpdateConfiguration.
type UpdateConfigurationJSONBody UpdateConfigurationRequest

// ReadProvidersParams defines parameters for ReadProviders.
type ReadProvidersParams struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// ReadRegionsParams defines parameters for ReadRegions.
type ReadRegionsParams struct {
	Provider SnowProviders `json:"provider"`
	Limit    *int          `json:"limit,omitempty"`
	Offset   *int          `json:"offset,omitempty"`
}

// ReadServiceTypesParams defines parameters for ReadServiceTypes.
type ReadServiceTypesParams struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// ReadSizesParams defines parameters for ReadSizes.
type ReadSizesParams struct {
	Region   string                  `json:"region"`
	Topology ReadSizesParamsTopology `json:"topology"`
	Provider SnowProviders           `json:"provider"`
	Tier     ReadSizesParamsTier     `json:"tier"`
	Limit    *int                    `json:"limit,omitempty"`
	Offset   *int                    `json:"offset,omitempty"`
}

// ReadSizesParamsTopology defines parameters for ReadSizes.
type ReadSizesParamsTopology string

// ReadSizesParamsTier defines parameters for ReadSizes.
type ReadSizesParamsTier string

// ReadTiersParams defines parameters for ReadTiers.
type ReadTiersParams struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// ReadTopologiesParams defines parameters for ReadTopologies.
type ReadTopologiesParams struct {
	ServiceType ReadTopologiesParamsServiceType `json:"service_type"`
	Limit       *int                            `json:"limit,omitempty"`
	Offset      *int                            `json:"offset,omitempty"`
}

// ReadTopologiesParamsServiceType defines parameters for ReadTopologies.
type ReadTopologiesParamsServiceType string

// ReadVersionsParams defines parameters for ReadVersions.
type ReadVersionsParams struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// ListServicesParams defines parameters for ListServices.
type ListServicesParams struct {
	Name   *string `json:"name,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
	Offset *int    `json:"offset,omitempty"`
}

// CreateServiceJSONBody defines parameters for CreateService.
type CreateServiceJSONBody ServiceIn

// UpdateServiceJSONBody defines parameters for UpdateService.
type UpdateServiceJSONBody ServiceUpdate

// RemoveAllowedAddressParams defines parameters for RemoveAllowedAddress.
type RemoveAllowedAddressParams struct {
	Address string `json:"address"`
}

// ListAllowedAddressesParams defines parameters for ListAllowedAddresses.
type ListAllowedAddressesParams struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

// AddAllowedAddressJSONBody defines parameters for AddAllowedAddress.
type AddAllowedAddressJSONBody IPAddress

// UpdateStatusJSONBody defines parameters for UpdateStatus.
type UpdateStatusJSONBody ServiceStatusUpdate

// CreateConfigurationJSONRequestBody defines body for CreateConfiguration for application/json ContentType.
type CreateConfigurationJSONRequestBody CreateConfigurationJSONBody

// UpdateConfigurationJSONRequestBody defines body for UpdateConfiguration for application/json ContentType.
type UpdateConfigurationJSONRequestBody UpdateConfigurationJSONBody

// CreateServiceJSONRequestBody defines body for CreateService for application/json ContentType.
type CreateServiceJSONRequestBody CreateServiceJSONBody

// UpdateServiceJSONRequestBody defines body for UpdateService for application/json ContentType.
type UpdateServiceJSONRequestBody UpdateServiceJSONBody

// AddAllowedAddressJSONRequestBody defines body for AddAllowedAddress for application/json ContentType.
type AddAllowedAddressJSONRequestBody AddAllowedAddressJSONBody

// UpdateStatusJSONRequestBody defines body for UpdateStatus for application/json ContentType.
type UpdateStatusJSONRequestBody UpdateStatusJSONBody
