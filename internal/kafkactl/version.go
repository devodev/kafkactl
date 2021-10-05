package kafkactl

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string
var Platform string

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Display the current version and platform",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("kafkactl version %s %s\n", Version, Platform)
			return nil
		},
	}

	return cmd
}
