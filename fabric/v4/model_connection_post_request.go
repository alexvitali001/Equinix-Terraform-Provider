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

// Create connection post request
type ConnectionPostRequest struct {
	Type_ *ConnectionType `json:"type,omitempty"`
	// Customer-provided connection name
	Name  string `json:"name,omitempty"`
	Order *Order `json:"order,omitempty"`
	// Preferences for notifications on connection configuration or status changes
	Notifications []SimplifiedNotification `json:"notifications,omitempty"`
	// Connection bandwidth in Mbps
	Bandwidth  int32                 `json:"bandwidth,omitempty"`
	Redundancy *ConnectionRedundancy `json:"redundancy,omitempty"`
	ASide      *ConnectionSide       `json:"aSide,omitempty"`
	ZSide      *ConnectionSide       `json:"zSide,omitempty"`
	Project    *Project              `json:"project,omitempty"`
	// Connection additional information
	AdditionalInfo []ConnectionSideAdditionalInfo `json:"additionalInfo,omitempty"`
}
