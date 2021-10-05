package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getTopicConfigOptions struct {
	cli *cli.CLI

	topic      string
	configName string
}

func newCmdGetTopicConfig() *cobra.Command {
	o := getTopicConfigOptions{}

	cmd := &cobra.Command{
		Use:     "topic-config TOPIC_NAME [CONFIG_NAME]",
		Aliases: []string{"topic-configs", "topicconfig", "topic-configs", "tc"},
		Short:   "Topic Config API",
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

func (o *getTopicConfigOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args)
	if err != nil {
		return err
	}
	o.cli = cli

	o.topic = args[0]

	if len(args) == 2 {
		o.configName = args[1]
	}

	return nil
}

func (o *getTopicConfigOptions) validate() error {
	return o.cli.Validate()
}

func (o *getTopicConfigOptions) run() error {
	var resp interface{}
	var err error
	if o.configName == "" {
		resp, err = o.cli.Client.TopicConfig.ListWide(context.Background(), o.cli.Context.ClusterID, o.topic)
	} else {
		resp, err = o.cli.Client.TopicConfig.GetWide(context.Background(), o.cli.Context.ClusterID, o.topic, o.configName)
	}
	if err != nil {
		return util.MakeCLIError("get", "topic-config", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
