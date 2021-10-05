package config

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Interact with the kafkactl config file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringP("config-file", "f", "", "Configuration file path")

	cmd.AddCommand(newCmdConfigAddContext())
	cmd.AddCommand(newCmdConfigGetContext())
	cmd.AddCommand(newCmdConfigRemoveContext())
	cmd.AddCommand(newCmdConfigUpdateContext())
	cmd.AddCommand(newCmdConfigUseContext())

	return cmd
}
