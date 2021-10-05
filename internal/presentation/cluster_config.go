package presentation

import (
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	clusterConfigHeader        = []string{"name", "value", "is_default", "is_read_only"}
	clusterConfigHeaderDataMap = func(data ClusterConfig) map[string]string {
		return map[string]string{
			"name":         data.Name,
			"value":        data.Value,
			"is_default":   fmt.Sprintf("%v", data.IsDefault),
			"is_read_only": fmt.Sprintf("%v", data.IsReadOnly),
		}
	}
)

func MapClusterConfig(data *v3.ClusterConfigData) *ClusterConfig {
	return &ClusterConfig{
		Name:       data.Name,
		Value:      data.Value,
		IsDefault:  data.IsDefault,
		IsReadOnly: data.IsReadOnly,
	}
}

type ClusterConfig struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	IsDefault  bool   `json:"is_default"`
	IsReadOnly bool   `json:"is_read_only"`
}

func (t ClusterConfig) TableHeader() []string {
	return clusterConfigHeader
}

func (t ClusterConfig) TableRows() []map[string]string {
	rows := []map[string]string{clusterConfigHeaderDataMap(t)}
	return rows
}

type ClusterConfigList []ClusterConfig

func (t ClusterConfigList) TableHeader() []string {
	return clusterConfigHeader
}

func (t ClusterConfigList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = clusterConfigHeaderDataMap(data)
	}
	return rows
}
