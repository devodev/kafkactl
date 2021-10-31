package delete

import (
	"github.com/devodev/kafkactl/internal/cli"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	c, err := cli.New()
	cobra.CheckErr(err)

	cmd.PersistentFlags().AddFlagSet(c.Flags())

	cmd.AddCommand(newCmdDeleteAcl(c))
	cmd.AddCommand(newCmdDeleteTopic(c))

	return cmd
}
