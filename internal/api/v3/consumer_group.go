package v3

type ConsumerGroupData struct {
	V3BaseData
	ClusterID         string            `json:"cluster_id"`
	ConsumerGroupID   string            `json:"consumer_group_id"`
	State             string            `json:"state"`
	PartitionAssignor string            `json:"partition_assignor"`
	IsSimple          bool              `json:"is_simple"`
	Coordinator       V3BaseDataRelated `json:"coordinator"`
	Consumers         V3BaseDataRelated `json:"consumers"`
	LagSummary        V3BaseDataRelated `json:"lag_summary"`
}

type ConsumerGroupListResponse struct {
	V3Base
	Data []ConsumerGroupData `json:"data"`
}

type ConsumerGroupGetResponse struct {
	ConsumerGroupData
}
