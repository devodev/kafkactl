package client

import (
	"context"
	"fmt"
	"sort"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	partitionReplicasListEndpoint       = "/v3/clusters/%s/topics/%s/partitions/%d/replicas"
	partitionReplicasGetEndpoint        = "/v3/clusters/%s/topics/%s/partitions/%d/replicas/%d"
	partitionReplicasBrokerListEndpoint = "/v3/clusters/%s/brokers/%d/partition-replicas"
)

type ServicePartitionReplica service

func (s *ServicePartitionReplica) List(ctx context.Context, clusterID, topicName string, partitionID int, resp *v3.ReplicaListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(partitionReplicasListEndpoint, clusterID, topicName, partitionID), resp)
}

func (s *ServicePartitionReplica) ListWide(ctx context.Context, clusterID, topicName string, partitionID int) (presentation.PartitionReplicaList, error) {
	var partitionRepResp v3.ReplicaListResponse
	if err := s.List(ctx, clusterID, topicName, partitionID, &partitionRepResp); err != nil {
		return nil, err
	}
	partitionReps := make(presentation.PartitionReplicaList, 0, len(partitionRepResp.Data))
	for _, partitionData := range partitionRepResp.Data {
		rep := presentation.MapPartitionReplica(&partitionData)
		partitionReps = append(partitionReps, *rep)
	}
	sort.Sort(presentation.PartitionReplicaAlphabeticSort(partitionReps))
	return partitionReps, nil
}

func (s *ServicePartitionReplica) ListAllWide(ctx context.Context, clusterID, topicName string) (presentation.PartitionReplicaList, error) {
	resp, err := s.client.Partition.ListWide(ctx, clusterID, topicName)
	if err != nil {
		return nil, err
	}
	partitionReps := make(presentation.PartitionReplicaList, 0, len(resp))
	for _, partitionData := range resp {
		replicas, err := s.ListWide(ctx, clusterID, topicName, partitionData.PartitionID)
		if err != nil {
			return nil, err
		}
		partitionReps = append(partitionReps, replicas...)
	}
	sort.Sort(presentation.PartitionReplicaAlphabeticSort(partitionReps))
	return partitionReps, nil
}

func (s *ServicePartitionReplica) Get(ctx context.Context, clusterID, topicName string, partitionID, brokerID int, resp *v3.ReplicaGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(partitionReplicasGetEndpoint, clusterID, topicName, partitionID, brokerID), resp)
}

func (s *ServicePartitionReplica) GetWide(ctx context.Context, clusterID, topicName string, partitionID, brokerID int) (*presentation.PartitionReplica, error) {
	var partitionRepResp v3.ReplicaGetResponse
	if err := s.Get(ctx, clusterID, topicName, partitionID, brokerID, &partitionRepResp); err != nil {
		return nil, err
	}
	partition := presentation.MapPartitionReplica(&partitionRepResp.ReplicaData)
	return partition, nil
}

func (s *ServicePartitionReplica) ListBroker(ctx context.Context, clusterID string, brokerID int, resp *v3.ReplicaListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(partitionReplicasBrokerListEndpoint, clusterID, brokerID), resp)
}

func (s *ServicePartitionReplica) ListBrokerWide(ctx context.Context, clusterID string, brokerID int) (presentation.PartitionReplicaList, error) {
	var partitionReplicasResp v3.ReplicaListResponse
	if err := s.ListBroker(ctx, clusterID, brokerID, &partitionReplicasResp); err != nil {
		return nil, err
	}
	partitionReplicas := make(presentation.PartitionReplicaList, 0, len(partitionReplicasResp.Data))
	for _, cgData := range partitionReplicasResp.Data {
		pr := presentation.MapPartitionReplica(&cgData)
		partitionReplicas = append(partitionReplicas, *pr)
	}
	return partitionReplicas, nil
}

func (s *ServicePartitionReplica) ListBrokerAllWide(ctx context.Context, clusterID string) (presentation.PartitionReplicaList, error) {
	brokerResp, err := s.client.Broker.ListWide(ctx, clusterID)
	if err != nil {
		return nil, err
	}
	partitionReplicas := make(presentation.PartitionReplicaList, 0, len(brokerResp))
	for _, bData := range brokerResp {
		partitionReplicasResp, err := s.ListBrokerWide(ctx, clusterID, bData.BrokerID)
		if err != nil {
			return nil, err
		}
		partitionReplicas = append(partitionReplicas, partitionReplicasResp...)
	}
	return partitionReplicas, nil
}
