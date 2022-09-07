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

// AccessPointType : Access point type
type AccessPointType string

// List of AccessPointType
const (
	VD_AccessPointType      AccessPointType = "VD"
	VG_AccessPointType      AccessPointType = "VG"
	SP_AccessPointType      AccessPointType = "SP"
	IGW_AccessPointType     AccessPointType = "IGW"
	COLO_AccessPointType    AccessPointType = "COLO"
	SUBNET_AccessPointType  AccessPointType = "SUBNET"
	GW_AccessPointType      AccessPointType = "GW"
	NETWORK_AccessPointType AccessPointType = "NETWORK"
)
