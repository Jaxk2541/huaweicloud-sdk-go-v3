/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type KeyDescriptionInfo struct {
	// 密钥ID。
	KeyId *string `json:"key_id,omitempty"`
	// 密钥描述。
	KeyDescription *string `json:"key_description,omitempty"`
}

func (o KeyDescriptionInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeyDescriptionInfo", string(data)}, " ")
}
