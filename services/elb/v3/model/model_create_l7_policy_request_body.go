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

// This is a auto create Body Object
type CreateL7PolicyRequestBody struct {
	L7policy *CreateL7PolicyOption `json:"l7policy"`
}

func (o CreateL7PolicyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateL7PolicyRequestBody", string(data)}, " ")
}