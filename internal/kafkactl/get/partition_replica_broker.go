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

type getPartitionReplicaBrokerOptions struct {
	cli *cli.CLI

	brokerID *int
}

func newCmdGetPartitionReplicaBroker() *cobra.Command {
	o := getPartitionReplicaBrokerOptions{}

	cmd := &cobra.Command{
		Use:     "partition-replica-broker [BROKER_ID]",
		Aliases: []string{"partition-replicas-broker"},
		Short:   "Partition Replica API (Broker)",
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

func (o *getPartitionReplicaBrokerOptions) setup(cmd *cobra.Command, args []string) error {
	cli, err := cli.New(cmd, args)
	if err != nil {
		return err
	}
	o.cli = cli

	if len(args) == 1 {
		var brokerID int
		brokerID, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid broker id: %s", err.Error())
		}
		o.brokerID = &brokerID
	}

	return nil
}

func (o *getPartitionReplicaBrokerOptions) validate() error {
	return o.cli.Validate()
}

func (o *getPartitionReplicaBrokerOptions) run() error {
	var resp interface{}
	var err error
	if o.brokerID == nil {
		resp, err = o.cli.Client.PartitionReplica.ListBrokerAllWide(context.Background(), o.cli.Context.ClusterID)
	} else {
		resp, err = o.cli.Client.PartitionReplica.ListBrokerWide(context.Background(), o.cli.Context.ClusterID, *o.brokerID)
	}
	if err != nil {
		return util.MakeCLIError("get", "partition-replica-broker", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
