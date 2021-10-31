package presentation

import (
	"fmt"
	"regexp"
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	coordinatorRe = regexp.MustCompile(`\d+$`)
)

var (
	consumerGroupHeader        = []string{"consumer_group_id", "state", "partition_assignor", "coordinator_broker", "consumers_count"}
	consumerGroupHeaderDataMap = func(data ConsumerGroup) map[string]string {
		coordinatorBroker := ""
		if data.CoordinatorBroker != 0 {
			coordinatorBroker = strconv.Itoa(data.CoordinatorBroker)
		}
		return map[string]string{
			"consumer_group_id":  data.ConsumerGroupID,
			"state":              data.State,
			"partition_assignor": data.PartitionAssignor,
			"coordinator_broker": coordinatorBroker,
			"consumers_count":    strconv.Itoa(data.ConsumersCount),
		}
	}
	consumerGroupLagSummaryHeader = []string{
		"consumer_group_id",
		"max_consumer_id", "max_instance_id", "max_client_id", "max_topic_name", "max_partition_id",
		"max", "total",
	}
	consumerGroupLagSummaryHeaderDataMap = func(data ConsumerGroupLagSummary) map[string]string {
		return map[string]string{
			"consumer_group_id": data.ConsumerGroupID,
			"max_consumer_id":   data.MaxLagConsumerID,
			"max_instance_id":   data.MaxLagInstanceID,
			"max_client_id":     data.MaxLagClientID,
			"max_topic_name":    data.MaxLagTopicName,
			"max_partition_id":  strconv.Itoa(data.MaxLagPartitionID),
			"max":               strconv.Itoa(data.MaxLag),
			"total":             strconv.Itoa(data.TotalLag),
		}
	}
)

func MapConsumerGroup(data *v3.ConsumerGroupData, consumers ConsumerList) *ConsumerGroup {
	coordinatorBrokerID := 0
	if id, err := extractCoordinator(data.Coordinator); err == nil {
		coordinatorBrokerID = id
	}
	return &ConsumerGroup{
		ConsumerGroupID:   data.ConsumerGroupID,
		State:             data.State,
		PartitionAssignor: data.PartitionAssignor,
		CoordinatorBroker: coordinatorBrokerID,
		ConsumersCount:    len(consumers),
	}
}

func extractCoordinator(related v3.V3BaseDataRelated) (int, error) {
	brokerIDStr := coordinatorRe.FindString(related.Related)
	if brokerIDStr == "" {
		return 0, fmt.Errorf("coordinator broker id not found")
	}
	brokerID, err := strconv.Atoi(brokerIDStr)
	if err != nil {
		return 0, fmt.Errorf("invalid broker id")
	}
	return brokerID, nil
}

type ConsumerGroup struct {
	ConsumerGroupID   string `json:"consumer_group_id"`
	State             string `json:"state"`
	PartitionAssignor string `json:"partition_assignor"`
	CoordinatorBroker int    `json:"coordinator_broker"`
	ConsumersCount    int    `json:"consumers_count"`
}

func (t ConsumerGroup) TableHeader() []string {
	return consumerGroupHeader
}

func (t ConsumerGroup) TableRows() []map[string]string {
	rows := []map[string]string{consumerGroupHeaderDataMap(t)}
	return rows
}

type ConsumerGroupList []ConsumerGroup

func (t ConsumerGroupList) TableHeader() []string {
	return consumerGroupHeader
}

func (t ConsumerGroupList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = consumerGroupHeaderDataMap(data)
	}
	return rows
}

func MapConsumerGroupLagSummary(data *v3.ConsumerGroupLagSummaryData) *ConsumerGroupLagSummary {
	return &ConsumerGroupLagSummary{
		ConsumerGroupID:   data.ConsumerGroupID,
		MaxLagTopicName:   data.MaxLagTopicName,
		MaxLagPartitionID: data.MaxLagPartitionID,
		MaxLagConsumerID:  data.MaxLagConsumerID,
		MaxLagClientID:    data.MaxLagClientID,
		MaxLagInstanceID:  data.MaxLagInstanceID,
		MaxLag:            data.MaxLag,
		TotalLag:          data.TotalLag,
	}
}

type ConsumerGroupLagSummary struct {
	ConsumerGroupID   string `json:"consumer_group_id"`
	MaxLagTopicName   string `json:"max_lag_topic_name"`
	MaxLagPartitionID int    `json:"max_lag_partition_id"`
	MaxLagConsumerID  string `json:"max_lag_consumer_id"`
	MaxLagClientID    string `json:"max_lag_client_id"`
	MaxLagInstanceID  string `json:"max_lag_instance_id"`
	MaxLag            int    `json:"max_lag"`
	TotalLag          int    `json:"total_lag"`
}

func (t ConsumerGroupLagSummary) TableHeader() []string {
	return consumerGroupLagSummaryHeader
}

func (t ConsumerGroupLagSummary) TableRows() []map[string]string {
	rows := []map[string]string{consumerGroupLagSummaryHeaderDataMap(t)}
	return rows
}

type ConsumerGroupLagSummaryList []ConsumerGroupLagSummary

func (t ConsumerGroupLagSummaryList) TableHeader() []string {
	return consumerGroupLagSummaryHeader
}

func (t ConsumerGroupLagSummaryList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = consumerGroupLagSummaryHeaderDataMap(data)
	}
	return rows
}
