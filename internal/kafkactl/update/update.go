package update

import (
	"github.com/devodev/kafkactl/internal/cli"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	c := cli.New()

	cmd.PersistentFlags().AddFlagSet(c.Flags())

	cmd.AddCommand(newCmdupdateBrokerConfig(c))
	cmd.AddCommand(newCmdUpdateTopicConfig(c))

	return cmd
}
