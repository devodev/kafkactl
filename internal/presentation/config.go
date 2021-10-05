package presentation

import (
	"github.com/devodev/kafkactl/internal/config"
)

var (
	configHeader        = []string{"current", "context_name", "base_url", "cluster_id"}
	configHeaderDataMap = func(data Config) map[string]string {
		current := " "
		if data.IsCurrent {
			current = "*"
		}
		return map[string]string{
			"current":      current,
			"context_name": data.ContextName,
			"base_url":     data.BaseURL,
			"cluster_id":   data.ClusterID,
		}
	}
)

func MapConfig(data *config.Context, isCurrent bool) *Config {
	return &Config{
		IsCurrent:   isCurrent,
		ContextName: data.Name,
		BaseURL:     data.BaseURL,
		ClusterID:   data.ClusterID,
	}
}

type Config struct {
	IsCurrent   bool   `json:"is_current"`
	ContextName string `json:"context_name"`
	BaseURL     string `json:"base_url"`
	ClusterID   string `json:"cluster_id"`
}

func (c Config) TableHeader() []string {
	return configHeader
}

func (c Config) TableRows() []map[string]string {
	rows := []map[string]string{configHeaderDataMap(c)}
	return rows
}

type ConfigList []Config

func (c ConfigList) TableHeader() []string {
	return configHeader
}

func (c ConfigList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(c))
	for idx, data := range c {
		rows[idx] = configHeaderDataMap(data)
	}
	return rows
}

type ConfigAlphabeticSort ConfigList

func (t ConfigAlphabeticSort) Len() int {
	return len(t)
}

func (t ConfigAlphabeticSort) Less(i, j int) bool {
	iStr := t[i].ContextName
	jStr := t[j].ContextName

	return iStr < jStr
}

func (t ConfigAlphabeticSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
