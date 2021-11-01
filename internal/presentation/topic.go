package presentation

import (
	"fmt"
	"strconv"
	"strings"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	topicHeader  = []string{"topic_name", "replication_factor", "partitions_count", "configs_overridden", "is_internal"}
	topicDataMap = func(data Topic) map[string]string {
		return map[string]string{
			"topic_name":         data.TopicName,
			"replication_factor": strconv.Itoa(data.ReplicationFactor),
			"partitions_count":   strconv.Itoa(data.PartitionsCount),
			"configs":            formatTopicConfigs(data.ConfigsOverridden),
			"is_internal":        fmt.Sprintf("%v", data.IsInternal),
		}
	}
	TopicDescribeTemplate = `
{{- $partitionsTable := tableify .Partitions  -}}
{{- $groupsTable := tableify .Groups  -}}
{{- printf "%-25s %s" "Topic:" .TopicName }}
{{ printf "%-25s %d" "ReplicationFactor:" .ReplicationFactor }}
{{ printf "%-25s %d" "PartitionsCount:" .PartitionsCount }}
{{ printf "%-25s %v" "IsInternal:" .IsInternal }}
{{- if eq (len .ConfigsOverridden) 0 }}
{{ printf "%-25s -" "Configs:" -}}
{{- end }}
{{- range $idx, $config := .ConfigsOverridden -}}
{{- $c := printf "%s=%s" $config.Name $config.Value -}}
{{- if eq $idx 0 }}
{{ printf "%-25s %s" "Configs:" $c -}}
{{ else }}
{{ printf "%-25s %s" "" $c -}}
{{ end -}}
{{- end }}
{{- if eq (len .Groups) 0 }}
{{ printf "%-25s -" "Consumer Groups:" -}}
{{ else }}
{{ printf "%-25s" "Active Consumer Groups:" }}
{{ $groupsTable -}}
{{- end }}
{{ printf "%-25s" "Partitions:" }}
{{ $partitionsTable -}}
`
)

func MapTopic(data *v3.TopicData, pData PartitionList, cData TopicConfigList) *Topic {
	partitionsCount := len(pData)

	configs := cData.shortAndOverridden()

	return &Topic{
		TopicName:         data.TopicName,
		ReplicationFactor: data.ReplicationFactor,
		PartitionsCount:   partitionsCount,
		Partitions:        pData,
		Configs:           cData,
		ConfigsOverridden: configs,
		IsInternal:        data.IsInternal,
	}
}

type Topic struct {
	TopicName         string             `json:"topic_name"`
	ReplicationFactor int                `json:"replication_factor"`
	PartitionsCount   int                `json:"partitions_count"`
	Partitions        PartitionList      `json:"partitions"`
	Configs           TopicConfigList    `json:"configs"`
	ConfigsOverridden []TopicConfigShort `json:"configs_overridden"`
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
