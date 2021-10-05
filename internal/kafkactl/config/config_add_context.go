package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devodev/kafkactl/internal/config"
	"github.com/spf13/cobra"
)

type configAddContextOptions struct {
	ctxName string
	baseURL string

	cfgFilename string

	clusterID string
}

func newCmdConfigAddContext() *cobra.Command {
	o := configAddContextOptions{}

	cmd := &cobra.Command{
		Use:   "add-context NAME BASE_URL",
		Short: "Add a new context",
		Args:  cobra.ExactArgs(2),
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

	cmd.Flags().StringVar(&o.clusterID, "cluster-id", "", "Kafka Rest Proxy Cluster ID")

	return cmd
}

func (o *configAddContextOptions) setup(cmd *cobra.Command, args []string) error {
	o.ctxName = args[0]
	o.baseURL = args[1]

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

func (o *configAddContextOptions) validate() error { return nil }

func (o *configAddContextOptions) run() error {
	cfg, err := config.LoadFromFile(o.cfgFilename)
	if err != nil {
		if _, ok := err.(config.FileNotFoundError); !ok {
			return err
		}
		cfg = config.New()
	}

	newCtx := &config.Context{
		Name:      o.ctxName,
		BaseURL:   o.baseURL,
		ClusterID: o.clusterID,
	}
	if err := cfg.AddContext(newCtx); err != nil {
		return err
	}

	if err := config.WriteToFile(cfg, o.cfgFilename); err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Context '%s' added successfully\n", o.ctxName)
	return nil
}
