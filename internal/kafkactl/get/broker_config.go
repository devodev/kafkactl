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

type getBrokerConfigOptions struct {
	cli *cli.CLI

	brokerID   *int
	configName string
}

func newCmdGetBrokerConfig(c *cli.CLI) *cobra.Command {
	o := getBrokerConfigOptions{cli: c}

	cmd := &cobra.Command{
		Use:     "broker-config [BROKER_ID]",
		Aliases: []string{"broker-configs", "brokerconfig", "brokerconfigs", "bc"},
		Short:   "Broker Config API",
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

func (o *getBrokerConfigOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	if len(args) == 1 {
		var brokerID int
		brokerID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid broker id: %s", err.Error())
		}
		o.brokerID = &brokerID
	}

	if len(args) == 2 {
		o.configName = args[1]
	}

	return nil
}

func (o *getBrokerConfigOptions) validate() error {
	return o.cli.Validate()
}

func (o *getBrokerConfigOptions) run() error {
	var resp interface{}
	var err error
	if o.brokerID == nil {
		resp, err = o.cli.Client.BrokerConfig.ListAllWide(context.Background(), o.cli.Context.ClusterID)
	} else if o.configName == "" {
		resp, err = o.cli.Client.BrokerConfig.ListWide(context.Background(), o.cli.Context.ClusterID, *o.brokerID)
	} else {
		resp, err = o.cli.Client.BrokerConfig.GetWide(context.Background(), o.cli.Context.ClusterID, *o.brokerID, o.configName)
	}
	if err != nil {
		return util.MakeCLIError("get", "broker-config", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
