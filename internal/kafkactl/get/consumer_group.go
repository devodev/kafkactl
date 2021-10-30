package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getConsumerGroupOptions struct {
	cli *cli.CLI

	groupID string
}

func newCmdGetConsumerGroup(c *cli.CLI) *cobra.Command {
	o := getConsumerGroupOptions{cli: c}

	cmd := &cobra.Command{
		Use:     "consumer-group [GROUP_ID]",
		Aliases: []string{"consumer-groups", "consumergroup", "consumergroups", "cg"},
		Short:   "Consumer Group API",
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

func (o *getConsumerGroupOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	if len(args) == 1 {
		o.groupID = args[0]
	}

	return nil
}

func (o *getConsumerGroupOptions) validate() error {
	return o.cli.Validate()
}

func (o *getConsumerGroupOptions) run() error {
	var resp interface{}
	var err error
	if o.groupID == "" {
		resp, err = o.cli.Client.ConsumerGroup.ListWide(context.Background(), o.cli.Context.ClusterID)
	} else {
		resp, err = o.cli.Client.ConsumerGroup.GetWide(context.Background(), o.cli.Context.ClusterID, o.groupID)
	}
	if err != nil {
		return util.MakeCLIError("get", "consumer-group", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
