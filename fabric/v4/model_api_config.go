/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.4
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Configuration for API based Integration for Service Profile
type ApiConfig struct {
	// Setting indicating whether the API is available (true) or not (false).
	ApiAvailable  bool   `json:"apiAvailable,omitempty"`
	IntegrationId string `json:"integrationId,omitempty"`
	// Setting indicating that the port is managed by Equinix (true) or not (false).
	EquinixManagedPort bool `json:"equinixManagedPort,omitempty"`
	// Setting indicating that the VLAN is managed by Equinix (true) or not (false).
	EquinixManagedVlan bool `json:"equinixManagedVlan,omitempty"`
	// Setting showing that oversubscription support is available (true) or not (false). The default is false. Oversubscription is the sale of more than the available network bandwidth. This practice is common and legitimate. After all, many customers use less bandwidth than they've purchased. And network users don't consume bandwidth all at the same time. The leftover bandwidth can be sold to other customers. When demand surges, operational and engineering resources can be shifted to accommodate the load.
	AllowOverSubscription bool `json:"allowOverSubscription,omitempty"`
	// A cap on oversubscription.
	OverSubscriptionLimit int32 `json:"overSubscriptionLimit,omitempty"`
	BandwidthFromApi      bool  `json:"bandwidthFromApi,omitempty"`
}