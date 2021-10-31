package presentation

import (
	"net/url"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
)

var (
	aclHeader        = []string{"principal", "host", "pattern_type", "resource_type", "resource_name", "permission", "operation"}
	aclHeaderDataMap = func(data Acl) map[string]string {
		return map[string]string{
			"principal":     data.Principal,
			"host":          data.Host,
			"pattern_type":  data.PatternType,
			"resource_type": data.ResourceType,
			"resource_name": data.ResourceName,
			"permission":    data.Permission,
			"operation":     data.Operation,
		}
	}
)

func MapAcl(data *v3.AclData) *Acl {
	return &Acl{
		ClusterID:    data.ClusterID,
		Principal:    data.Principal,
		Host:         data.Host,
		ResourceType: string(data.ResourceType),
		PatternType:  string(data.PatternType),
		ResourceName: data.ResourceName,
		Permission:   string(data.Permission),
		Operation:    string(data.Operation),
	}
}

type Acl struct {
	ClusterID    string `json:"cluster_id"`
	Principal    string `json:"principal"`
	Host         string `json:"host"`
	ResourceType string `json:"resource_type"`
	PatternType  string `json:"pattern_type"`
	ResourceName string `json:"resource_name"`
	Permission   string `json:"permission"`
	Operation    string `json:"operation"`
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
	Principal    string
	Host         string
	ResourceType string
	PatternType  string
	ResourceName string
	Permission   string
	Operation    string
}

func (q AclListQueryParams) Encode() string {
	queryParams := url.Values{}
	if q.Principal != "" {
		queryParams.Add("principal", q.Principal)
	}
	if q.Host != "" {
		queryParams.Add("host", q.Host)
	}
	if q.ResourceType != "" {
		queryParams.Add("resource_type", q.ResourceType)
	}
	if q.PatternType != "" {
		queryParams.Add("pattern_type", q.PatternType)
	}
	if q.ResourceName != "" {
		queryParams.Add("resource_name", q.ResourceName)
	}
	if q.Permission != "" {
		queryParams.Add("permission", q.Permission)
	}
	if q.Operation != "" {
		queryParams.Add("operation", q.Operation)
	}
	return queryParams.Encode()
}
