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

// Request Object
type DisableKeyRequest struct {
	VersionId string                 `json:"version_id"`
	Body      *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o DisableKeyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DisableKeyRequest", string(data)}, " ")
}
