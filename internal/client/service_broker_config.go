package client

import (
	"context"
	"fmt"
	"net/http"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	brokerConfigListEndpoint       = "/v3/clusters/%s/brokers/%d/configs"
	brokerConfigGetEndpoint        = "/v3/clusters/%s/brokers/%d/configs/%s"
	brokerConfigBatchAlterEndpoint = "/v3/clusters/%s/brokers/%d/configs:alter"
)

type ServiceBrokerConfig service

func (s *ServiceBrokerConfig) List(ctx context.Context, clusterID string, brokerID int, resp *v3.BrokerConfigListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(brokerConfigListEndpoint, clusterID, brokerID), resp)
}

func (s *ServiceBrokerConfig) ListWide(ctx context.Context, clusterID string, brokerID int) (presentation.BrokerConfigList, error) {
	var brokerConfigsResp v3.BrokerConfigListResponse
	if err := s.List(ctx, clusterID, brokerID, &brokerConfigsResp); err != nil {
		return nil, err
	}
	brokerConfigs := make(presentation.BrokerConfigList, 0, len(brokerConfigsResp.Data))
	for _, cgData := range brokerConfigsResp.Data {
		cg := presentation.MapBrokerConfig(&cgData)
		brokerConfigs = append(brokerConfigs, *cg)
	}
	return brokerConfigs, nil
}

func (s *ServiceBrokerConfig) ListAllWide(ctx context.Context, clusterID string) (presentation.BrokerConfigList, error) {
	brokerResp, err := s.client.Broker.ListWide(ctx, clusterID)
	if err != nil {
		return nil, err
	}
	brokerConfigs := make(presentation.BrokerConfigList, 0, len(brokerResp))
	for _, bData := range brokerResp {
		bcs, err := s.ListWide(ctx, clusterID, bData.BrokerID)
		if err != nil {
			return nil, err
		}
		brokerConfigs = append(brokerConfigs, bcs...)
	}
	return brokerConfigs, nil
}

func (s *ServiceBrokerConfig) Get(ctx context.Context, clusterID string, brokerID int, configName string, resp *v3.BrokerConfigGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(brokerConfigGetEndpoint, clusterID, brokerID, configName), resp)
}

func (s *ServiceBrokerConfig) GetWide(ctx context.Context, clusterID string, brokerID int, configName string) (*presentation.BrokerConfig, error) {
	var brokerConfigResp v3.BrokerConfigGetResponse
	if err := s.Get(ctx, clusterID, brokerID, configName, &brokerConfigResp); err != nil {
		return nil, err
	}
	brokerConfig := presentation.MapBrokerConfig(&brokerConfigResp.BrokerConfigData)
	return brokerConfig, nil
}

func (s *ServiceBrokerConfig) BatchAlter(ctx context.Context, clusterID string, brokerID int, payload *v3.BrokerConfigBatchAlterRequest) (string, error) {
	var statusRetriever StatusRetriever
	if err := s.client.Post(ctx, fmt.Sprintf(brokerConfigBatchAlterEndpoint, clusterID, brokerID), payload, nil, statusRetriever.HttpOption); err != nil {
		return "", err
	}

	if statusRetriever.Code != http.StatusNoContent {
		return "", fmt.Errorf(statusRetriever.Status)
	}
	response := fmt.Sprintf("Configs of broker with ID '%d' updated/reset successfully", brokerID)
	return response, nil
}
