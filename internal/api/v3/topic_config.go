package v3

var (
	TopicConfigDeleteOperation string = "DELETE"
)

type TopicConfigData struct {
	V3BaseData
	ClusterID   string `json:"cluster_id"`
	TopicName   string `json:"topic_name"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	IsDefault   bool   `json:"is_default"`
	IsReadOnly  bool   `json:"is_read_only"`
	IsSensitive bool   `json:"is_sensitive"`
	Source      string `json:"source"`
	Synonyms    []struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Source string `json:"source"`
	} `json:"synonyms"`
}

type TopicConfigListResponse struct {
	V3Base
	Data []TopicConfigData `json:"data"`
}

type TopicConfigGetResponse struct {
	TopicConfigData
}

type TopicConfigBatchAlterData struct {
	Name      string `json:"name"`
	Value     string `json:"value,omitempty"`
	Operation string `json:"operation,omitempty"`
}

type TopicConfigBatchAlterRequest struct {
	Data []TopicConfigBatchAlterData `json:"data"`
}
