package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getClusterOptions struct {
	cli *cli.CLI

	clusterID string
}

func newCmdGetCluster() *cobra.Command {
	o := getClusterOptions{}

	cmd := &cobra.Command{
		Use:     "cluster [CLUSTER_ID]",
		Aliases: []string{"clusters", "cl"},
		Short:   "Cluster API",
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

func (o *getClusterOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args, cli.WithIgnoreClusterIDUnset())
	if err != nil {
		return err
	}
	o.cli = cli

	if len(args) == 1 {
		o.clusterID = args[0]
	}

	return nil
}

func (o *getClusterOptions) validate() error {
	return o.cli.Validate()
}

func (o *getClusterOptions) run() error {
	var resp interface{}
	var err error
	if o.clusterID == "" {
		resp, err = o.cli.Client.Cluster.ListWide(context.Background())
	} else {
		resp, err = o.cli.Client.Cluster.GetWide(context.Background(), o.clusterID)
	}
	if err != nil {
		return util.MakeCLIError("get", "cluster", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
