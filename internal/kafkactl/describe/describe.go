package describe

import (
	"github.com/devodev/kafkactl/internal/cli"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	c, err := cli.New(cli.WithLimitedOutputs("template", "json"))
	cobra.CheckErr(err)

	cmd.PersistentFlags().AddFlagSet(c.Flags())

	cmd.AddCommand(newCmdDescribeTopic(c))

	return cmd
}
