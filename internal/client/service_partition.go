package client

import (
	"context"
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	partitionListEndpoint = "/v3/clusters/%s/topics/%s/partitions"
	partitionGetEndpoint  = "/v3/clusters/%s/topics/%s/partitions/%d"
)

type ServicePartition service

func (s *ServicePartition) List(ctx context.Context, clusterID, topicName string, resp *v3.PartitionListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(partitionListEndpoint, clusterID, topicName), resp)
}

func (s *ServicePartition) ListWide(ctx context.Context, clusterID, topicName string) (presentation.PartitionList, error) {
	var partitionResp v3.PartitionListResponse
	if err := s.List(ctx, clusterID, topicName, &partitionResp); err != nil {
		return nil, err
	}
	partitions := make(presentation.PartitionList, 0, len(partitionResp.Data))
	for _, partitionData := range partitionResp.Data {
		replicas, err := s.client.PartitionReplica.ListWide(ctx, clusterID, topicName, partitionData.PartitionID)
		if err != nil {
			return nil, err
		}
		partition := presentation.MapPartition(&partitionData, replicas)
		partitions = append(partitions, *partition)
	}
	return partitions, nil
}

func (s *ServicePartition) Get(ctx context.Context, clusterID, topicName string, partitionID int, resp *v3.PartitionGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(partitionGetEndpoint, clusterID, topicName, partitionID), resp)
}

func (s *ServicePartition) GetWide(ctx context.Context, clusterID, topicName string, partitionID int) (*presentation.Partition, error) {
	var partitionResp v3.PartitionGetResponse
	if err := s.Get(ctx, clusterID, topicName, partitionID, &partitionResp); err != nil {
		return nil, err
	}
	replicas, err := s.client.PartitionReplica.ListWide(ctx, clusterID, topicName, partitionID)
	if err != nil {
		return nil, err
	}
	partition := presentation.MapPartition(&partitionResp.PartitionData, replicas)
	return partition, nil
}
