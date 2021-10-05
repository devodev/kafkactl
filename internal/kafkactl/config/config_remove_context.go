package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devodev/kafkactl/internal/config"
	"github.com/spf13/cobra"
)

type configRemoveContextOptions struct {
	ctxName string

	cfgFilename string
}

func newCmdConfigRemoveContext() *cobra.Command {
	o := configRemoveContextOptions{}

	cmd := &cobra.Command{
		Use:   "remove-context NAME",
		Short: "Remove a context",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.setup(cmd, args); err != nil {
				return err
			}
			if err := o.validate(); err != nil {
				return err
			}
			if err := o.run(); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}

func (o *configRemoveContextOptions) setup(cmd *cobra.Command, args []string) error {
	o.ctxName = args[0]

	cfgFilename, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return err
	}

	// if flag not set, load config using default filepath
	if cfgFilename == "" {
		var home string
		home, err = os.UserHomeDir()
		if err != nil {
			return err
		}
		cfgFilename = filepath.Join(home, ".kafkactl.yaml")
	}

	o.cfgFilename = cfgFilename

	return nil
}

func (o *configRemoveContextOptions) validate() error { return nil }

func (o *configRemoveContextOptions) run() error {
	cfg, err := config.LoadFromFile(o.cfgFilename)
	if err != nil {
		return err
	}

	if err := cfg.RemoveContext(o.ctxName); err != nil {
		return err
	}

	if err := config.WriteToFile(cfg, o.cfgFilename); err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Context '%s' removed successfully\n", o.ctxName)
	return nil
}
