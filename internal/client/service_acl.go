package client

import (
	"context"
	"fmt"
	"net/http"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	aclEndpoint = "/v3/clusters/%s/acls"
)

type ServiceAcl service

func (s *ServiceAcl) Create(ctx context.Context, clusterID string, req *v3.AclCreateRequest) (string, error) {
	endpoint := fmt.Sprintf(aclEndpoint, clusterID)

	var statusRetriever StatusRetriever
	// TODO: retrieve and display list of deleted acls
	if err := s.client.Post(ctx, endpoint, req, nil, statusRetriever.HttpOption); err != nil {
		return "", err
	}
	if statusRetriever.Code != http.StatusCreated {
		return "", fmt.Errorf(statusRetriever.Status)
	}
	response := "ACL(s) created successfully"
	return response, nil
}

func (s *ServiceAcl) Delete(ctx context.Context, clusterID string, queryParams *v3.AclParams) (string, error) {
	endpoint := fmt.Sprintf(aclEndpoint, clusterID)
	params, err := queryParams.Encode()
	if err != nil {
		return "", err
	}

	var statusRetriever StatusRetriever
	// TODO: retrieve and display list of deleted acls
	if err := s.client.DeleteWithParams(ctx, endpoint, params, nil, statusRetriever.HttpOption); err != nil {
		return "", err
	}
	if statusRetriever.Code != http.StatusOK {
		return "", fmt.Errorf(statusRetriever.Status)
	}
	response := "ACL(s) deleted successfully"
	return response, nil
}

func (s *ServiceAcl) List(ctx context.Context, clusterID string, queryParams *v3.AclParams, resp *v3.AclListResponse) error {
	endpoint := fmt.Sprintf(aclEndpoint, clusterID)
	params, err := queryParams.Encode()
	if err != nil {
		return err
	}

	return s.client.GetWithParams(ctx, endpoint, params, resp)
}

func (s *ServiceAcl) ListWide(ctx context.Context, clusterID string, queryParams *v3.AclParams) (presentation.AclList, error) {
	var aclResp v3.AclListResponse
	if err := s.List(ctx, clusterID, queryParams, &aclResp); err != nil {
		return nil, err
	}

	acls := make(presentation.AclList, 0, len(aclResp.Data))
	for _, aclData := range aclResp.Data {
		acl := presentation.MapAcl(&aclData)
		acls = append(acls, *acl)
	}
	return acls, nil
}
