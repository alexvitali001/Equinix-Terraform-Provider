/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.7
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// IPv4 or IPv6 specific configuration
type LinkProtocolIpv4Ipv6Config struct {
	// Link subnet prefix
	LinkPrefix string `json:"linkPrefix,omitempty"`
	// Prefix datatype when linkPrefix not specified
	LocalIfaceIp string `json:"localIfaceIp,omitempty"`
	// Equinix-side link interface address
	RemoteIfaceIp string `json:"remoteIfaceIp,omitempty"`
}
