package update

import (
	"context"
	"fmt"
	"os"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type updateTopicConfigOptions struct {
	cli *cli.CLI

	topic   string
	configs []string

	request *v3.ConfigBatchAlterRequest
}

func newCmdUpdateTopicConfig(c *cli.CLI) *cobra.Command {
	o := updateTopicConfigOptions{cli: c}

	cmd := &cobra.Command{
		Use:     "topic-config TOPIC_NAME KEYVALUE_PAIR [KEYVALUE_PAIR..]",
		Aliases: []string{"topic-configs", "topicconfig", "topic-configs", "tc"},
		Short:   "Topic Config API",
		Args:    cobra.MinimumNArgs(2),
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

func (o *updateTopicConfigOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	o.topic = args[0]
	o.configs = args[1:]

	req, err := o.makeRequest()
	if err != nil {
		return err
	}
	o.request = req

	return nil
}

func (o *updateTopicConfigOptions) validate() error {
	if err := o.cli.Validate(); err != nil {
		return err
	}
	if len(o.request.Data) == 0 {
		return fmt.Errorf("nothing to do")
	}
	return nil
}

func (o *updateTopicConfigOptions) run() error {
	status, err := o.cli.Client.TopicConfig.BatchAlter(context.Background(), o.cli.Context.ClusterID, o.topic, o.request)
	if err != nil {
		return util.MakeCLIError("update", "topic-config", err)
	}
	fmt.Fprintln(os.Stdout, status)
	return nil
}

func (o *updateTopicConfigOptions) makeRequest() (*v3.ConfigBatchAlterRequest, error) {
	var req v3.ConfigBatchAlterRequest

	topicConfigs, err := util.KeyValueDeleteParse("=", "-", o.configs)
	if err != nil {
		return nil, fmt.Errorf("could not parse config key-value pair: %s", err.Error())
	}
	if len(topicConfigs) > 0 {
		req.Data = make([]v3.ConfigBatchAlterData, 0, len(topicConfigs))
	}
	for key, value := range topicConfigs {
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
