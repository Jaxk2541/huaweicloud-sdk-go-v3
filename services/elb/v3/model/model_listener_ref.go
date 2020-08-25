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

// 监听器
type ListenerRef struct {
	// 监听器ID。
	Id string `json:"id"`
}

func (o ListenerRef) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListenerRef", string(data)}, " ")
}
