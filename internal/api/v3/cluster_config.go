package v3

type ClusterConfigData struct {
	V3BaseData
	ClusterID   string `json:"cluster_id"`
	ConfigType  string `json:"config_type"`
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

type ClusterConfigListResponse struct {
	V3Base
	Data []ClusterConfigData `json:"data"`
}

type ClusterConfigGetResponse struct {
	ClusterConfigData
}
