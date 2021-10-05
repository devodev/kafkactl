package presentation

import (
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	clusterHeader        = []string{"cluster_id", "brokers", "topics_count"}
	clusterHeaderDataMap = func(data Cluster) map[string]string {
		return map[string]string{
			"cluster_id":   data.ClusterID,
			"brokers":      intSliceToString(data.Brokers.brokerIDs(), ","),
			"topics_count": strconv.Itoa(data.TopicsCount),
		}
	}
)

func MapCluster(cData *v3.ClusterData, bData BrokerList, tData TopicList) *Cluster {
	return &Cluster{
		ClusterID:   cData.ClusterID,
		Brokers:     bData,
		TopicsCount: len(tData),
	}
}

type Cluster struct {
	ClusterID   string     `json:"cluster_id"`
	Brokers     BrokerList `json:"brokers"`
	TopicsCount int        `json:"topics_count"`
}

func (t Cluster) TableHeader() []string {
	return clusterHeader
}

func (t Cluster) TableRows() []map[string]string {
	rows := []map[string]string{clusterHeaderDataMap(t)}
	return rows
}

type ClusterList []Cluster

func (t ClusterList) TableHeader() []string {
	return clusterHeader
}

func (t ClusterList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = clusterHeaderDataMap(data)
	}
	return rows
}
