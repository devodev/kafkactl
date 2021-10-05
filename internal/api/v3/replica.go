package v3

type ReplicaData struct {
	V3BaseData
	ClusterID   string            `json:"cluster_id"`
	TopicName   string            `json:"topic_name"`
	BrokerID    int               `json:"broker_id"`
	PartitionID int               `json:"partition_id"`
	IsLeader    bool              `json:"is_leader"`
	IsInSync    bool              `json:"is_in_sync"`
	Broker      V3BaseDataRelated `json:"broker"`
}

type ReplicaListResponse struct {
	V3Base
	Data []ReplicaData `json:"data"`
}

type ReplicaGetResponse struct {
	ReplicaData
}
