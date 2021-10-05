package presentation

import (
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	topicConfigHeader        = []string{"topic_name", "name", "value", "is_default"}
	topicConfigHeaderDataMap = func(data TopicConfig) map[string]string {
		return map[string]string{
			"topic_name": data.TopicName,
			"name":       data.Name,
			"value":      data.Value,
			"is_default": fmt.Sprintf("%v", data.IsDefault),
		}
	}
)

func MapTopicConfig(data *v3.TopicConfigData) *TopicConfig {
	return &TopicConfig{
		TopicName: data.TopicName,
		Value:     data.Value,
		Name:      data.Name,
		IsDefault: data.IsDefault,
	}
}

type TopicConfigShort struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TopicConfig struct {
	TopicName string `json:"topic_name"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	IsDefault bool   `json:"is_default"`
}

func (t TopicConfig) TableHeader() []string {
	return topicConfigHeader
}

func (t TopicConfig) TableRows() []map[string]string {
	rows := []map[string]string{topicConfigHeaderDataMap(t)}
	return rows
}

type TopicConfigList []TopicConfig

func (t TopicConfigList) TableHeader() []string {
	return topicConfigHeader
}

func (t TopicConfigList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = topicConfigHeaderDataMap(data)
	}
	return rows
}
