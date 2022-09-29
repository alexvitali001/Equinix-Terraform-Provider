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

// Fabric Gateway Package
type VirtualGatewayPackage struct {
	// Gateway package URI
	Href string `json:"href,omitempty"`
	// Type of Gateway package
	Type_ string `json:"type,omitempty"`
	Code  *Code  `json:"code,omitempty"`
	// Fabric Gateway Package description
	Description string `json:"description,omitempty"`
	// Gateway package BGP IPv4 routes limit
	TotalIPv4RoutesMax int32 `json:"totalIPv4RoutesMax,omitempty"`
	// Gateway package BGP IPv6 routes limit
	TotalIPv6RoutesMax int32 `json:"totalIPv6RoutesMax,omitempty"`
	// Gateway package static IPv4 routes limit
	StaticIPv4RoutesMax int32 `json:"staticIPv4RoutesMax,omitempty"`
	// Gateway package static IPv6 routes limit
	StaticIPv6RoutesMax int32 `json:"staticIPv6RoutesMax,omitempty"`
	// Gateway package NACLs limit
	NaclsMax int32 `json:"naclsMax,omitempty"`
	// Gateway package NACLs rules limit
	NaclRulesMax int32 `json:"naclRulesMax,omitempty"`
	// Gateway package high-available configuration support
	HaSupported bool `json:"haSupported,omitempty"`
	// Gateway package route filter support
	RouteFilterSupported bool `json:"routeFilterSupported,omitempty"`
	// Gateway package NAT supported type
	NatType   string            `json:"natType,omitempty"`
	ChangeLog *PackageChangeLog `json:"changeLog,omitempty"`
}