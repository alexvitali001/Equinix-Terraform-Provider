/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b> </br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// TimeServiceSortBy : Possible field names to use on sorting
type TimeServiceSortBy string

// List of TimeServiceSortBy
const (
	NAME_TimeServiceSortBy                        TimeServiceSortBy = "/name"
	UUID_TimeServiceSortBy                        TimeServiceSortBy = "/uuid"
	STATE_TimeServiceSortBy                       TimeServiceSortBy = "/state"
	TYPE__TimeServiceSortBy                       TimeServiceSortBy = "/type"
	PACKAGECODE_TimeServiceSortBy                 TimeServiceSortBy = "/package/code"
	CHANGE_LOGCREATED_DATE_TIME_TimeServiceSortBy TimeServiceSortBy = "/changeLog/createdDateTime"
	CHANGE_LOGUPDATED_DATE_TIME_TimeServiceSortBy TimeServiceSortBy = "/changeLog/updatedDateTime"
)