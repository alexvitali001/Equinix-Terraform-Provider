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

// Configuration for dot1q to qinq translation support
type ServiceProfileLinkProtocolConfig struct {
	// was tagType - missing on wiki
	EncapsulationStrategy string   `json:"encapsulationStrategy,omitempty"`
	NamedTags             []string `json:"namedTags,omitempty"`
	// was ctagLabel
	VlanCTagLabel string `json:"vlanCTagLabel,omitempty"`
	ReuseVlanSTag bool   `json:"reuseVlanSTag,omitempty"`
	// Port encapsulation - Derived response attribute. Ignored on request payloads.
	Encapsulation string `json:"encapsulation,omitempty"`
}
