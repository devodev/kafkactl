package delete

import (
	"context"
	"fmt"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type deleteTopicOptions struct {
	cli *cli.CLI

	topic string
}

func newCmdDeleteTopic(c *cli.CLI) *cobra.Command {
	o := deleteTopicOptions{cli: c}

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

func (o *deleteTopicOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	o.topic = args[0]

	return nil
}

func (o *deleteTopicOptions) validate() error {
	return o.cli.Validate()
}

func (o *deleteTopicOptions) run() error {
	status, err := o.cli.Client.Topic.Delete(context.Background(), o.cli.Context.ClusterID, o.topic)
	if err != nil {
		return util.MakeCLIError("delete", "topic", err)
	}
	fmt.Fprintln(os.Stdout, status)
	return nil
}
