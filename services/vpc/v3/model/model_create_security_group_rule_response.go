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
type CreateSecurityGroupRuleResponse struct {
	// 请求ID
	RequestId         *string            `json:"request_id,omitempty"`
	SecurityGroupRule *SecurityGroupRule `json:"security_group_rule,omitempty"`
}

func (o CreateSecurityGroupRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSecurityGroupRuleResponse", string(data)}, " ")
}
