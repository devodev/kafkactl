package describe

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/devodev/kafkactl/internal/presentation"
	"github.com/devodev/kafkactl/internal/serializers"
	"github.com/spf13/cobra"
)

type describeTopicOptions struct {
	cli *cli.CLI

	topic string
}

func newCmdDescribeTopic(c *cli.CLI) *cobra.Command {
	o := describeTopicOptions{cli: c}

	cmd := &cobra.Command{
		Use:     "topic TOPIC_NAME",
		Aliases: []string{"topics", "to"},
		Short:   "Topic API",
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

func (o *describeTopicOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	o.topic = args[0]

	return nil
}

func (o *describeTopicOptions) validate() error {
	return o.cli.Validate()
}

func (o *describeTopicOptions) run() error {
	topic, err := o.cli.Client.Topic.GetWide(context.Background(), o.cli.Context.ClusterID, o.topic)
	if err != nil {
		return util.MakeCLIError("describe", "topic", err)
	}

	groups, err := o.cli.Client.ConsumerGroup.ListWide(context.Background(), o.cli.Context.ClusterID)
	if err != nil {
		return util.MakeCLIError("describe", "topic", err)
	}

	// TODO: filter groups by topic (use consumer assignments)

	// TODO: Add offset/lag to group output

	container := describeTopicContainer{
		Topic:  topic,
		Groups: groups,
	}

	tc := &serializers.TemplateContainer{
		Templates: []string{
			presentation.TopicDescribeTemplate,
		},
		Data: container,
	}

	if err := o.cli.Serializer.Serialize(tc, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}

type describeTopicContainer struct {
	*presentation.Topic
	Groups presentation.ConsumerGroupList
}
