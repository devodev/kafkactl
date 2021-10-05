package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devodev/kafkactl/internal/config"
	"github.com/spf13/cobra"
)

type configUpdateContextOptions struct {
	ctxName string

	cfgFilename string

	baseURL   string
	clusterID string
}

func newCmdConfigUpdateContext() *cobra.Command {
	o := configUpdateContextOptions{}

	cmd := &cobra.Command{
		Use:   "update-context NAME",
		Short: "Update a context",
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

	cmd.Flags().StringVar(&o.baseURL, "base-url", "", "Kafka Rest Proxy Base URL")
	cmd.Flags().StringVar(&o.clusterID, "cluster-id", "", "Kafka Rest Proxy Cluster ID")

	return cmd
}

func (o *configUpdateContextOptions) setup(cmd *cobra.Command, args []string) error {
	o.ctxName = args[0]

	cfgFilename, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return err
	}
	baseURL, err := cmd.Flags().GetString("base-url")
	if err != nil {
		return err
	}
	clusterID, err := cmd.Flags().GetString("cluster-id")
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
	o.baseURL = baseURL
	o.clusterID = clusterID

	return nil
}

func (o *configUpdateContextOptions) validate() error {
	if o.baseURL == "" && o.clusterID == "" {
		return fmt.Errorf("nothing to do")
	}
	return nil
}

func (o *configUpdateContextOptions) run() error {
	cfg, err := config.LoadFromFile(o.cfgFilename)
	if err != nil {
		return err
	}

	ctx, err := cfg.GetContext(o.ctxName)
	if err != nil {
		return err
	}
	if o.baseURL != "" {
		ctx.BaseURL = o.baseURL
	}
	if o.clusterID != "" {
		ctx.ClusterID = o.clusterID
	}

	if err := cfg.UpdateContext(ctx); err != nil {
		return err
	}

	if err := config.WriteToFile(cfg, o.cfgFilename); err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Context '%s' updated successfully\n", o.ctxName)
	return nil
}
