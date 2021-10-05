package client

import (
	"context"
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	clusterListEndpoint = "/v3/clusters"
	clusterGetEndpoint  = "/v3/clusters/%s"
)

type ServiceCluster service

func (s *ServiceCluster) List(ctx context.Context, resp *v3.ClusterListResponse) error {
	return s.client.Get(ctx, clusterListEndpoint, resp)
}

func (s *ServiceCluster) ListWide(ctx context.Context) (presentation.ClusterList, error) {
	var clustersResp v3.ClusterListResponse
	if err := s.List(ctx, &clustersResp); err != nil {
		return nil, err
	}
	clusters := make(presentation.ClusterList, 0, len(clustersResp.Data))
	for _, cData := range clustersResp.Data {
		brokersResp, err := s.client.Broker.ListWide(ctx, cData.ClusterID)
		if err != nil {
			return nil, err
		}
		topicsResp, err := s.client.Topic.ListWide(ctx, cData.ClusterID)
		if err != nil {
			return nil, err
		}
		cg := presentation.MapCluster(&cData, brokersResp, topicsResp)
		clusters = append(clusters, *cg)
	}
	return clusters, nil
}

func (s *ServiceCluster) Get(ctx context.Context, clusterID string, resp *v3.ClusterGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(clusterGetEndpoint, clusterID), resp)
}

func (s *ServiceCluster) GetWide(ctx context.Context, clusterID string) (*presentation.Cluster, error) {
	var clusterResp v3.ClusterGetResponse
	if err := s.Get(ctx, clusterID, &clusterResp); err != nil {
		return nil, err
	}
	brokersResp, err := s.client.Broker.ListWide(ctx, clusterResp.ClusterID)
	if err != nil {
		return nil, err
	}
	topicsResp, err := s.client.Topic.ListWide(ctx, clusterID)
	if err != nil {
		return nil, err
	}
	cluster := presentation.MapCluster(&clusterResp.ClusterData, brokersResp, topicsResp)
	return cluster, nil
}
