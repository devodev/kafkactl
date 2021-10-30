package client

import (
	"context"
	"fmt"
	"net/http"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	clusterConfigListEndpoint       = "/v3/clusters/%s/broker-configs"
	clusterConfigGetEndpoint        = "/v3/clusters/%s/broker-configs/%s"
	clusterConfigBatchAlterEndpoint = "/v3/clusters/%s/broker-configs:alter"
)

type ServiceClusterConfig service

func (s *ServiceClusterConfig) List(ctx context.Context, clusterID string, resp *v3.ClusterConfigListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(clusterConfigListEndpoint, clusterID), resp)
}

func (s *ServiceClusterConfig) ListWide(ctx context.Context, clusterID string) (presentation.ClusterConfigList, error) {
	var clusterConfigsResp v3.ClusterConfigListResponse
	if err := s.List(ctx, clusterID, &clusterConfigsResp); err != nil {
		return nil, err
	}
	clusterConfigs := make(presentation.ClusterConfigList, 0, len(clusterConfigsResp.Data))
	for _, cgData := range clusterConfigsResp.Data {
		cg := presentation.MapClusterConfig(&cgData)
		clusterConfigs = append(clusterConfigs, *cg)
	}
	return clusterConfigs, nil
}

func (s *ServiceClusterConfig) Get(ctx context.Context, clusterID, configName string, resp *v3.ClusterConfigGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(clusterConfigGetEndpoint, clusterID, configName), resp)
}

func (s *ServiceClusterConfig) GetWide(ctx context.Context, clusterID, configName string) (*presentation.ClusterConfig, error) {
	var clusterConfigResp v3.ClusterConfigGetResponse
	if err := s.Get(ctx, clusterID, configName, &clusterConfigResp); err != nil {
		return nil, err
	}
	clusterConfig := presentation.MapClusterConfig(&clusterConfigResp.ClusterConfigData)
	return clusterConfig, nil
}

func (s *ServiceClusterConfig) BatchAlter(ctx context.Context, clusterID string, payload *v3.ClusterConfigBatchAlterRequest) (string, error) {
	var statusRetriever StatusRetriever
	if err := s.client.Post(ctx, fmt.Sprintf(clusterConfigBatchAlterEndpoint, clusterID), payload, nil, statusRetriever.HttpOption); err != nil {
		return "", err
	}

	if statusRetriever.Code != http.StatusNoContent {
		return "", fmt.Errorf(statusRetriever.Status)
	}
	response := "Configs updated/reset successfully"
	return response, nil
}
