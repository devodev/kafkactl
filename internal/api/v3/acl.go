package v3

type AclData struct {
	V3BaseData
	ClusterID    string `json:"cluster_id"`
	ResourceType string `json:"resource_type"`
	ResourceName string `json:"resource_name"`
	PatternType  string `json:"pattern_type"`
	Principal    string `json:"principal"`
	Host         string `json:"host"`
	Operation    string `json:"operation"`
	Permission   string `json:"permission"`
}

type AclListResponse struct {
	V3Base
	Data []AclData `json:"data"`
}
