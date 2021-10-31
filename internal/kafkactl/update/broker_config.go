package update

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type updateBrokerConfigOptions struct {
	cli *cli.CLI

	brokerIDList []string

	brokerIDs map[int]struct{}
	configs   []string

	request *v3.ConfigBatchAlterRequest
}

func newCmdupdateBrokerConfig(c *cli.CLI) *cobra.Command {
	o := updateBrokerConfigOptions{cli: c}

	cmd := &cobra.Command{
		Use:     "broker-config KEYVALUE_PAIR [KEYVALUE_PAIR..]",
		Aliases: []string{"broker-configs", "brokerconfig", "brokerconfigs", "bc"},
		Short:   "Broker Config API",
		Args:    cobra.MinimumNArgs(1),
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

	cmd.Flags().StringSliceVarP(&o.brokerIDList, "broker-id", "b", []string{}, "Comma-delimited list of broker IDs to update")

	return cmd
}

func (o *updateBrokerConfigOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	o.configs = args

	o.brokerIDs = make(map[int]struct{}, len(o.brokerIDList))
	for _, idStr := range o.brokerIDList {
		brokerID, err := strconv.Atoi(idStr)
		if err != nil {
			return fmt.Errorf("invalid broker id: %s", err.Error())
		}
		o.brokerIDs[brokerID] = struct{}{}
	}

	req, err := o.makeRequest()
	if err != nil {
		return err
	}
	o.request = req

	return nil
}

func (o *updateBrokerConfigOptions) validate() error {
	if err := o.cli.Validate(); err != nil {
		return err
	}

	if len(o.request.Data) == 0 {
		return fmt.Errorf("nothing to do")
	}
	return nil
}

func (o *updateBrokerConfigOptions) run() error {
	// get all brokers
	resp, err := o.cli.Client.Broker.ListWide(context.Background(), o.cli.Context.ClusterID)
	if err != nil {
		return err
	}

	currBrokerIDs := resp.BrokerIDMap()

	if len(o.brokerIDs) > 0 {
		notFound := make([]string, 0)
		for id := range o.brokerIDs {
			if _, ok := currBrokerIDs[id]; !ok {
				notFound = append(notFound, strconv.Itoa(id))
			}
		}
		if len(notFound) > 0 {
			return fmt.Errorf("broker IDs not found: %s", strings.Join(notFound, ","))
		}
		// if all IDs found, use them instead of complete list
		currBrokerIDs = o.brokerIDs
	}

	statuses := make([]string, 0, len(currBrokerIDs))
	errors := make([]string, 0, len(currBrokerIDs))
	for id := range currBrokerIDs {
		status, err := o.cli.Client.BrokerConfig.BatchAlter(context.Background(), o.cli.Context.ClusterID, id, o.request)
		if err != nil {
			errors = append(errors, fmt.Sprintf("[id:%d] %s", id, err.Error()))
			continue
		}
		statuses = append(statuses, status)
	}
	if len(errors) > 0 {
		return util.MakeCLIError("update", "broker-config", fmt.Errorf(strings.Join(errors, ": ")))
	}
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(statuses, "\n"))
	return nil
}

func (o *updateBrokerConfigOptions) makeRequest() (*v3.ConfigBatchAlterRequest, error) {
	var req v3.ConfigBatchAlterRequest

	brokerConfigs, err := util.KeyValueDeleteParse("=", "-", o.configs)
	if err != nil {
		return nil, fmt.Errorf("could not parse config key-value pair: %s", err.Error())
	}
	if len(brokerConfigs) > 0 {
		req.Data = make([]v3.ConfigBatchAlterData, 0, len(brokerConfigs))
	}
	for key, value := range brokerConfigs {
		var config v3.ConfigBatchAlterData
		if value == "" {
			config = v3.ConfigBatchAlterData{Name: key, Operation: v3.ConfigDeleteOperation}
		} else {
			config = v3.ConfigBatchAlterData{Name: key, Value: value}
		}
		req.Data = append(req.Data, config)
	}

	return &req, nil
}
