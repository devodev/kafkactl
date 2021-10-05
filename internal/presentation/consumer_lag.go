package presentation

import (
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	consumerLagHeader = []string{
		"consumer_group_id", "consumer_id",
		"topic_name", "partition_id",
		"current_offset", "log_end_offset", "lag",
	}
	consumerLagHeaderDataMap = func(data ConsumerLag) map[string]string {
		return map[string]string{
			"consumer_group_id": data.ConsumerGroupID,
			"consumer_id":       data.ConsumerID,
			"topic_name":        data.TopicName,
			"partition_id":      strconv.Itoa(data.PartitionID),
			"current_offset":    strconv.Itoa(data.CurrentOffset),
			"log_end_offset":    strconv.Itoa(data.LogEndOffset),
			"lag":               strconv.Itoa(data.Lag),
		}
	}
)

func MapConsumerLag(data *v3.ConsumerLagData) *ConsumerLag {
	return &ConsumerLag{
		ConsumerGroupID: data.ConsumerGroupID,
		ConsumerID:      data.ConsumerID,
		ClientID:        data.ClientID,
		InstanceID:      data.InstanceID,
		TopicName:       data.TopicName,
		PartitionID:     data.PartitionID,
		CurrentOffset:   data.CurrentOffset,
		LogEndOffset:    data.LogEndOffset,
		Lag:             data.Lag,
	}
}

type ConsumerLag struct {
	ConsumerGroupID string `json:"consumer_group_id"`
	ConsumerID      string `json:"consumer_id"`
	InstanceID      string `json:"instance_id"`
	ClientID        string `json:"client_id"`
	TopicName       string `json:"topic_name"`
	PartitionID     int    `json:"partition_id"`
	CurrentOffset   int    `json:"current_offset"`
	LogEndOffset    int    `json:"log_end_offset"`
	Lag             int    `json:"lag"`
}

func (c ConsumerLag) TableHeader() []string {
	return consumerLagHeader
}

func (c ConsumerLag) TableRows() []map[string]string {
	rows := []map[string]string{consumerLagHeaderDataMap(c)}
	return rows
}

type ConsumerLagList []ConsumerLag

func (c ConsumerLagList) TableHeader() []string {
	return consumerLagHeader
}

func (c ConsumerLagList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(c))
	for idx, data := range c {
		rows[idx] = consumerLagHeaderDataMap(data)
	}
	return rows
}
