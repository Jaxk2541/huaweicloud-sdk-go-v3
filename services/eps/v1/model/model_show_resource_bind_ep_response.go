/*
 * EPS
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
type ShowResourceBindEpResponse struct {
	// 资源列表
	Resources *[]Resources `json:"resources,omitempty"`
	// 查询失败的企业项目下的资源
	Errors *[]Errors `json:"errors,omitempty"`
	// 企业项目下的资源总数
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o ShowResourceBindEpResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowResourceBindEpResponse", string(data)}, " ")
}
