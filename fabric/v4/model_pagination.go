/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.10
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Pagination response information
type Pagination struct {
	// Index of the first item returned in the response. The default is 0.
	Offset int32 `json:"offset,omitempty"`
	// Maximum number of search results returned per page. Number must be between 1 and 100, and the default is 20.
	Limit int32 `json:"limit"`
	// Total number of elements returned.
	Total int32 `json:"total"`
	// URL relative to the next item in the response.
	Next string `json:"next,omitempty"`
	// URL relative to the previous item in the response.
	Previous string `json:"previous,omitempty"`
}
