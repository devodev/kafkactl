package client

import (
	"context"
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	consumerAssignmentListEndpoint = "/v3/clusters/%s/consumer-groups/%s/consumers/%s/assignments"
)

type ServiceConsumerAssignment service

func (s *ServiceConsumerAssignment) List(ctx context.Context, clusterID, groupID, consumerID string, resp *v3.ConsumerAssignmentListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(consumerAssignmentListEndpoint, clusterID, groupID, consumerID), resp)
}

func (s *ServiceConsumerAssignment) ListWide(ctx context.Context, clusterID, groupID, consumerID string) (presentation.ConsumerAssignmentList, error) {
	var consumersResp v3.ConsumerAssignmentListResponse
	if err := s.List(ctx, clusterID, groupID, consumerID, &consumersResp); err != nil {
		return nil, err
	}
	consumers := make(presentation.ConsumerAssignmentList, 0, len(consumersResp.Data))
	for _, cgData := range consumersResp.Data {
		cg := presentation.MapConsumerAssignment(&cgData)
		consumers = append(consumers, *cg)
	}
	return consumers, nil
}
