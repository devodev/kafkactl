package presentation

import (
	"net/url"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	aclHeader        = []string{"resource_type", "resource_name", "pattern_type", "principal", "host", "operation", "permission"}
	aclHeaderDataMap = func(data Acl) map[string]string {
		return map[string]string{
			"resource_type": data.ResourceType,
			"resource_name": data.ResourceName,
			"pattern_type":  data.PatternType,
			"principal":     data.Principal,
			"host":          data.Host,
			"operation":     data.Operation,
			"permission":    data.Permission,
		}
	}
)

func MapAcl(data *v3.AclData) *Acl {
	return &Acl{
		ClusterID:    data.ClusterID,
		ResourceType: data.ResourceType,
		ResourceName: data.ResourceName,
		PatternType:  data.PatternType,
		Principal:    data.Principal,
		Host:         data.Host,
		Operation:    data.Operation,
		Permission:   data.Permission,
	}
}

type Acl struct {
	ClusterID    string `json:"cluster_id"`
	ResourceType string `json:"resource_type"`
	ResourceName string `json:"resource_name"`
	PatternType  string `json:"pattern_type"`
	Principal    string `json:"principal"`
	Host         string `json:"host"`
	Operation    string `json:"operation"`
	Permission   string `json:"permission"`
}

func (t Acl) TableHeader() []string {
	return aclHeader
}

func (t Acl) TableRows() []map[string]string {
	rows := []map[string]string{aclHeaderDataMap(t)}
	return rows
}

type AclList []Acl

func (t AclList) TableHeader() []string {
	return aclHeader
}

func (t AclList) TableRows() []map[string]string {
	rows := make([]map[string]string, len(t))
	for idx, data := range t {
		rows[idx] = aclHeaderDataMap(data)
	}
	return rows
}

type AclListQueryParams struct {
	ResourceType string
	ResourceName string
	PatternType  string
	Principal    string
	Host         string
	Operation    string
	Permission   string
}

func (q AclListQueryParams) Encode() string {
	queryParams := url.Values{}
	if q.ResourceType != "" {
		queryParams.Add("resource_type", q.ResourceType)
	}
	if q.ResourceName != "" {
		queryParams.Add("resource_name", q.ResourceName)
	}
	if q.PatternType != "" {
		queryParams.Add("pattern_type", q.PatternType)
	}
	if q.Principal != "" {
		queryParams.Add("principal", q.Principal)
	}
	if q.Host != "" {
		queryParams.Add("host", q.Host)
	}
	if q.Operation != "" {
		queryParams.Add("operation", q.Operation)
	}
	if q.Permission != "" {
		queryParams.Add("permission", q.Permission)
	}
	return queryParams.Encode()
}
