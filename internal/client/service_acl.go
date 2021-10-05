package client

import (
	"context"
	"fmt"
	"strings"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/presentation"
)

const (
	aclListEndpoint = "/v3/clusters/%s/acls"
)

type ServiceAcl service

func (s *ServiceAcl) List(ctx context.Context, clusterID string, queryParams *presentation.AclListQueryParams, resp *v3.AclListResponse) error {
	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf(aclListEndpoint, clusterID))

	urlQuery := queryParams.Encode()
	if urlQuery != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", urlQuery))
	}

	return s.client.Get(ctx, endpoint.String(), resp)
}

func (s *ServiceAcl) ListWide(ctx context.Context, clusterID string, queryParams *presentation.AclListQueryParams) (presentation.AclList, error) {
	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf(aclListEndpoint, clusterID))

	urlQuery := queryParams.Encode()
	if urlQuery != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", urlQuery))
	}

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
