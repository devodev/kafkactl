package create

import (
	"github.com/devodev/kafkactl/internal/cli"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	c := cli.New()

	cmd.PersistentFlags().AddFlagSet(c.Flags())

	cmd.AddCommand(newCmdCreateAcl(c))
	cmd.AddCommand(newCmdCreateTopic(c))

	return cmd
}
