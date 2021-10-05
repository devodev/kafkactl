package presentation

import (
	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	consumerHeader        = []string{"consumer_group_id", "consumer_id", "client_id", "instance_id"}
	consumerHeaderDataMap = func(data Consumer) map[string]string {
		return map[string]string{
			"consumer_group_id": data.ConsumerGroupID,
			"consumer_id":       data.ConsumerID,
			"client_id":         data.ClientID,
			"instance_id":       data.InstanceID,
		}
	}
)

func MapConsumer(data *v3.ConsumerData) *Consumer {
	return &Consumer{
		ConsumerGroupID: data.ConsumerGroupID,
		ConsumerID:      data.ConsumerID,
		ClientID:        data.ClientID,
		InstanceID:      data.InstanceID,
	}
}

type Consumer struct {
	ConsumerGroupID string `json:"consumer_group_id"`
	ConsumerID      string `json:"consumer_id"`
	ClientID        string `json:"client_id"`
	InstanceID      string `json:"instance_id"`
}

func (t Consumer) TableHeader() []string {
	return consumerHeader
}

func (t Consumer) TableRows() []map[string]string {
	rows := []map[string]string{consumerHeaderDataMap(t)}
	return rows
}

type ConsumerList []Consumer

func (t ConsumerList) TableHeader() []string {
	return consumerHeader
}

func (t ConsumerList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = consumerHeaderDataMap(data)
	}
	return rows
}
