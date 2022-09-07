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

// FabricGatewayCode : Virtual port gateway code
type FabricGatewayCode string

// List of FabricGatewayCode
const (
	LIMITED_FabricGatewayCode  FabricGatewayCode = "LIMITED"
	BASIC_FabricGatewayCode    FabricGatewayCode = "BASIC"
	ADVANCED_FabricGatewayCode FabricGatewayCode = "ADVANCED"
	PREMIUM_FabricGatewayCode  FabricGatewayCode = "PREMIUM"
)
