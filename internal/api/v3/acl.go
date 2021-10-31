package v3

import "net/url"

// TODO: use the following enums to validate params
// TODO: implement json interfaces and validate value
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

type AclPatternType string

const (
	PatternTypeUnknown  AclPatternType = "UNKNOWN"
	PatternTypeAny      AclPatternType = "ANY"
	PatternTypeMatch    AclPatternType = "MATCH"
	PatternTypeLiteral  AclPatternType = "LITERAL"
	PatternTypePrefixed AclPatternType = "PREFIXED"
)

type AclPermission string

const (
	PermissionUnknown AclPermission = "UNKNOWN"
	PermissionAny     AclPermission = "ANY"
	PermissionDeny    AclPermission = "DENY"
	PermissionAllow   AclPermission = "ALLOW"
)

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

type AclData struct {
	V3BaseData
	ClusterID    string          `json:"cluster_id"`
	ResourceType AclResourceType `json:"resource_type"`
	ResourceName string          `json:"resource_name"`
	PatternType  AclPatternType  `json:"pattern_type"`
	Principal    string          `json:"principal"`
	Host         string          `json:"host"`
	Operation    AclOperation    `json:"operation"`
	Permission   AclPermission   `json:"permission"`
}

type AclListResponse struct {
	V3Base
	Data []AclData `json:"data"`
}
