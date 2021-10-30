package presentation

import (
	"fmt"
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	brokerHeader        = []string{"broker_id", "host", "partition_replicas_count"}
	brokerHeaderDataMap = func(data Broker) map[string]string {
		return map[string]string{
			"broker_id":                strconv.Itoa(data.BrokerID),
			"host":                     fmt.Sprintf("%s:%d", data.Host, data.Port),
			"partition_replicas_count": strconv.Itoa(data.PartitionReplicasCount),
		}
	}
)

func MapBroker(bData *v3.BrokerData, prData PartitionReplicaList) *Broker {
	return &Broker{
		BrokerID:               bData.BrokerID,
		Host:                   bData.Host,
		Port:                   bData.Port,
		PartitionReplicasCount: len(prData.partitionReplicas(bData.BrokerID)),
	}
}

type Broker struct {
	BrokerID               int    `json:"broker_id"`
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	PartitionReplicasCount int    `json:"partition_replicas_count"`
}

func (t Broker) TableHeader() []string {
	return brokerHeader
}

func (t Broker) TableRows() []map[string]string {
	rows := []map[string]string{brokerHeaderDataMap(t)}
	return rows
}

type BrokerList []Broker

func (b BrokerList) brokerIDs() []int {
	brokerIDs := make([]int, 0, len(b))
	for _, broker := range b {
		brokerIDs = append(brokerIDs, broker.BrokerID)
	}
	return brokerIDs
}

func (b BrokerList) BrokerIDMap() map[int]struct{} {
	m := make(map[int]struct{})

	brokerIDs := b.brokerIDs()
	for _, id := range brokerIDs {
		m[id] = struct{}{}
	}
	return m
}

func (b BrokerList) TableHeader() []string {
	return brokerHeader
}

func (b BrokerList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(b))
	for idx, data := range b {
		rows[idx] = brokerHeaderDataMap(data)
	}
	return rows
}
