package create

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringP("output", "o", "table", "How to format the output (table, json)")
	cmd.PersistentFlags().StringP("config-file", "f", "", "Configuration file path")

	cmd.AddCommand(newCmdCreateTopic())

	return cmd
}
