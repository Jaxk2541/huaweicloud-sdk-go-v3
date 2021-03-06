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
type KeystoneListPermissionsResponse struct {
	Links *Links `json:"links,omitempty"`
	// 权限信息列表。
	Roles *[]RoleResult `json:"roles,omitempty"`
}

func (o KeystoneListPermissionsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListPermissionsResponse", string(data)}, " ")
}
