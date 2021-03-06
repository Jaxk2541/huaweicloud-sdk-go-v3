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

//
type NeutronPageLink struct {
	// API链接
	Href string `json:"href"`
	// API链接与该API版本的关系
	Rel string `json:"rel"`
}

func (o NeutronPageLink) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronPageLink", string(data)}, " ")
}
