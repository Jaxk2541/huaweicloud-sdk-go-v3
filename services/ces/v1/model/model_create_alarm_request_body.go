/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

//
type CreateAlarmRequestBody struct {
	// 告警名称，只能包含0-9/a-z/A-Z/_/-或汉字。
	AlarmName string `json:"alarm_name"`
	//
	AlarmDescription *string        `json:"alarm_description,omitempty"`
	Metric           *MetricInfoExt `json:"metric"`
	Condition        *Condition     `json:"condition"`
	// 是否启用该条告警，默认为true。
	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`
	// 是否启用该条告警触发的动作，默认为true。注：若alarm_action_enabled为true，对应的alarm_actions、ok_actions至少有一个不能为空。若alarm_actions、ok_actions同时存在时，notificationList值保持一致。
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`
	// 告警级别，默认为2，级别为1、2、3、4。分别对应紧急、重要、次要、提示。
	AlarmLevel *int32 `json:"alarm_level,omitempty"`
	// 告警类型。 仅针对事件告警的参数，枚举类型：EVENT.SYS或者EVENT.CUSTOM
	AlarmType *CreateAlarmRequestBodyAlarmType `json:"alarm_type,omitempty"`
	// 告警触发的动作。 结构样例如下： { \"type\": \"notification\",\"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"] } type取值： notification：通知。 autoscaling：弹性伸缩。
	AlarmActions *[]AlarmActions `json:"alarm_actions,omitempty"`
	// 数据不足触发的动作（该参数已废弃，建议无需配置）。
	InsufficientdataActions *[]AlarmActions `json:"insufficientdata_actions,omitempty"`
	// 告警恢复触发的动作
	OkActions *[]AlarmActions `json:"ok_actions,omitempty"`
}

func (o CreateAlarmRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateAlarmRequestBody", string(data)}, " ")
}

type CreateAlarmRequestBodyAlarmType struct {
	value string
}

type CreateAlarmRequestBodyAlarmTypeEnum struct {
	EVENT_SYS    CreateAlarmRequestBodyAlarmType
	EVENT_CUSTOM CreateAlarmRequestBodyAlarmType
}

func GetCreateAlarmRequestBodyAlarmTypeEnum() CreateAlarmRequestBodyAlarmTypeEnum {
	return CreateAlarmRequestBodyAlarmTypeEnum{
		EVENT_SYS: CreateAlarmRequestBodyAlarmType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: CreateAlarmRequestBodyAlarmType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c CreateAlarmRequestBodyAlarmType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateAlarmRequestBodyAlarmType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
