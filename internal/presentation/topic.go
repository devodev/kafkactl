package presentation

import (
	"fmt"
	"strconv"
	"strings"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	topicHeader  = []string{"topic_name", "replication_factor", "partitions_count", "configs", "is_internal"}
	topicDataMap = func(data Topic) map[string]string {
		return map[string]string{
			"topic_name":         data.TopicName,
			"replication_factor": strconv.Itoa(data.ReplicationFactor),
			"partitions_count":   strconv.Itoa(data.PartitionsCount),
			"configs":            formatTopicConfigs(data.Configs),
			"is_internal":        fmt.Sprintf("%v", data.IsInternal),
		}
	}
	TopicDescribeTemplate = `{{ printf "%-20s %s" "Topic:" .TopicName }}
{{ printf "%-20s %d" "ReplicationFactor:" .ReplicationFactor }}
{{ printf "%-20s %d" "PartitionsCount:" .PartitionsCount }}
{{ printf "%-20s %v" "IsInternal:" .IsInternal }}
{{- range $idx, $config := .Configs -}}
{{- $c := printf "%s=%s" $config.Name $config.Value -}}
{{- if eq $idx 0 }}
{{ printf "%-20s %s" "Configs:" $c -}}
{{ else }}
{{ printf "%-20s %s" "" $c -}}
{{ end -}}
{{- end }}
`
)

func MapTopic(data *v3.TopicData, pData PartitionList, cData TopicConfigList) *Topic {
	partitionsCount := len(pData)

	configs := make([]TopicConfigShort, 0, len(cData))
	for _, topicConfigData := range cData {
		// filter out default config
		if topicConfigData.IsDefault {
			continue
		}
		configs = append(configs, TopicConfigShort{
			Name:  topicConfigData.Name,
			Value: topicConfigData.Value,
		})
	}

	return &Topic{
		TopicName:         data.TopicName,
		ReplicationFactor: data.ReplicationFactor,
		PartitionsCount:   partitionsCount,
		Configs:           configs,
		IsInternal:        data.IsInternal,
	}
}

type Topic struct {
	TopicName         string             `json:"topic_name"`
	ReplicationFactor int                `json:"replication_factor"`
	PartitionsCount   int                `json:"partitions_count"`
	Configs           []TopicConfigShort `json:"configs"`
	IsInternal        bool               `json:"is_internal"`
}

func (t Topic) TableHeader() []string {
	return topicHeader
}

func (t Topic) TableRows() []map[string]string {
	rows := []map[string]string{topicDataMap(t)}
	return rows
}

type TopicList []Topic

func (t TopicList) TableHeader() []string {
	return topicHeader
}

func (t TopicList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = topicDataMap(data)
	}
	return rows
}

type TopicAlphabeticSort TopicList

func (t TopicAlphabeticSort) Len() int {
	return len(t)
}

func (t TopicAlphabeticSort) Less(i, j int) bool {
	iStr := t[i].TopicName
	jStr := t[j].TopicName

	return iStr < jStr
}

func (t TopicAlphabeticSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func formatTopicConfigs(data []TopicConfigShort) string {
	var configValues []string
	for _, configData := range data {
		configValues = append(configValues, fmt.Sprintf("%s=%s", configData.Name, configData.Value))
	}
	return strings.Join(configValues, ",")
}
