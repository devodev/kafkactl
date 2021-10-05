package v3

type BrokerData struct {
	V3BaseData
	ClusterID         string            `json:"cluster_id"`
	BrokerID          int               `json:"broker_id"`
	Host              string            `json:"host"`
	Port              int               `json:"port"`
	Configs           V3BaseDataRelated `json:"configs"`
	PartitionReplicas V3BaseDataRelated `json:"partition_replicas"`
}

type BrokerListResponse struct {
	V3Base
	Data []BrokerData `json:"data"`
}

type BrokerGetResponse struct {
	BrokerData
}
