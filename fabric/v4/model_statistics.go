/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.9
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

import (
	"time"
)

// This API provides service-level traffic metrics so that you can view access and gather key information required to manage service subscription sizing and capacity.
type Statistics struct {
	// Start and duration of the statistical analysis interval.
	StartDateTime time.Time `json:"startDateTime,omitempty"`
	// End and duration of the statistical analysis interval.
	EndDateTime time.Time `json:"endDateTime,omitempty"`
	// Point of view for connection metrics - aSide or zSide
	ViewPoint            string                `json:"viewPoint,omitempty"`
	BandwidthUtilization *BandwidthUtilization `json:"bandwidthUtilization,omitempty"`
}
