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

// VirtualConnectionPriceAccessPointType : Virtual Connection access point type
type VirtualConnectionPriceAccessPointType string

// List of VirtualConnectionPriceAccessPointType
const (
	VD_VirtualConnectionPriceAccessPointType         VirtualConnectionPriceAccessPointType = "VD"
	SP_VirtualConnectionPriceAccessPointType         VirtualConnectionPriceAccessPointType = "SP"
	COLO_VirtualConnectionPriceAccessPointType       VirtualConnectionPriceAccessPointType = "COLO"
	GW_VirtualConnectionPriceAccessPointType         VirtualConnectionPriceAccessPointType = "GW"
	CHAINGROUP_VirtualConnectionPriceAccessPointType VirtualConnectionPriceAccessPointType = "CHAINGROUP"
	NETWORK_VirtualConnectionPriceAccessPointType    VirtualConnectionPriceAccessPointType = "NETWORK"
)