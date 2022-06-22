/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.3
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Connection link protocol configuration - UNTAGGED
type LinkProtocolUntagged struct {
	Type_ *LinkProtocolType           `json:"type"`
	Ipv4  *LinkProtocolIpv4Ipv6Config `json:"ipv4,omitempty"`
	Ipv6  *LinkProtocolIpv4Ipv6Config `json:"ipv6,omitempty"`
}
