package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getClusterConfigOptions struct {
	cli *cli.CLI

	configName string
}

func newCmdGetClusterConfig() *cobra.Command {
	o := getClusterConfigOptions{}

	cmd := &cobra.Command{
		Use:     "cluster-config [CONFIG_NAME]",
		Aliases: []string{"cluster-configs", "clusterconfig", "clusterconfigs", "cc"},
		Short:   "Cluster Config API",
		Args:    cobra.RangeArgs(0, 1),
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

func (o *getClusterConfigOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args)
	if err != nil {
		return err
	}
	o.cli = cli

	if len(args) == 1 {
		o.configName = args[0]
	}

	return nil
}

func (o *getClusterConfigOptions) validate() error {
	return o.cli.Validate()
}

func (o *getClusterConfigOptions) run() error {
	var resp interface{}
	var err error
	if o.configName == "" {
		resp, err = o.cli.Client.ClusterConfig.ListWide(context.Background(), o.cli.Context.ClusterID)
	} else {
		resp, err = o.cli.Client.ClusterConfig.GetWide(context.Background(), o.cli.Context.ClusterID, o.configName)
	}
	if err != nil {
		return util.MakeCLIError("get", "cluster-config", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
