/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.3.52
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Connection link protocol Configuration
type SimplifiedLinkProtocol struct {
	Type_ *LinkProtocolType `json:"type,omitempty"`
	// vlanTag value specified for DOT1Q connections
	VlanTag int32 `json:"vlanTag,omitempty"`
	// vlanSTag value specified for QINQ connections
	VlanSTag int32 `json:"vlanSTag,omitempty"`
	// vlanCTag value specified for QINQ connections
	VlanCTag int32 `json:"vlanCTag,omitempty"`
	Unit     int32 `json:"unit,omitempty"`
	Vni      int32 `json:"vni,omitempty"`
	IntUnit  int32 `json:"intUnit,omitempty"`
}
