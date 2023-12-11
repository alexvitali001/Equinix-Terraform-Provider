/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// EPT service instance
type PrecisionTimeServiceCreateResponse struct {
	Type_ string `json:"type"`
	Href  string `json:"href"`
	// uuid of the ept service
	Uuid string `json:"uuid"`
	// name of the ept service
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// indicate service state
	State    string           `json:"state"`
	Package_ *PackageResponse `json:"package"`
	// fabric l2 connections used for the ept service
	Connections          []FabricConnectionUuid `json:"connections,omitempty"`
	Ipv4                 *Ipv4                  `json:"ipv4"`
	Account              *Account               `json:"account,omitempty"`
	AdvanceConfiguration *AdvanceConfiguration  `json:"advanceConfiguration,omitempty"`
	Project              *Project               `json:"project,omitempty"`
}
