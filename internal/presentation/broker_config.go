package presentation

import (
	"fmt"
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	brokerConfigHeader        = []string{"broker_id", "name", "value", "is_default", "is_read_only"}
	brokerConfigHeaderDataMap = func(data BrokerConfig) map[string]string {
		return map[string]string{
			"broker_id":    strconv.Itoa(data.BrokerID),
			"name":         data.Name,
			"value":        data.Value,
			"is_default":   fmt.Sprintf("%v", data.IsDefault),
			"is_read_only": fmt.Sprintf("%v", data.IsReadOnly),
		}
	}
)

func MapBrokerConfig(data *v3.BrokerConfigData) *BrokerConfig {
	return &BrokerConfig{
		Name:       data.Name,
		BrokerID:   data.BrokerID,
		Value:      data.Value,
		IsDefault:  data.IsDefault,
		IsReadOnly: data.IsReadOnly,
	}
}

type BrokerConfig struct {
	BrokerID   int    `json:"broker_id"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	IsDefault  bool   `json:"is_default"`
	IsReadOnly bool   `json:"is_read_only"`
}

func (t BrokerConfig) TableHeader() []string {
	return brokerConfigHeader
}

func (t BrokerConfig) TableRows() []map[string]string {
	rows := []map[string]string{brokerConfigHeaderDataMap(t)}
	return rows
}

type BrokerConfigList []BrokerConfig

func (t BrokerConfigList) TableHeader() []string {
	return brokerConfigHeader
}

func (t BrokerConfigList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = brokerConfigHeaderDataMap(data)
	}
	return rows
}
