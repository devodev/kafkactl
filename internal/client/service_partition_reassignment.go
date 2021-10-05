package client

import (
	"context"
	"fmt"
	"strconv"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	partitionReassignmentEndpoint = "/v3/clusters/%s/topics/%s/partitions/%s/reassignment"
)

type ServicePartitionReassignment service

func (s *ServicePartitionReassignment) ReassignmentGet(ctx context.Context, clusterID, topicName string, partitionID int, resp *v3.PartitionReassignmentGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(partitionReassignmentEndpoint, clusterID, topicName, strconv.Itoa(partitionID)), resp)
}

func (s *ServicePartitionReassignment) GetWide(ctx context.Context, clusterID, topicName string, partitionID int) (*presentation.PartitionReassignment, error) {
	var partitionReasResp v3.PartitionReassignmentGetResponse
	if err := s.ReassignmentGet(ctx, clusterID, topicName, partitionID, &partitionReasResp); err != nil {
		return nil, err
	}
	partition := presentation.MapPartitionReassignment(&partitionReasResp.PartitionReassignmentData)
	return partition, nil
}

func (s *ServicePartitionReassignment) ReassignmentList(ctx context.Context, clusterID string, topicName string, resp *v3.PartitionReassignmentListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(partitionReassignmentEndpoint, clusterID, topicName, "-"), resp)
}

func (s *ServicePartitionReassignment) ListWide(ctx context.Context, clusterID string, topicName string) (presentation.PartitionReassignmentList, error) {
	var partitionReasResp v3.PartitionReassignmentListResponse
	if err := s.ReassignmentList(ctx, clusterID, topicName, &partitionReasResp); err != nil {
		return nil, err
	}
	partitionReas := make(presentation.PartitionReassignmentList, 0, len(partitionReasResp.Data))
	for _, reasData := range partitionReasResp.Data {
		reassignment := presentation.MapPartitionReassignment(&reasData)
		partitionReas = append(partitionReas, *reassignment)
	}
	return partitionReas, nil
}

func (s *ServicePartitionReassignment) ListAllWide(ctx context.Context, clusterID string) (presentation.PartitionReassignmentList, error) {
	return s.ListWide(ctx, clusterID, "-")
}
