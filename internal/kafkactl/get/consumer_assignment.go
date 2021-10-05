package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getConsumerAssignmentOptions struct {
	cli *cli.CLI

	groupID    string
	consumerID string
}

func newCmdGetConsumerAssignment() *cobra.Command {
	o := getConsumerAssignmentOptions{}

	cmd := &cobra.Command{
		Use:     "consumer-assignment GROUP_ID CONSUMER_ID",
		Aliases: []string{"consumer-assignments", "consumerassignment", "consumerassignments", "ca"},
		Short:   "Consumer Assignment API",
		Args:    cobra.ExactArgs(2),
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

func (o *getConsumerAssignmentOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args)
	if err != nil {
		return err
	}
	o.cli = cli

	o.groupID = args[0]
	o.consumerID = args[1]

	return nil
}

func (o *getConsumerAssignmentOptions) validate() error {
	return o.cli.Validate()
}

func (o *getConsumerAssignmentOptions) run() error {
	resp, err := o.cli.Client.ConsumerAssignment.ListWide(context.Background(), o.cli.Context.ClusterID, o.groupID, o.consumerID)
	if err != nil {
		return util.MakeCLIError("get", "consumer-assignment", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
