/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br> </br> <b>Integrations (SDKs, Tools) links:</b> </br> <a href=\"https://deploy.equinix.com/labs/fabric-java\\\">Fabric Java SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/equinix-sdk-go\\\">Fabric Go SDK</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-provider-equinix\\\">Equinix Terraform Provider</a> </br> <a href=\"https://deploy.equinix.com/labs/terraform-equinix-fabric\\\">Fabric Terraform Modules</a> </br> <a href=\"https://deploy.equinix.com/labs/pulumi-provider-equinix/\">Equinix Pulumi Provider</a> </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Actions : Connection action type
type Actions string

// List of Actions
const (
	CONNECTION_CREATION_ACCEPTANCE_Actions        Actions = "CONNECTION_CREATION_ACCEPTANCE"
	CONNECTION_CREATION_REJECTION_Actions         Actions = "CONNECTION_CREATION_REJECTION"
	CONNECTION_UPDATE_ACCEPTANCE_Actions          Actions = "CONNECTION_UPDATE_ACCEPTANCE"
	CONNECTION_UPDATE_REJECTION_Actions           Actions = "CONNECTION_UPDATE_REJECTION"
	CONNECTION_DELETION_ACCEPTANCE_Actions        Actions = "CONNECTION_DELETION_ACCEPTANCE"
	CONNECTION_REJECTION_ACCEPTANCE_Actions       Actions = "CONNECTION_REJECTION_ACCEPTANCE"
	CONNECTION_UPDATE_REQUEST_Actions             Actions = "CONNECTION_UPDATE_REQUEST"
	MIGRATION_EVPL_VC_Actions                     Actions = "MIGRATION_EVPL_VC"
	CONNECTION_PROVIDER_STATUS_REQUEST_Actions    Actions = "CONNECTION_PROVIDER_STATUS_REQUEST"
	CONNECTION_PROVIDER_BANDWIDTH_REQUEST_Actions Actions = "CONNECTION_PROVIDER_BANDWIDTH_REQUEST"
	ACCEPT_HOSTED_CONNECTION_Actions              Actions = "ACCEPT_HOSTED_CONNECTION"
	CANCEL_EVPL_VC_DRAFT_ORDERS_Actions           Actions = "CANCEL_EVPL_VC_DRAFT_ORDERS"
)
