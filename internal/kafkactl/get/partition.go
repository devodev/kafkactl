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

type getPartitionOptions struct {
	cli *cli.CLI

	topic       string
	partitionID *int
}

func newCmdGetPartition() *cobra.Command {
	o := getPartitionOptions{}

	cmd := &cobra.Command{
		Use:     "partition TOPIC_NAME [PARTITION_ID]",
		Aliases: []string{"partitions", "pa"},
		Short:   "Partition API",
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

func (o *getPartitionOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args)
	if err != nil {
		return err
	}
	o.cli = cli

	o.topic = args[0]

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

func (o *getPartitionOptions) validate() error {
	return o.cli.Validate()
}

func (o *getPartitionOptions) run() error {
	var resp interface{}
	var err error
	if o.partitionID == nil {
		resp, err = o.cli.Client.Partition.ListWide(context.Background(), o.cli.Context.ClusterID, o.topic)
	} else {
		resp, err = o.cli.Client.Partition.GetWide(context.Background(), o.cli.Context.ClusterID, o.topic, *o.partitionID)
	}
	if err != nil {
		return util.MakeCLIError("get", "partition", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
