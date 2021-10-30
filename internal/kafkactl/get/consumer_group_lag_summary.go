package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getConsumerGroupLagSummaryOptions struct {
	cli *cli.CLI

	groupID string
}

func newCmdGetConsumerGroupLagSummary(c *cli.CLI) *cobra.Command {
	o := getConsumerGroupLagSummaryOptions{cli: c}

	cmd := &cobra.Command{
		Use:     "lag-summary [GROUP_ID]",
		Aliases: []string{"lagsummary"},
		Short:   "Consumer Group API (lag-summary)",
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

func (o *getConsumerGroupLagSummaryOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	if len(args) == 1 {
		o.groupID = args[0]
	}

	return nil
}

func (o *getConsumerGroupLagSummaryOptions) validate() error {
	return o.cli.Validate()
}

func (o *getConsumerGroupLagSummaryOptions) run() error {
	var resp interface{}
	var err error
	if o.groupID == "" {
		resp, err = o.cli.Client.ConsumerGroup.LagSummaryAllWide(context.Background(), o.cli.Context.ClusterID)
	} else {
		resp, err = o.cli.Client.ConsumerGroup.LagSummaryWide(context.Background(), o.cli.Context.ClusterID, o.groupID)
	}
	if err != nil {
		return util.MakeCLIError("get", "lag-summary", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
