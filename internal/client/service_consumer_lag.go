package client

import (
	"context"
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	consumerLagListEndpoint = "/v3/clusters/%s/consumer-groups/%s/lags"
)

type ServiceConsumerLag service

func (s *ServiceConsumerLag) List(ctx context.Context, clusterID, groupID string, resp *v3.ConsumerLagListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(consumerLagListEndpoint, clusterID, groupID), resp)
}

func (s *ServiceConsumerLag) ListWide(ctx context.Context, clusterID, groupID string) (presentation.ConsumerLagList, error) {
	var consumersResp v3.ConsumerLagListResponse
	if err := s.List(ctx, clusterID, groupID, &consumersResp); err != nil {
		return nil, err
	}
	consumers := make(presentation.ConsumerLagList, 0, len(consumersResp.Data))
	for _, cgData := range consumersResp.Data {
		cg := presentation.MapConsumerLag(&cgData)
		consumers = append(consumers, *cg)
	}
	return consumers, nil
}
