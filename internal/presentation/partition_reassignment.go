package presentation

import (
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	partitionReassignmentHeader        = []string{"topic_name", "partition_id", "adding_replicas", "removing_replicas"}
	partitionReassignmentHeaderDataMap = func(data PartitionReassignment) map[string]string {
		return map[string]string{
			"topic_name":        data.TopicName,
			"partition_id":      strconv.Itoa(data.PartitionID),
			"adding_replicas":   intSliceToString(data.AddingReplicas, ","),
			"removing_replicas": intSliceToString(data.RemovingReplicas, ","),
		}
	}
)

func MapPartitionReassignment(tcData *v3.PartitionReassignmentData) *PartitionReassignment {
	return &PartitionReassignment{
		TopicName:        tcData.TopicName,
		PartitionID:      tcData.PartitionID,
		AddingReplicas:   tcData.AddingReplicas,
		RemovingReplicas: tcData.RemovingReplicas,
	}
}

type PartitionReassignment struct {
	TopicName        string `json:"topic_name"`
	PartitionID      int    `json:"partition_id"`
	AddingReplicas   []int  `json:"adding_replicas"`
	RemovingReplicas []int  `json:"removing_replicas"`
}

func (p PartitionReassignment) TableHeader() []string {
	return partitionReassignmentHeader
}

func (p PartitionReassignment) TableRows() []map[string]string {
	rows := []map[string]string{partitionReassignmentHeaderDataMap(p)}
	return rows
}

type PartitionReassignmentList []PartitionReassignment

func (p PartitionReassignmentList) TableHeader() []string {
	return partitionReassignmentHeader
}

func (p PartitionReassignmentList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(p))
	for idx, data := range p {
		rows[idx] = partitionReassignmentHeaderDataMap(data)
	}
	return rows
}
