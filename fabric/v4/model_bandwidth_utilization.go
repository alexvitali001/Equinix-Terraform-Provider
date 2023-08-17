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

// Bandwidth utilization statistics (octet counters-based)
type BandwidthUtilization struct {
	// Aggregated data transfer capacity. Possible values- Mbps, megabits (1 million bits) per second; Gbps, gigabits (1 billion bits) per second.
	Unit string `json:"unit,omitempty"`
	// An interval formatted value, indicating the time-interval the metric objects within the response represent
	MetricInterval string     `json:"metricInterval,omitempty"`
	Inbound        *Direction `json:"inbound,omitempty"`
	Outbound       *Direction `json:"outbound,omitempty"`
}
