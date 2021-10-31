package v3

import (
	"fmt"
	"net/url"
)

type AclOperation string

const (
	OperationUnknown         AclOperation = "UNKNOWN"
	OperationAny             AclOperation = "ANY"
	OperationAll             AclOperation = "ALL"
	OperationRead            AclOperation = "READ"
	OperationWrite           AclOperation = "WRITE"
	OperationCreate          AclOperation = "CREATE"
	OperationDelete          AclOperation = "DELETE"
	OperationAlter           AclOperation = "ALTER"
	OperationDescribe        AclOperation = "DESCRIBE"
	OperationClusterAction   AclOperation = "CLUSTER_ACTION"
	OperationDescribeConfigs AclOperation = "DESCRIBE_CONFIGS"
	OperationAlterConfigs    AclOperation = "ALTER_CONFIGS"
	OperationIdempotentWrite AclOperation = "IDEMPOTENT_WRITE"
)

var (
	AclOperationDefault = OperationAny
	AclOperationMap     = map[string]AclOperation{
		"UNKNOWN":          OperationUnknown,
		"ANY":              OperationAny,
		"ALL":              OperationAll,
		"READ":             OperationRead,
		"WRITE":            OperationWrite,
		"CREATE":           OperationCreate,
		"DELETE":           OperationDelete,
		"ALTER":            OperationAlter,
		"DESCRIBE":         OperationDescribe,
		"CLUSTER_ACTION":   OperationClusterAction,
		"DESCRIBE_CONFIGS": OperationDescribeConfigs,
		"ALTER_CONFIGS":    OperationAlterConfigs,
		"IDEMPOTENT_WRITE": OperationIdempotentWrite,
	}
)

func AclOperationFrom(s string) (AclOperation, error) {
	val, ok := AclOperationMap[s]
	if !ok {
		return "", fmt.Errorf("invalid operation: %s", s)
	}
	return val, nil
}

type AclPatternType string

const (
	PatternTypeUnknown  AclPatternType = "UNKNOWN"
	PatternTypeAny      AclPatternType = "ANY"
	PatternTypeMatch    AclPatternType = "MATCH"
	PatternTypeLiteral  AclPatternType = "LITERAL"
	PatternTypePrefixed AclPatternType = "PREFIXED"
)

var (
	AclPatternTypeDefault = PatternTypeLiteral
	AclPatternTypeMap     = map[string]AclPatternType{
		"UNKNOWN":  PatternTypeUnknown,
		"ANY":      PatternTypeAny,
		"MATCH":    PatternTypeMatch,
		"LITERAL":  PatternTypeLiteral,
		"PREFIXED": PatternTypePrefixed,
	}
)

func AclPatternTypeFrom(s string) (AclPatternType, error) {
	val, ok := AclPatternTypeMap[s]
	if !ok {
		return "", fmt.Errorf("invalid pattern type: %s", s)
	}
	return val, nil
}

type AclPermission string

const (
	PermissionUnknown AclPermission = "UNKNOWN"
	PermissionAny     AclPermission = "ANY"
	PermissionDeny    AclPermission = "DENY"
	PermissionAllow   AclPermission = "ALLOW"
)

var (
	AclPermissionDefault = PermissionAny
	AclPermissionMap     = map[string]AclPermission{
		"UNKNOWN": PermissionUnknown,
		"ANY":     PermissionAny,
		"DENY":    PermissionDeny,
		"ALLOW":   PermissionAllow,
	}
)

func AclPermissionFrom(s string) (AclPermission, error) {
	val, ok := AclPermissionMap[s]
	if !ok {
		return "", fmt.Errorf("invalid permission: %s", s)
	}
	return val, nil
}

type AclResourceType string

const (
	ResourceTypeUnknown         AclResourceType = "UNKNOWN"
	ResourceTypeAny             AclResourceType = "ANY"
	ResourceTypeTopic           AclResourceType = "TOPIC"
	ResourceTypeGroup           AclResourceType = "GROUP"
	ResourceTypeCluster         AclResourceType = "CLUSTER"
	ResourceTypeTransactionalID AclResourceType = "TRANSACTIONAL_ID"
	ResourceTypeDelegationToken AclResourceType = "DELEGATION_TOKEN"
)

var (
	AclResourceTypeDefault = ResourceTypeAny
	AclResourceTypeMap     = map[string]AclResourceType{
		"UNKNOWN":          ResourceTypeUnknown,
		"ANY":              ResourceTypeAny,
		"TOPIC":            ResourceTypeTopic,
		"GROUP":            ResourceTypeGroup,
		"CLUSTER":          ResourceTypeCluster,
		"TRANSACTIONAL_ID": ResourceTypeTransactionalID,
		"DELEGATION_TOKEN": ResourceTypeDelegationToken,
	}
)

func AclResourceTypeFrom(s string) (AclResourceType, error) {
	val, ok := AclResourceTypeMap[s]
	if !ok {
		return "", fmt.Errorf("invalid resource type: %s", s)
	}
	return val, nil
}

type AclBaseData struct {
	ResourceType AclResourceType `json:"resource_type"`
	ResourceName string          `json:"resource_name"`
	PatternType  AclPatternType  `json:"pattern_type"`
	Principal    string          `json:"principal"`
	Host         string          `json:"host"`
	Operation    AclOperation    `json:"operation"`
	Permission   AclPermission   `json:"permission"`
}

type AclData struct {
	V3BaseData
	AclBaseData
	ClusterID string `json:"cluster_id"`
}

type AclListResponse struct {
	V3Base
	Data []AclData `json:"data"`
}

type AclCreateRequest AclBaseData

type AclParams struct {
	ResourceType string
	ResourceName string
	PatternType  string
	Principal    string
	Host         string
	Operation    string
	Permission   string
}

func (q AclParams) Request() (*AclCreateRequest, error) {
	req := &AclCreateRequest{
		Operation:    AclOperationDefault,
		PatternType:  AclPatternTypeDefault,
		Permission:   AclPermissionDefault,
		ResourceType: AclResourceTypeDefault,
		ResourceName: q.ResourceName,
		Principal:    q.Principal,
		Host:         q.Host,
	}

	var err error
	if q.Operation != "" {
		req.Operation, err = AclOperationFrom(q.Operation)
		if err != nil {
			return nil, err
		}
	}
	if q.PatternType != "" {
		req.PatternType, err = AclPatternTypeFrom(q.PatternType)
		if err != nil {
			return nil, err
		}
	}
	if q.Permission != "" {
		req.Permission, err = AclPermissionFrom(q.Permission)
		if err != nil {
			return nil, err
		}
	}
	if q.ResourceType != "" {
		req.ResourceType, err = AclResourceTypeFrom(q.ResourceType)
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}

func (q AclParams) Encode() (url.Values, error) {
	queryParams := url.Values{}

	if q.Operation != "" {
		if _, err := AclOperationFrom(q.Operation); err != nil {
			return nil, err
		}
		queryParams.Set("operation", q.Operation)
	}
	if q.PatternType != "" {
		if _, err := AclPatternTypeFrom(q.PatternType); err != nil {
			return nil, err
		}
		queryParams.Set("pattern_type", q.PatternType)
	}
	if q.Permission != "" {
		if _, err := AclPermissionFrom(q.Permission); err != nil {
			return nil, err
		}
		queryParams.Set("permission", q.Permission)
	}
	if q.ResourceType != "" {
		if _, err := AclResourceTypeFrom(q.ResourceType); err != nil {
			return nil, err
		}
		queryParams.Set("resource_type", q.ResourceType)
	}

	if q.ResourceName != "" {
		queryParams.Set("resource_name", q.ResourceName)
	}
	if q.Principal != "" {
		queryParams.Set("principal", q.Principal)
	}
	if q.Host != "" {
		queryParams.Set("host", q.Host)
	}

	return queryParams, nil
}
