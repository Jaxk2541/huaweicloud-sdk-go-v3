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

// Request Object
type BatchListMetricDataRequest struct {
	ContentType string                          `json:"Content-Type"`
	Body        *BatchListMetricDataRequestBody `json:"body,omitempty"`
}

func (o BatchListMetricDataRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchListMetricDataRequest", string(data)}, " ")
}
