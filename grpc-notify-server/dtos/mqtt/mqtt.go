package dtos

// MQTTResultDto MQTTResultDto
type MQTTResultDto struct {
	ActionName string      `json:"actionName"`
	Message    interface{} `json:"message"`
}
