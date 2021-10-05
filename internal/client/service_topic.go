package client

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	// endpoints
	topicCreateEndpoint = "/v3/clusters/%s/topics"
	topicListEndpoint   = "/v3/clusters/%s/topics"
	topicGetEndpoint    = "/v3/clusters/%s/topics/%s"
)

type ServiceTopic service

func (s *ServiceTopic) Create(ctx context.Context, clusterID string, payload *v3.TopicCreateRequest) (string, error) {
	var statusRetriever StatusRetriever
	if err := s.client.Post(ctx, fmt.Sprintf(topicCreateEndpoint, clusterID), payload, nil, statusRetriever.HttpOption); err != nil {
		return "", err
	}
	response := statusRetriever.Status
	if statusRetriever.Code == http.StatusCreated {
		response = fmt.Sprintf("Topic %s created successfully", payload.TopicName)
	}
	return response, nil
}

func (s *ServiceTopic) List(ctx context.Context, clusterID string, resp *v3.TopicListResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(topicListEndpoint, clusterID), resp)
}

func (s *ServiceTopic) ListWide(ctx context.Context, clusterID string) (presentation.TopicList, error) {
	var topicsResp v3.TopicListResponse
	if err := s.List(ctx, clusterID, &topicsResp); err != nil {
		return nil, err
	}
	topics := make(presentation.TopicList, 0, len(topicsResp.Data))
	for _, topicData := range topicsResp.Data {
		partitionResp, err := s.client.Partition.ListWide(ctx, clusterID, topicData.TopicName)
		if err != nil {
			return nil, err
		}
		configResp, err := s.client.TopicConfig.ListWide(ctx, clusterID, topicData.TopicName)
		if err != nil {
			return nil, err
		}
		topic := presentation.MapTopic(&topicData, partitionResp, configResp)
		topics = append(topics, *topic)
	}
	sort.Sort(presentation.TopicAlphabeticSort(topics))
	return topics, nil
}

func (s *ServiceTopic) Get(ctx context.Context, clusterID, topicName string, resp *v3.TopicGetResponse) error {
	return s.client.Get(ctx, fmt.Sprintf(topicGetEndpoint, clusterID, topicName), resp)
}

func (s *ServiceTopic) GetWide(ctx context.Context, clusterID, topicName string) (*presentation.Topic, error) {
	var topicResp v3.TopicGetResponse
	if err := s.Get(ctx, clusterID, topicName, &topicResp); err != nil {
		return nil, err
	}
	partitionResp, err := s.client.Partition.ListWide(ctx, clusterID, topicResp.TopicName)
	if err != nil {
		return nil, err
	}
	configResp, err := s.client.TopicConfig.ListWide(ctx, clusterID, topicResp.TopicName)
	if err != nil {
		return nil, err
	}
	topic := presentation.MapTopic(&topicResp.TopicData, partitionResp, configResp)
	return topic, nil
}
