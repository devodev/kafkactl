package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getConsumerLagOptions struct {
	cli *cli.CLI

	groupID string
}

func newCmdGetConsumerLag(c *cli.CLI) *cobra.Command {
	o := getConsumerLagOptions{cli: c}

	cmd := &cobra.Command{
		Use:     "consumer-lag GROUP_ID",
		Aliases: []string{"consumer-lags", "consumerlag", "consumerlags", "lag", "lags"},
		Short:   "Consumer Lag API",
		Args:    cobra.ExactArgs(1),
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

func (o *getConsumerLagOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	o.groupID = args[0]

	return nil
}

func (o *getConsumerLagOptions) validate() error {
	return o.cli.Validate()
}

func (o *getConsumerLagOptions) run() error {
	resp, err := o.cli.Client.ConsumerLag.ListWide(context.Background(), o.cli.Context.ClusterID, o.groupID)
	if err != nil {
		return util.MakeCLIError("get", "consumer-lag", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
