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

// Port specification
type Port struct {
	Type_ *PortType `json:"type"`
	// Equinix assigned response attribute for Port Id
	Id int32 `json:"id,omitempty"`
	// Equinix assigned response attribute for an absolute URL that is the subject of the link's context.
	Href string `json:"href,omitempty"`
	// Equinix assigned response attribute for  port identifier
	Uuid string `json:"uuid,omitempty"`
	// Equinix assigned response attribute for Port name
	Name string `json:"name,omitempty"`
	// Equinix assigned response attribute for Port description
	Description string `json:"description,omitempty"`
	// Physical Ports Speed in Mbps
	PhysicalPortsSpeed int32 `json:"physicalPortsSpeed"`
	// Equinix assigned response attribute for Connection count
	ConnectionsCount int32 `json:"connectionsCount,omitempty"`
	// Physical Ports Type
	PhysicalPortsType string `json:"physicalPortsType"`
	// Port connectivity type
	ConnectivitySourceType string     `json:"connectivitySourceType"`
	Project                *Project   `json:"project,omitempty"`
	State                  *PortState `json:"state,omitempty"`
	Order                  *PortOrder `json:"order,omitempty"`
	// Equinix assigned response attribute for Unique ID for a virtual port.
	CvpId     string             `json:"cvpId,omitempty"`
	Operation *PortOperation     `json:"operation,omitempty"`
	Account   *SimplifiedAccount `json:"account"`
	Changelog *Changelog         `json:"changelog,omitempty"`
	// Port service Type
	ServiceType string `json:"serviceType,omitempty"`
	// Equinix assigned response attribute for Port bandwidth in Mbps
	Bandwidth int32 `json:"bandwidth,omitempty"`
	// Equinix assigned response attribute for Port available bandwidth in Mbps
	AvailableBandwidth int32 `json:"availableBandwidth,omitempty"`
	// Equinix assigned response attribute for Port used bandwidth in Mbps
	UsedBandwidth    int32                 `json:"usedBandwidth,omitempty"`
	Location         *SimplifiedLocation   `json:"location"`
	Device           *PortDevice           `json:"device,omitempty"`
	Interface_       *PortInterface        `json:"interface,omitempty"`
	Tether           *PortTether           `json:"tether,omitempty"`
	DemarcationPoint *PortDemarcationPoint `json:"demarcationPoint,omitempty"`
	Redundancy       *PortRedundancy       `json:"redundancy,omitempty"`
	Encapsulation    *PortEncapsulation    `json:"encapsulation"`
	Lag              *PortLag              `json:"lag"`
	// Port ASN
	Asn      int32         `json:"asn,omitempty"`
	Settings *PortSettings `json:"settings"`
	// Number of physical ports
	PhysicalPortQuantity int32 `json:"physicalPortQuantity,omitempty"`
	// Notification preferences
	Notifications []PortNotification `json:"notifications,omitempty"`
	// Port additional information
	AdditionalInfo []PortAdditionalInfo `json:"additionalInfo,omitempty"`
	// Physical ports that implement this port
	PhysicalPorts []PhysicalPort `json:"physicalPorts,omitempty"`
	// Port Loas
	Loas []PortLoa `json:"loas,omitempty"`
}
