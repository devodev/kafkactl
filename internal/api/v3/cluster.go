package v3

type ClusterData struct {
	V3BaseData
	ClusterID              string            `json:"cluster_id"`
	Acls                   V3BaseDataRelated `json:"acls"`
	Brokers                V3BaseDataRelated `json:"brokers"`
	BrokerConfigs          V3BaseDataRelated `json:"broker_configs"`
	ConsumerGroups         V3BaseDataRelated `json:"consumer_groups"`
	Controller             V3BaseDataRelated `json:"controller"`
	PartitionReassignments V3BaseDataRelated `json:"partition_reassignments"`
	Topics                 V3BaseDataRelated `json:"topics"`
}

type ClusterListResponse struct {
	V3Base
	Data []ClusterData `json:"data"`
}

type ClusterGetResponse struct {
	ClusterData
}
