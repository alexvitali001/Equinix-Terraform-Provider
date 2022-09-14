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

// Network specification
type SimplifiedNetwork struct {
	// Network URI
	Href string `json:"href,omitempty"`
	// Equinix-assigned network identifier
	Uuid string `json:"uuid"`
	// Customer-assigned network name
	Name      string                   `json:"name,omitempty"`
	State     *NetworkState            `json:"state,omitempty"`
	Account   *SimplifiedAccount       `json:"account,omitempty"`
	Change    *SimplifiedNetworkChange `json:"change,omitempty"`
	Operation *NetworkOperation        `json:"operation,omitempty"`
	ChangeLog *Changelog               `json:"changeLog,omitempty"`
	// Network sub-resources links
	Links    []Link              `json:"links,omitempty"`
	Type_    *NetworkType        `json:"type,omitempty"`
	Scope    *NetworkScope       `json:"scope,omitempty"`
	Location *SimplifiedLocation `json:"location,omitempty"`
}
