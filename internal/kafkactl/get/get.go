package get

import (
	"github.com/devodev/kafkactl/internal/cli"
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

	c := cli.New()

	cmd.PersistentFlags().AddFlagSet(c.Flags())

	cmd.AddCommand(newCmdGetAcl(c))
	cmd.AddCommand(newCmdGetBroker(c))
	cmd.AddCommand(newCmdGetBrokerConfig(c))
	cmd.AddCommand(newCmdGetCluster(c))
	cmd.AddCommand(newCmdGetClusterConfig(c))
	cmd.AddCommand(newCmdGetConsumer(c))
	cmd.AddCommand(newCmdGetConsumerAssignment(c))
	cmd.AddCommand(newCmdGetConsumerLag(c))
	cmd.AddCommand(newCmdGetConsumerGroup(c))
	cmd.AddCommand(newCmdGetConsumerGroupLagSummary(c))
	cmd.AddCommand(newCmdGetPartition(c))
	cmd.AddCommand(newCmdGetPartitionReassignment(c))
	cmd.AddCommand(newCmdGetPartitionReplica(c))
	cmd.AddCommand(newCmdGetPartitionReplicaBroker(c))
	cmd.AddCommand(newCmdGetTopic(c))
	cmd.AddCommand(newCmdGetTopicConfig(c))

	return cmd
}
