package client

import (
	"context"
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	brokerListEndpoint = "/v3/clusters/%s/brokers"
	brokerGetEndpoint  = "/v3/clusters/%s/brokers/%d"
)

type ServiceBroker service

func (s *ServiceBroker) List(ctx context.Context, clusterID string, resp *v3.BrokerListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(brokerListEndpoint, clusterID), resp)
}

func (s *ServiceBroker) ListWide(ctx context.Context, clusterID string) (presentation.BrokerList, error) {
	var brokersResp v3.BrokerListResponse
	if err := s.List(ctx, clusterID, &brokersResp); err != nil {
		return nil, err
	}
	brokers := make(presentation.BrokerList, 0, len(brokersResp.Data))
	for _, cgData := range brokersResp.Data {
		replicas, err := s.client.PartitionReplica.ListBrokerWide(ctx, clusterID, cgData.BrokerID)
		if err != nil {
			return nil, err
		}
		cg := presentation.MapBroker(&cgData, replicas)
		brokers = append(brokers, *cg)
	}
	return brokers, nil
}

func (s *ServiceBroker) Get(ctx context.Context, clusterID string, brokerID int, resp *v3.BrokerGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(brokerGetEndpoint, clusterID, brokerID), resp)
}

func (s *ServiceBroker) GetWide(ctx context.Context, clusterID string, brokerID int) (*presentation.Broker, error) {
	var brokerResp v3.BrokerGetResponse
	if err := s.client.Broker.Get(ctx, clusterID, brokerID, &brokerResp); err != nil {
		return nil, err
	}
	replicas, err := s.client.PartitionReplica.ListBrokerWide(ctx, clusterID, brokerID)
	if err != nil {
		return nil, err
	}
	broker := presentation.MapBroker(&brokerResp.BrokerData, replicas)
	return broker, nil
}
