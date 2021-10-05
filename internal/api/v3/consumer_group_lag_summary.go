package v3

type ConsumerGroupLagSummaryData struct {
	V3BaseData
	ClusterID         string            `json:"cluster_id"`
	ConsumerGroupID   string            `json:"consumer_group_id"`
	MaxLagConsumerID  string            `json:"max_lag_consumer_id"`
	MaxLagInstanceID  string            `json:"max_lag_instance_id"`
	MaxLagClientID    string            `json:"max_lag_client_id"`
	MaxLagTopicName   string            `json:"max_lag_topic_name"`
	MaxLagPartitionID int               `json:"max_lag_partition_id"`
	MaxLag            int               `json:"max_lag"`
	TotalLag          int               `json:"total_lag"`
	MaxLagConsumer    V3BaseDataRelated `json:"max_lag_consumer"`
	MaxLagPartitions  V3BaseDataRelated `json:"max_lag_partition"`
}

type ConsumerGroupLagSummaryResponse struct {
	ConsumerGroupLagSummaryData
}
