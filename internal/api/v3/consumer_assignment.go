package v3

type ConsumerAssignmentData struct {
	V3BaseData
	ClusterID       string            `json:"cluster_id"`
	ConsumerGroupID string            `json:"consumer_group_id"`
	ConsumerID      string            `json:"consumer_id"`
	TopicName       string            `json:"topic_name"`
	PartitionID     int               `json:"partition_id"`
	Partition       V3BaseDataRelated `json:"partition"`
	Lag             V3BaseDataRelated `json:"lag"`
}

type ConsumerAssignmentListResponse struct {
	V3Base
	Data []ConsumerAssignmentData `json:"data"`
}

type ConsumerAssignmentGetResponse struct {
	ConsumerAssignmentData
}
