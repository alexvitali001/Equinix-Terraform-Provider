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

type VirtualGatewaySimpleExpression struct {
	// Possible field names to use on filters:  * `/project/projectId` - project id (mandatory)  * `/name` - Fabric Gateway name  * `/uuid` - Fabric Gateway uuid  * `/state` - Fabric Gateway status  * `/location/metroCode` - Fabric Gateway metro code  * `/location/metroName` - Fabric Gateway metro name  * `/package/code` - Fabric Gateway package  * `/_*` - all-category search
	Property string `json:"property,omitempty"`
	// Possible operators to use on filters:  * `=` - equal  * `!=` - not equal  * `>` - greater than  * `>=` - greater than or equal to  * `<` - less than  * `<=` - less than or equal to  * `[NOT] BETWEEN` - (not) between  * `[NOT] LIKE` - (not) like  * `[NOT] IN` - (not) in  * `ILIKE` - case-insensitive like
	Operator string   `json:"operator,omitempty"`
	Values   []string `json:"values,omitempty"`
}
