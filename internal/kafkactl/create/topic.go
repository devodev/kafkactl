package create

import (
	"context"
	"fmt"
	"os"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type createTopicOptions struct {
	cli *cli.CLI

	replicationFactor int
	partitionsCount   int
	configs           []string

	request *v3.TopicCreateRequest
}

func newCmdCreateTopic(c *cli.CLI) *cobra.Command {
	o := createTopicOptions{cli: c}

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

	cmd.Flags().IntVar(&o.replicationFactor, "replication-factor", 0, "Topic replication factor")
	cmd.Flags().IntVar(&o.partitionsCount, "partitions-count", 0, "Topic partitions count")
	cmd.Flags().StringArrayVar(&o.configs, "config", nil, "Topic config (key=value)")

	return cmd
}

func (o *createTopicOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	req, err := o.makeRequest(args[0])
	if err != nil {
		return err
	}
	o.request = req

	return nil
}

func (o *createTopicOptions) validate() error {
	return o.cli.Validate()
}

func (o *createTopicOptions) run() error {
	status, err := o.cli.Client.Topic.Create(context.Background(), o.cli.Context.ClusterID, o.request)
	if err != nil {
		return util.MakeCLIError("get", "topic", err)
	}
	fmt.Fprintln(os.Stdout, status)
	return nil
}

func (o *createTopicOptions) makeRequest(topic string) (*v3.TopicCreateRequest, error) {
	var req v3.TopicCreateRequest

	req.TopicName = topic
	if o.replicationFactor != 0 {
		req.ReplicationFactor = &o.replicationFactor
	}
	if o.partitionsCount != 0 {
		req.PartitionsCount = &o.partitionsCount
	}

	topicConfigs, err := util.KeyValueParse("=", o.configs)
	if err != nil {
		return nil, fmt.Errorf("could not parse config key-value pair: %s", err.Error())
	}
	if len(topicConfigs) > 0 {
		req.Configs = make([]v3.TopicConfig, 0, len(topicConfigs))
	}
	for key, value := range topicConfigs {
		config := v3.TopicConfig{Name: key, Value: value}
		req.Configs = append(req.Configs, config)
	}

	return &req, nil
}
