package update

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringP("output", "o", "table", "How to format the output (table, json)")
	cmd.PersistentFlags().StringP("config-file", "f", "", "Configuration file path")

	cmd.AddCommand(newCmdUpdateTopicConfig())

	return cmd
}
