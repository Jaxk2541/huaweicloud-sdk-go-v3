/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartNewPipelineResponse struct {
	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`
	// 流水线构建ID
	BuildId *string `json:"build_id,omitempty"`
}

func (o StartNewPipelineResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"StartNewPipelineResponse", string(data)}, " ")
}
