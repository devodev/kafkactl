package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getTopicOptions struct {
	cli *cli.CLI

	topic string
}

func newCmdGetTopic() *cobra.Command {
	o := getTopicOptions{}

	cmd := &cobra.Command{
		Use:     "topic [TOPIC_NAME]",
		Aliases: []string{"topics", "to"},
		Short:   "Topic API",
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

func (o *getTopicOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args)
	if err != nil {
		return err
	}
	o.cli = cli

	if len(args) == 1 {
		o.topic = args[0]
	}
	return nil
}

func (o *getTopicOptions) validate() error {
	return o.cli.Validate()
}

func (o *getTopicOptions) run() error {
	var resp interface{}
	var err error
	if o.topic == "" {
		resp, err = o.cli.Client.Topic.ListWide(context.Background(), o.cli.Context.ClusterID)
	} else {
		resp, err = o.cli.Client.Topic.GetWide(context.Background(), o.cli.Context.ClusterID, o.topic)
	}
	if err != nil {
		return util.MakeCLIError("get", "topic", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
