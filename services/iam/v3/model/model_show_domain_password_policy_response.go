/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainPasswordPolicyResponse struct {
	PasswordPolicy *PasswordPolicyResult `json:"password_policy,omitempty"`
}

func (o ShowDomainPasswordPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowDomainPasswordPolicyResponse", string(data)}, " ")
}