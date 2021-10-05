package get

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getPartitionReassignmentOptions struct {
	cli *cli.CLI

	topic       string
	partitionID *int
}

func newCmdGetPartitionReassignment() *cobra.Command {
	o := getPartitionReassignmentOptions{}

	cmd := &cobra.Command{
		Use:     "partition-reassignment [TOPIC_NAME] [PARTITION_ID]",
		Aliases: []string{"partition-reassignments"},
		Short:   "Partition Reassignment API",
		Args:    cobra.RangeArgs(0, 2),
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

func (o *getPartitionReassignmentOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args)
	if err != nil {
		return err
	}
	o.cli = cli

	if len(args) > 0 {
		o.topic = args[0]
	}
	if len(args) == 2 {
		var partitionID int
		partitionID, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid partition id: %s", err.Error())
		}
		o.partitionID = &partitionID
	}

	return nil
}

func (o *getPartitionReassignmentOptions) validate() error {
	return o.cli.Validate()
}

func (o *getPartitionReassignmentOptions) run() error {
	var resp interface{}
	var err error
	if o.topic == "" {
		resp, err = o.cli.Client.PartitionReassignment.ListAllWide(context.Background(), o.cli.Context.ClusterID)
	} else if o.partitionID == nil {
		resp, err = o.cli.Client.PartitionReassignment.ListWide(context.Background(), o.cli.Context.ClusterID, o.topic)
	} else {
		resp, err = o.cli.Client.PartitionReassignment.GetWide(context.Background(), o.cli.Context.ClusterID, o.topic, *o.partitionID)
	}
	if err != nil {
		return util.MakeCLIError("get", "partition-reassignment", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
