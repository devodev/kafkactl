package client

import (
	"context"
	"fmt"
	"net/http"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	topicConfigListEndpoint       = "/v3/clusters/%s/topics/%s/configs"
	topicConfigGetEndpoint        = "/v3/clusters/%s/topics/%s/configs/%s"
	topicConfigBatchAlterEndpoint = "/v3/clusters/%s/topics/%s/configs:alter"
)

type ServiceTopicConfig service

func (s *ServiceTopicConfig) List(ctx context.Context, clusterID, topicName string, resp *v3.TopicConfigListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(topicConfigListEndpoint, clusterID, topicName), resp)
}

func (s *ServiceTopicConfig) ListWide(ctx context.Context, clusterID, topicName string) (presentation.TopicConfigList, error) {
	var topicConfigResp v3.TopicConfigListResponse
	if err := s.List(ctx, clusterID, topicName, &topicConfigResp); err != nil {
		return nil, err
	}
	topicConfigs := make(presentation.TopicConfigList, 0, len(topicConfigResp.Data))
	for _, topicConfigData := range topicConfigResp.Data {
		topicConfig := presentation.MapTopicConfig(&topicConfigData)
		topicConfigs = append(topicConfigs, *topicConfig)
	}
	return topicConfigs, nil
}

func (s *ServiceTopicConfig) Get(ctx context.Context, clusterID, topicName, configName string, resp *v3.TopicConfigGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(topicConfigGetEndpoint, clusterID, topicName, configName), resp)
}

func (s *ServiceTopicConfig) GetWide(ctx context.Context, clusterID, topicName, configName string) (*presentation.TopicConfig, error) {
	var topicConfigResp v3.TopicConfigGetResponse
	if err := s.Get(ctx, clusterID, topicName, configName, &topicConfigResp); err != nil {
		return nil, err
	}
	topicConfig := presentation.MapTopicConfig(&topicConfigResp.TopicConfigData)
	return topicConfig, nil
}

func (s *ServiceTopicConfig) BatchAlter(ctx context.Context, clusterID, topicName string, payload *v3.ConfigBatchAlterRequest) (string, error) {
	var statusRetriever StatusRetriever
	if err := s.client.Post(ctx, fmt.Sprintf(topicConfigBatchAlterEndpoint, clusterID, topicName), payload, nil, statusRetriever.HttpOption); err != nil {
		return "", err
	}
	if statusRetriever.Code != http.StatusNoContent {
		return "", fmt.Errorf(statusRetriever.Status)
	}
	response := fmt.Sprintf("Configs of topic with name '%s' updated/reset successfully", topicName)
	return response, nil
}
