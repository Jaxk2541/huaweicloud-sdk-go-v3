/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowNetworkIpAvailabilitiesResponse struct {
	NetworkIpAvailability *NetworkIpAvailability `json:"network_ip_availability,omitempty"`
}

func (o ShowNetworkIpAvailabilitiesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowNetworkIpAvailabilitiesResponse", string(data)}, " ")
}
