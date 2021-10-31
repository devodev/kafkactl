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

func (s *ServiceAcl) Delete(ctx context.Context, clusterID string, queryParams *v3.AclQueryParams) (string, error) {
	endpoint := fmt.Sprintf(aclEndpoint, clusterID)
	params := queryParams.Encode()

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

func (s *ServiceAcl) List(ctx context.Context, clusterID string, queryParams *v3.AclQueryParams, resp *v3.AclListResponse) error {
	endpoint := fmt.Sprintf(aclEndpoint, clusterID)
	params := queryParams.Encode()

	return s.client.GetWithParams(ctx, endpoint, params, resp)
}

func (s *ServiceAcl) ListWide(ctx context.Context, clusterID string, queryParams *v3.AclQueryParams) (presentation.AclList, error) {
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
