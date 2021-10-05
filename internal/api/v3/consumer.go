package v3

type ConsumerData struct {
	V3BaseData
	ClusterID       string            `json:"cluster_id"`
	ConsumerGroupID string            `json:"consumer_group_id"`
	ConsumerID      string            `json:"consumer_id"`
	InstanceID      string            `json:"instance_id"`
	ClientID        string            `json:"client_id"`
	Assignments     V3BaseDataRelated `json:"assignments"`
}

type ConsumerListResponse struct {
	V3Base
	Data []ConsumerData `json:"data"`
}

type ConsumerGetResponse struct {
	ConsumerData
}
