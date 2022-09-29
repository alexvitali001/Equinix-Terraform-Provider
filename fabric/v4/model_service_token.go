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

import (
	"time"
)

// Create Service Tokens (v4) generates Equinix Fabric™ service tokens. These tokens authorize users to access protected resources and services. The tokens remove sensitive content from the environment, rather than just masking it, making the protected data impossible to unencrypt or decrypt. Resource owners can distribute the tokens to trusted partners and vendors, allowing selected third parties to work directly with Equinix network assets.
type ServiceToken struct {
	Type_ *ServiceTokenType `json:"type,omitempty"`
	// An absolute URL that is the subject of the link's context.
	Href string `json:"href,omitempty"`
	// Equinix-assigned service token identifier
	Uuid string `json:"uuid,omitempty"`
	// Service token description
	Description string `json:"description,omitempty"`
	// Expiration date and time of the service token.
	ExpirationDateTime time.Time               `json:"expirationDateTime,omitempty"`
	Connection         *ServiceTokenConnection `json:"connection,omitempty"`
	State              *ServiceTokenState      `json:"state,omitempty"`
	// Service token related notifications
	Notifications []SimplifiedNotification `json:"notifications,omitempty"`
	Account       *SimplifiedAccount       `json:"account,omitempty"`
	Changelog     *Changelog               `json:"changelog,omitempty"`
	Project       *Project                 `json:"project,omitempty"`
}