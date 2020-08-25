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

// Response Object
type ListL7PoliciesResponse struct {
	// 请求ID。 注：自动生成 。
	RequestId *string   `json:"request_id,omitempty"`
	PageInfo  *PageInfo `json:"page_info,omitempty"`
	// 转发策略对象列表。
	L7policies []L7Policy `json:"l7policies,omitempty"`
}

func (o ListL7PoliciesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListL7PoliciesResponse", string(data)}, " ")
}
