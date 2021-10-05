package presentation

import (
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	partitionHeader        = []string{"topic_name", "partition_id", "leader_broker", "replicas", "isr"}
	partitionHeaderDataMap = func(data Partition) map[string]string {
		leader := ""
		if data.LeaderBroker != 0 {
			leader = strconv.Itoa(data.LeaderBroker)
		}
		return map[string]string{
			"topic_name":    data.TopicName,
			"partition_id":  strconv.Itoa(data.PartitionID),
			"leader_broker": leader,
			"replicas":      intSliceToString(data.ReplicaBrokers, ","),
			"isr":           intSliceToString(data.IsrBrokers, ","),
		}
	}
)

func MapPartition(data *v3.PartitionData, replicas PartitionReplicaList) *Partition {
	leaderBroker := 0
	if leader, err := replicas.leader(data.PartitionID); err == nil {
		leaderBroker = leader
	}
	replicaBrokers := replicas.replicaBrokers(data.PartitionID)
	isrBrokers := replicas.isrBrokers(data.PartitionID)
	return &Partition{
		TopicName:      data.TopicName,
		PartitionID:    data.PartitionID,
		LeaderBroker:   leaderBroker,
		ReplicaBrokers: replicaBrokers,
		IsrBrokers:     isrBrokers,
	}
}

type Partition struct {
	TopicName      string `json:"topic_name"`
	PartitionID    int    `json:"partition_id"`
	LeaderBroker   int    `json:"leader_broker"`
	ReplicaBrokers []int  `json:"replica_brokers"`
	IsrBrokers     []int  `json:"isr_brokers"`
}

func (t Partition) TableHeader() []string {
	return partitionHeader
}

func (t Partition) TableRows() []map[string]string {
	rows := []map[string]string{partitionHeaderDataMap(t)}
	return rows
}

type PartitionList []Partition

func (t PartitionList) TableHeader() []string {
	return partitionHeader
}

func (t PartitionList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = partitionHeaderDataMap(data)
	}
	return rows
}
