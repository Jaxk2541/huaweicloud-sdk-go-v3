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

// Request Object
type ShowSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o ShowSecurityGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowSecurityGroupRequest", string(data)}, " ")
}
