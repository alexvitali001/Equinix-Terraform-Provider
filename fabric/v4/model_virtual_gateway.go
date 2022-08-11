/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.3
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Fabric Gateway Access point object
type VirtualGateway struct {
	Type_ string `json:"type"`
	// Customer-provided Fabric Gateway name
	Name string `json:"name"`
	Location *SimplifiedLocationWithoutIbx `json:"location"`
	Package_ *VirtualGatewayPackageType `json:"package"`
	Order *Order `json:"order,omitempty"`
	Project *Project `json:"project,omitempty"`
	Account *SimplifiedAccount `json:"account"`
	// Preferences for notifications on connection configuration or status changes
	Notifications []SimplifiedNotification `json:"notifications"`
	// Fabric Gateway URI
	Href string `json:"href"`
	// Equinix-assigned access point identifier
	Uuid string `json:"uuid"`
	State *VirtualGatewayAccessPointState `json:"state"`
	Operation *VirtualGatewayOperation `json:"operation"`
	ChangeLog *Changelog `json:"changeLog"`
	Change *GatewayChange `json:"change,omitempty"`
}
