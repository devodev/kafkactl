package presentation

import (
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	consumerAssignmentHeader        = []string{"consumer_group_id", "consumer_id", "topic_name", "partition_id"}
	consumerAssignmentHeaderDataMap = func(data ConsumerAssignment) map[string]string {
		return map[string]string{
			"consumer_group_id": data.ConsumerGroupID,
			"consumer_id":       data.ConsumerID,
			"topic_name":        data.TopicName,
			"partition_id":      strconv.Itoa(data.PartitionID),
		}
	}
)

func MapConsumerAssignment(data *v3.ConsumerAssignmentData) *ConsumerAssignment {
	return &ConsumerAssignment{
		ConsumerGroupID: data.ConsumerGroupID,
		ConsumerID:      data.ConsumerID,
		TopicName:       data.TopicName,
		PartitionID:     data.PartitionID,
	}
}

type ConsumerAssignment struct {
	ConsumerGroupID string `json:"consumer_group_id"`
	ConsumerID      string `json:"consumer_id"`
	TopicName       string `json:"topic_name"`
	PartitionID     int    `json:"partition_id"`
}

func (c ConsumerAssignment) TableHeader() []string {
	return consumerAssignmentHeader
}

func (c ConsumerAssignment) TableRows() []map[string]string {
	rows := []map[string]string{consumerAssignmentHeaderDataMap(c)}
	return rows
}

type ConsumerAssignmentList []ConsumerAssignment

func (c ConsumerAssignmentList) TableHeader() []string {
	return consumerAssignmentHeader
}

func (c ConsumerAssignmentList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(c))
	for idx, data := range c {
		rows[idx] = consumerAssignmentHeaderDataMap(data)
	}
	return rows
}
