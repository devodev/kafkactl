package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devodev/kafkactl/internal/config"
	"github.com/spf13/cobra"
)

type configUseContextOptions struct {
	ctxName string

	cfgFilename string
}

func newCmdConfigUseContext() *cobra.Command {
	o := configUseContextOptions{}

	cmd := &cobra.Command{
		Use:   "use-context NAME",
		Short: "Set current context",
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

func (o *configUseContextOptions) setup(cmd *cobra.Command, args []string) error {
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

func (o *configUseContextOptions) validate() error { return nil }

func (o *configUseContextOptions) run() error {
	cfg, err := config.LoadFromFile(o.cfgFilename)
	if err != nil {
		return err
	}
	if _, err := cfg.GetContext(o.ctxName); err != nil {
		return err
	}

	cfg.SetCurrentContext(o.ctxName)

	if err := config.WriteToFile(cfg, o.cfgFilename); err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Current context set to '%s' successfully\n", o.ctxName)
	return nil
}
