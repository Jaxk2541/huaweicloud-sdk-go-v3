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
type ShowUserLoginProtectResponse struct {
	LoginProtect *LoginProtectResult `json:"login_protect,omitempty"`
}

func (o ShowUserLoginProtectResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowUserLoginProtectResponse", string(data)}, " ")
}
