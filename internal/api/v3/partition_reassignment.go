package v3

type PartitionReassignmentData struct {
	V3BaseData
	ClusterID        string            `json:"cluster_id"`
	TopicName        string            `json:"topic_name"`
	PartitionID      int               `json:"partition_id"`
	AddingReplicas   []int             `json:"adding_replicas"`
	RemovingReplicas []int             `json:"removing_replicas"`
	Replicas         V3BaseDataRelated `json:"replicas"`
}

type PartitionReassignmentListResponse struct {
	V3Base
	Data []PartitionReassignmentData `json:"data"`
}

type PartitionReassignmentGetResponse struct {
	PartitionReassignmentData
}
