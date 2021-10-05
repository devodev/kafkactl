package v3

type ConsumerLagData struct {
	V3BaseData
	ClusterID       string `json:"cluster_id"`
	ConsumerGroupID string `json:"consumer_group_id"`
	ConsumerID      string `json:"consumer_id"`
	InstanceID      string `json:"instance_id"`
	ClientID        string `json:"client_id"`
	TopicName       string `json:"topic_name"`
	PartitionID     int    `json:"partition_id"`
	CurrentOffset   int    `json:"current_offset"`
	LogEndOffset    int    `json:"log_end_offset"`
	Lag             int    `json:"lag"`
}

type ConsumerLagListResponse struct {
	V3Base
	Data []ConsumerLagData `json:"data"`
}

type ConsumerLagGetResponse struct {
	ConsumerLagData
}
