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

type getPartitionReplicaOptions struct {
	cli *cli.CLI

	topic       string
	partitionID *int
	brokerID    *int
}

func newCmdGetPartitionReplica(c *cli.CLI) *cobra.Command {
	o := getPartitionReplicaOptions{cli: c}

	cmd := &cobra.Command{
		Use:     "partition-replica TOPIC_NAME [PARTITION_ID] [BROKER_ID]",
		Aliases: []string{"partition-replicas", "partitionreplica", "partitionreplicas", "replica", "replicas", "pr"},
		Short:   "Partition Replica API",
		Args:    cobra.RangeArgs(1, 3),
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

func (o *getPartitionReplicaOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	o.topic = args[0]
	if len(args) > 1 {
		var partitionID int
		partitionID, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid partition id: %s", err.Error())
		}
		o.partitionID = &partitionID
	}
	if len(args) == 3 {
		var brokerID int
		brokerID, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("invalid broker id: %s", err.Error())
		}
		o.brokerID = &brokerID
	}

	return nil
}

func (o *getPartitionReplicaOptions) validate() error {
	return o.cli.Validate()
}

func (o *getPartitionReplicaOptions) run() error {
	var resp interface{}
	var err error
	if o.partitionID == nil {
		resp, err = o.cli.Client.PartitionReplica.ListAllWide(context.Background(), o.cli.Context.ClusterID, o.topic)
	} else if o.brokerID == nil {
		resp, err = o.cli.Client.PartitionReplica.ListWide(context.Background(), o.cli.Context.ClusterID, o.topic, *o.partitionID)
	} else {
		resp, err = o.cli.Client.PartitionReplica.GetWide(context.Background(), o.cli.Context.ClusterID, o.topic, *o.partitionID, *o.brokerID)
	}
	if err != nil {
		return util.MakeCLIError("get", "partition-replica", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
