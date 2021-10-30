package get

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Display resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringP("output", "o", "table", "How to format the output (table, json)")
	cmd.PersistentFlags().StringP("config-file", "f", "", "Configuration file path")
	cmd.PersistentFlags().StringArrayP("header", "H", []string{}, "Additional HTTP header(s)")

	cmd.AddCommand(newCmdGetAcl())
	cmd.AddCommand(newCmdGetBroker())
	cmd.AddCommand(newCmdGetBrokerConfig())
	cmd.AddCommand(newCmdGetCluster())
	cmd.AddCommand(newCmdGetClusterConfig())
	cmd.AddCommand(newCmdGetConsumer())
	cmd.AddCommand(newCmdGetConsumerAssignment())
	cmd.AddCommand(newCmdGetConsumerLag())
	cmd.AddCommand(newCmdGetConsumerGroup())
	cmd.AddCommand(newCmdGetConsumerGroupLagSummary())
	cmd.AddCommand(newCmdGetPartition())
	cmd.AddCommand(newCmdGetPartitionReassignment())
	cmd.AddCommand(newCmdGetPartitionReplica())
	cmd.AddCommand(newCmdGetPartitionReplicaBroker())
	cmd.AddCommand(newCmdGetTopic())
	cmd.AddCommand(newCmdGetTopicConfig())

	return cmd
}
