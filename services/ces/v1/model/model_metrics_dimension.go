/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 指标维度
type MetricsDimension struct {
	// 维度名
	Name *string `json:"name,omitempty"`
	// 维度值
	Value *string `json:"value,omitempty"`
}

func (o MetricsDimension) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MetricsDimension", string(data)}, " ")
}
