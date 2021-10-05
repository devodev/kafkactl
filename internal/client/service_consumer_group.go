package client

import (
	"context"
	"fmt"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	consumerGroupListEndpoint       = "/v3/clusters/%s/consumer-groups"
	consumerGroupGetEndpoint        = "/v3/clusters/%s/consumer-groups/%s"
	consumerGroupLagSummaryEndpoint = "/v3/clusters/%s/consumer-groups/%s/lag-summary"
)

type ServiceConsumerGroup service

func (s *ServiceConsumerGroup) List(ctx context.Context, clusterID string, resp *v3.ConsumerGroupListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(consumerGroupListEndpoint, clusterID), resp)
}

func (s *ServiceConsumerGroup) ListWide(ctx context.Context, clusterID string) (presentation.ConsumerGroupList, error) {
	var consumerGroupsResp v3.ConsumerGroupListResponse
	if err := s.List(ctx, clusterID, &consumerGroupsResp); err != nil {
		return nil, err
	}
	consumerGroups := make(presentation.ConsumerGroupList, 0, len(consumerGroupsResp.Data))
	for _, cgData := range consumerGroupsResp.Data {
		consumers, err := s.client.Consumer.ListWide(ctx, clusterID, cgData.ConsumerGroupID)
		if err != nil {
			return nil, err
		}
		cg := presentation.MapConsumerGroup(&cgData, consumers)
		consumerGroups = append(consumerGroups, *cg)
	}
	return consumerGroups, nil
}

func (s *ServiceConsumerGroup) Get(ctx context.Context, clusterID, groupID string, resp *v3.ConsumerGroupGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(consumerGroupGetEndpoint, clusterID, groupID), resp)
}

func (s *ServiceConsumerGroup) GetWide(ctx context.Context, clusterID, groupID string) (*presentation.ConsumerGroup, error) {
	var consumerGroupResp v3.ConsumerGroupGetResponse
	if err := s.Get(ctx, clusterID, groupID, &consumerGroupResp); err != nil {
		return nil, err
	}
	consumers, err := s.client.Consumer.ListWide(ctx, clusterID, groupID)
	if err != nil {
		return nil, err
	}
	consumerGroup := presentation.MapConsumerGroup(&consumerGroupResp.ConsumerGroupData, consumers)
	return consumerGroup, nil
}

func (s *ServiceConsumerGroup) LagSummary(ctx context.Context, clusterID, groupID string, resp *v3.ConsumerGroupLagSummaryResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(consumerGroupLagSummaryEndpoint, clusterID, groupID), resp)
}

func (s *ServiceConsumerGroup) LagSummaryWide(ctx context.Context, clusterID, groupID string) (*presentation.ConsumerGroupLagSummary, error) {
	var lagSummaryResp v3.ConsumerGroupLagSummaryResponse
	if err := s.LagSummary(ctx, clusterID, groupID, &lagSummaryResp); err != nil {
		return nil, err
	}

	lagSummary := presentation.MapConsumerGroupLagSummary(&lagSummaryResp.ConsumerGroupLagSummaryData)
	return lagSummary, nil
}

func (s *ServiceConsumerGroup) LagSummaryAllWide(ctx context.Context, clusterID string) (presentation.ConsumerGroupLagSummaryList, error) {
	var consumerGroupsResp v3.ConsumerGroupListResponse
	if err := s.List(ctx, clusterID, &consumerGroupsResp); err != nil {
		return nil, err
	}

	lagSummaries := make(presentation.ConsumerGroupLagSummaryList, 0, len(consumerGroupsResp.Data))
	for _, cgData := range consumerGroupsResp.Data {
		lagSummary, err := s.LagSummaryWide(ctx, clusterID, cgData.ConsumerGroupID)
		if err != nil {
			return nil, err
		}
		lagSummaries = append(lagSummaries, *lagSummary)
	}
	return lagSummaries, nil
}
