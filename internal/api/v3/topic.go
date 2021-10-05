package v3

type TopicData struct {
	V3BaseData
	ClusterID              string            `json:"cluster_id"`
	TopicName              string            `json:"topic_name"`
	ReplicationFactor      int               `json:"replication_factor"`
	IsInternal             bool              `json:"is_internal"`
	Partitions             V3BaseDataRelated `json:"partitions"`
	Configs                V3BaseDataRelated `json:"configs"`
	PartitionReassignments V3BaseDataRelated `json:"partition_reassignments"`
}

type TopicListResponse struct {
	V3Base
	Data []TopicData `json:"data"`
}

type TopicGetResponse struct {
	TopicData
}

type TopicConfig struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TopicCreateRequest struct {
	TopicName           string `json:"topic_name"`
	PartitionsCount     *int   `json:"partitions_count"`
	ReplicationFactor   *int   `json:"replication_factor"`
	ReplicasAssignments []struct {
		PartitionID int   `json:"partition_id"`
		BrokerIDs   []int `json:"broker_ids"`
	} `json:"replicas_assignments"`
	Configs []TopicConfig `json:"configs"`
}
