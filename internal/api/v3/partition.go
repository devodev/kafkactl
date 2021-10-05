package v3

type PartitionData struct {
	V3BaseData
	ClusterID    string            `json:"cluster_id"`
	TopicName    string            `json:"topic_name"`
	PartitionID  int               `json:"partition_id"`
	Leader       V3BaseDataRelated `json:"leader"`
	Replicas     V3BaseDataRelated `json:"replicas"`
	Reassignment V3BaseDataRelated `json:"reassignment"`
}

type PartitionListResponse struct {
	V3Base
	Data []PartitionData `json:"data"`
}

type PartitionGetResponse struct {
	PartitionData
}
