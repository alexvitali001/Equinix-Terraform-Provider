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

// NetworkChangeStatus : Current outcome of the change flow
type NetworkChangeStatus string

// List of NetworkChangeStatus
const (
	APPROVED_NetworkChangeStatus               NetworkChangeStatus = "APPROVED"
	COMPLETED_NetworkChangeStatus              NetworkChangeStatus = "COMPLETED"
	FAILED_NetworkChangeStatus                 NetworkChangeStatus = "FAILED"
	REJECTED_NetworkChangeStatus               NetworkChangeStatus = "REJECTED"
	REQUESTED_NetworkChangeStatus              NetworkChangeStatus = "REQUESTED"
	SUBMITTED_FOR_APPROVAL_NetworkChangeStatus NetworkChangeStatus = "SUBMITTED_FOR_APPROVAL"
)
