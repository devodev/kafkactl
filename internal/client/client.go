package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
)

type service struct {
	client *KafkaRest
}

type HttpOption func(*http.Request, *http.Response) error

type StatusRetriever struct {
	Status string
	Code   int
}

func (s *StatusRetriever) HttpOption(req *http.Request, resp *http.Response) error {
	s.Status = resp.Status
	s.Code = resp.StatusCode
	return nil
}

type kafkaRestOption func(*KafkaRest) error

func WithHeader(header, value string) kafkaRestOption {
	return func(c *KafkaRest) error {
		if _, ok := c.headers[header]; ok {
			return fmt.Errorf("header '%s' already set", header)
		}
		c.headers[header] = value
		return nil
	}
}

func WithHeaders(headers map[string]string) kafkaRestOption {
	return func(c *KafkaRest) error {
		for h, v := range headers {
			if _, ok := c.headers[h]; ok {
				return fmt.Errorf("header '%s' already set", h)
			}
			c.headers[h] = v
		}
		return nil
	}
}

type KafkaRest struct {
	client *http.Client

	BaseURL *url.URL

	headers map[string]string

	// inspired by go-github:
	// https://github.com/google/go-github/blob/d913de9ce1e8ed5550283b448b37b721b61cc3b3/github/github.go#L159
	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	Acl                   *ServiceAcl
	Broker                *ServiceBroker
	BrokerConfig          *ServiceBrokerConfig
	Cluster               *ServiceCluster
	ClusterConfig         *ServiceClusterConfig
	Consumer              *ServiceConsumer
	ConsumerAssignment    *ServiceConsumerAssignment
	ConsumerGroup         *ServiceConsumerGroup
	ConsumerLag           *ServiceConsumerLag
	Partition             *ServicePartition
	PartitionReassignment *ServicePartitionReassignment
	PartitionReplica      *ServicePartitionReplica
	Topic                 *ServiceTopic
	TopicConfig           *ServiceTopicConfig
}

func New(baseURL string, opts ...kafkaRestOption) (*KafkaRest, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("could not parse base url: %w", err)
	}

	client := &KafkaRest{
		BaseURL: url,
		headers: map[string]string{
			"Content-Type": "application/json",
		},
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: false,
			},
			Timeout: 30 * time.Second,
		},
	}

	client.common.client = client

	client.Acl = (*ServiceAcl)(&client.common)
	client.Broker = (*ServiceBroker)(&client.common)
	client.BrokerConfig = (*ServiceBrokerConfig)(&client.common)
	client.Cluster = (*ServiceCluster)(&client.common)
	client.ClusterConfig = (*ServiceClusterConfig)(&client.common)
	client.Consumer = (*ServiceConsumer)(&client.common)
	client.ConsumerAssignment = (*ServiceConsumerAssignment)(&client.common)
	client.ConsumerGroup = (*ServiceConsumerGroup)(&client.common)
	client.ConsumerLag = (*ServiceConsumerLag)(&client.common)
	client.PartitionReassignment = (*ServicePartitionReassignment)(&client.common)
	client.PartitionReplica = (*ServicePartitionReplica)(&client.common)
	client.Partition = (*ServicePartition)(&client.common)
	client.Topic = (*ServiceTopic)(&client.common)
	client.TopicConfig = (*ServiceTopicConfig)(&client.common)

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func (c *KafkaRest) Get(ctx context.Context, endpoint string, result interface{}, opts ...HttpOption) error {
	return c.roundtrip(ctx, "GET", endpoint, nil, result, opts...)
}

func (c *KafkaRest) Post(ctx context.Context, endpoint string, payload interface{}, result interface{}, opts ...HttpOption) error {
	return c.roundtrip(ctx, "POST", endpoint, payload, result, opts...)
}

func (c *KafkaRest) roundtrip(ctx context.Context, method, endpoint string, payload interface{}, result interface{}, opts ...HttpOption) error {
	log.WithField("source", "KafkaRest.roundtrip").Debugf("%s %s", method, endpoint)
	req, err := c.makeRequest(ctx, method, endpoint, payload)
	if err != nil {
		return err
	}
	resp, err := c.sendRequest(req)
	if err != nil {
		return err
	}
	for _, opt := range opts {
		if err := opt(req, resp); err != nil {
			return fmt.Errorf("error returned from executing http option: %w", err)
		}
	}
	return c.handleResponse(resp, result)
}

func (c *KafkaRest) makeRequest(ctx context.Context, method, endpoint string, payload interface{}) (*http.Request, error) {
	url := *c.BaseURL
	url.Path = path.Join(url.Path, endpoint)
	reqURL := url.String()

	jsonPayload := new(bytes.Buffer)
	if payload != nil {
		if err := json.NewEncoder(jsonPayload).Encode(payload); err != nil {
			return nil, fmt.Errorf("could not encode payload: %w", err)
		}
	}

	funcLog := log.WithField("source", "KafkaRest.makeRequest")
	funcLog.Tracef("makeRequest (%s) [url] %s", method, reqURL)
	funcLog.Tracef("makeRequest (%s) [payload] %s", method, jsonPayload.String())

	req, err := http.NewRequest(method, reqURL, jsonPayload)
	if err != nil {
		return nil, fmt.Errorf("could not create http request (%s): %w", reqURL, err)
	}

	for header, value := range c.headers {
		req.Header.Add(header, value)
	}

	req = req.WithContext(ctx)

	funcLog.Tracef("makeRequest (%s) [req.Header] %+v", method, req.Header)

	return req, nil
}

func (c *KafkaRest) sendRequest(req *http.Request) (*http.Response, error) {
	log.WithField("source", "KafkaRest.sendRequest").Trace()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	return resp, nil
}

func (c *KafkaRest) handleResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}

	funcLog := log.WithField("source", "KafkaRest.handleResponse")
	funcLog.Tracef("handleResponse (%d) %s [body] %s", resp.StatusCode, resp.Status, string(body))

	if len(body) == 0 {
		funcLog.Debugf("handleResponse will not decode empty body")
		return nil
	}
	if result == nil {
		funcLog.Debugf("handleResponse will not decode because nil passed as result receiver")
		return nil
	}

	jsonDecoder := json.NewDecoder(bytes.NewReader(body))
	// handle http error
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var clientError ClientError
		if err = jsonDecoder.Decode(&clientError); err == nil {
			return clientError
		}
		return fmt.Errorf("unknown error (%d): %s", resp.StatusCode, resp.Status)
	}

	if err := jsonDecoder.Decode(&result); err != nil {
		return fmt.Errorf("could not decode json response: %w", err)
	}
	return nil
}
