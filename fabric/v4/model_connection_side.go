/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Connection configuration object for each side of multi-segment connection
type ConnectionSide struct {
	ServiceToken   *ServiceToken             `json:"serviceToken,omitempty"`
	AccessPoint    *AccessPoint              `json:"accessPoint,omitempty"`
	CompanyProfile *ConnectionCompanyProfile `json:"companyProfile,omitempty"`
	Invitation     *ConnectionInvitation     `json:"invitation,omitempty"`
	// Any additional information, which is not part of connection metadata or configuration
	AdditionalInfo []ConnectionSideAdditionalInfo `json:"additionalInfo,omitempty"`
}
