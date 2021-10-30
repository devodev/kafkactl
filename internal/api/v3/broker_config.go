package v3

var (
	BrokerConfigDeleteOperation string = "DELETE"
)

type BrokerConfigData struct {
	V3BaseData
	ClusterID   string `json:"cluster_id"`
	BrokerID    int    `json:"broker_id"`
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

type BrokerConfigListResponse struct {
	V3Base
	Data []BrokerConfigData `json:"data"`
}

type BrokerConfigGetResponse struct {
	BrokerConfigData
}

type BrokerConfigBatchAlterData struct {
	Name      string `json:"name"`
	Value     string `json:"value,omitempty"`
	Operation string `json:"operation,omitempty"`
}

type BrokerConfigBatchAlterRequest struct {
	Data []BrokerConfigBatchAlterData `json:"data"`
}
