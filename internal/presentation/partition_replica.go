package presentation

import (
	"fmt"
	"strconv"
	"strings"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	partitionReplicaHeader        = []string{"topic_name", "partition_id", "broker_id", "is_leader", "is_in_sync"}
	partitionReplicaHeaderDataMap = func(data PartitionReplica) map[string]string {
		return map[string]string{
			"topic_name":   data.TopicName,
			"partition_id": strconv.Itoa(data.PartitionID),
			"broker_id":    strconv.Itoa(data.BrokerID),
			"is_leader":    fmt.Sprintf("%v", data.IsLeader),
			"is_in_sync":   fmt.Sprintf("%v", data.IsInSync),
		}
	}
)

func MapPartitionReplica(data *v3.ReplicaData) *PartitionReplica {
	return &PartitionReplica{
		TopicName:   data.TopicName,
		PartitionID: data.PartitionID,
		BrokerID:    data.BrokerID,
		IsLeader:    data.IsLeader,
		IsInSync:    data.IsInSync,
	}
}

type PartitionReplica struct {
	TopicName   string `json:"topic_name"`
	PartitionID int    `json:"partition_id"`
	BrokerID    int    `json:"broker_id"`
	IsLeader    bool   `json:"is_leader"`
	IsInSync    bool   `json:"is_in_sync"`
}

func (t PartitionReplica) TableHeader() []string {
	return partitionReplicaHeader
}

func (t PartitionReplica) TableRows() []map[string]string {
	rows := []map[string]string{partitionReplicaHeaderDataMap(t)}
	return rows
}

type PartitionReplicaList []PartitionReplica

func (t PartitionReplicaList) leader(PartitionID int) (int, error) {
	for _, replica := range t {
		if replica.PartitionID == PartitionID && replica.IsLeader {
			return replica.BrokerID, nil
		}
	}
	return 0, fmt.Errorf("no found")
}

func (t PartitionReplicaList) replicaBrokers(PartitionID int) []int {
	lookup := make(map[int]struct{}, len(t))
	for _, replica := range t {
		if replica.PartitionID == PartitionID {
			lookup[replica.BrokerID] = struct{}{}
		}
	}
	brokerIDs := make([]int, 0, len(lookup))
	for id := range lookup {
		brokerIDs = append(brokerIDs, id)
	}
	return brokerIDs
}

func (t PartitionReplicaList) partitionReplicas(BrokerID int) []int {
	lookup := make(map[int]struct{}, len(t))
	for _, replica := range t {
		if replica.BrokerID == BrokerID {
			lookup[replica.PartitionID] = struct{}{}
		}
	}
	partitionIDs := make([]int, 0, len(lookup))
	for id := range lookup {
		partitionIDs = append(partitionIDs, id)
	}
	return partitionIDs
}

func (t PartitionReplicaList) isrBrokers(PartitionID int) []int {
	lookup := make(map[int]struct{}, len(t))
	for _, replica := range t {
		if replica.PartitionID == PartitionID && replica.IsInSync {
			lookup[replica.BrokerID] = struct{}{}
		}
	}
	brokerIDs := make([]int, 0, len(lookup))
	for id := range lookup {
		brokerIDs = append(brokerIDs, id)
	}
	return brokerIDs
}

func (t PartitionReplicaList) TableHeader() []string {
	return partitionReplicaHeader
}

func (t PartitionReplicaList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = partitionReplicaHeaderDataMap(data)
	}
	return rows
}

type PartitionReplicaAlphabeticSort PartitionReplicaList

func (t PartitionReplicaAlphabeticSort) Len() int {
	return len(t)
}

func (t PartitionReplicaAlphabeticSort) Less(i, j int) bool {
	iStr := strings.Join([]string{t[i].TopicName, strconv.Itoa(t[i].PartitionID), strconv.Itoa(t[i].BrokerID)}, "|")
	jStr := strings.Join([]string{t[j].TopicName, strconv.Itoa(t[j].PartitionID), strconv.Itoa(t[j].BrokerID)}, "|")

	return iStr < jStr
}

func (t PartitionReplicaAlphabeticSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
