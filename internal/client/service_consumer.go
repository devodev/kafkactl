package client

import (
	"context"
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	consumerListEndpoint = "/v3/clusters/%s/consumer-groups/%s/consumers"
	consumerGetEndpoint  = "/v3/clusters/%s/consumer-groups/%s/consumers/%s"
)

type ServiceConsumer service

func (s *ServiceConsumer) List(ctx context.Context, clusterID, groupID string, resp *v3.ConsumerListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(consumerListEndpoint, clusterID, groupID), resp)
}

func (s *ServiceConsumer) ListWide(ctx context.Context, clusterID, groupID string) (presentation.ConsumerList, error) {
	var consumersResp v3.ConsumerListResponse
	if err := s.List(ctx, clusterID, groupID, &consumersResp); err != nil {
		return nil, err
	}
	consumers := make(presentation.ConsumerList, 0, len(consumersResp.Data))
	for _, cgData := range consumersResp.Data {
		cg := presentation.MapConsumer(&cgData)
		consumers = append(consumers, *cg)
	}
	return consumers, nil
}

func (s *ServiceConsumer) Get(ctx context.Context, clusterID, groupID, consumerID string, resp *v3.ConsumerGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(consumerGetEndpoint, clusterID, groupID, consumerID), resp)
}

func (s *ServiceConsumer) GetWide(ctx context.Context, clusterID, groupID, consumerID string) (*presentation.Consumer, error) {
	var consumerResp v3.ConsumerGetResponse
	if err := s.Get(ctx, clusterID, groupID, consumerID, &consumerResp); err != nil {
		return nil, err
	}
	consumer := presentation.MapConsumer(&consumerResp.ConsumerData)
	return consumer, nil
}
