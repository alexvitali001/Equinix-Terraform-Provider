/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b> </br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

type TimeServiceSimpleExpression struct {
	// Possible field names to use on filters:  * `/project/projectId` - project id (mandatory)  * `/name` - Precision Time Service name  * `/uuid` - Precision Time Service uuid  * `/type` - Precision Time Service protocol  * `/state` - Precision Time Service status  * `/account/accountNumber` - Precision Time Service account number  * `/package/code` - Precision Time Service package  * `/_*` - all-category search
	Property string `json:"property,omitempty"`
	// Possible operators to use on filters:  * `=` - equal  * `!=` - not equal  * `[NOT] BETWEEN` - (not) between  * `[NOT] LIKE` - (not) like  * `[NOT] IN` - (not) in  * `ILIKE` - case-insensitive like
	Operator string   `json:"operator,omitempty"`
	Values   []string `json:"values,omitempty"`
}
