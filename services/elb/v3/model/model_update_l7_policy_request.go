/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateL7PolicyRequest struct {
	L7policyId string                     `json:"l7policy_id"`
	Body       *UpdateL7PolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateL7PolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateL7PolicyRequest", string(data)}, " ")
}
