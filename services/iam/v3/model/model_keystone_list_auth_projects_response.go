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
type KeystoneListAuthProjectsResponse struct {
	Links *LinksSelf `json:"links,omitempty"`
	// 项目信息列表。
	Projects *[]AuthProjectResult `json:"projects,omitempty"`
}

func (o KeystoneListAuthProjectsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListAuthProjectsResponse", string(data)}, " ")
}
