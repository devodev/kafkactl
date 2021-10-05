package kafkactl

import (
	"fmt"
	"os"

	configCmd "github.com/devodev/kafkactl/internal/kafkactl/config"
	createCmd "github.com/devodev/kafkactl/internal/kafkactl/create"
	getcmd "github.com/devodev/kafkactl/internal/kafkactl/get"
	updateCmd "github.com/devodev/kafkactl/internal/kafkactl/update"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func Execute() error {
	return newRootCmd().Execute()
}

func GenDoc(filepath string) error {
	return doc.GenMarkdownTree(newRootCmd(), filepath)
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kafkactl",
		Short: "kafkactl is a CLI tool for interacting with Kafka through the confluent Kafka Rest Proxy",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentPreRunE = initLogger

	cmd.PersistentFlags().StringP("log-level", "v", log.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")

	cmd.AddCommand(configCmd.New())
	cmd.AddCommand(createCmd.New())
	cmd.AddCommand(getcmd.New())
	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(updateCmd.New())

	return cmd
}

func initLogger(cmd *cobra.Command, args []string) error {
	level, err := cmd.Flags().GetString("log-level")
	if err != nil {
		return err
	}
	log.SetOutput(os.Stdout)
	logLevel, err := log.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("could not set log level %s: %s", level, err.Error())
	}
	log.SetLevel(logLevel)
	return nil
}
