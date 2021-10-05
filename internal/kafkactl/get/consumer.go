package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getConsumerOptions struct {
	cli *cli.CLI

	groupID    string
	consumerID string
}

func newCmdGetConsumer() *cobra.Command {
	o := getConsumerOptions{}

	cmd := &cobra.Command{
		Use:     "consumer GROUP_ID [CONSUMER_ID]",
		Aliases: []string{"consumers", "co"},
		Short:   "Consumer API",
		Args:    cobra.RangeArgs(1, 2),
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

func (o *getConsumerOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args)
	if err != nil {
		return err
	}
	o.cli = cli

	o.groupID = args[0]
	if len(args) == 2 {
		o.consumerID = args[1]
	}

	return nil
}

func (o *getConsumerOptions) validate() error {
	return o.cli.Validate()
}

func (o *getConsumerOptions) run() error {
	var resp interface{}
	var err error
	if o.consumerID == "" {
		resp, err = o.cli.Client.Consumer.ListWide(context.Background(), o.cli.Context.ClusterID, o.groupID)
	} else {
		resp, err = o.cli.Client.Consumer.GetWide(context.Background(), o.cli.Context.ClusterID, o.groupID, o.consumerID)
	}
	if err != nil {
		return util.MakeCLIError("get", "consumer", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
