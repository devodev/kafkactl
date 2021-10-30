package cli

import (
	"fmt"
	"os"
	"path/filepath"

	kafkaclient "github.com/devodev/kafkactl/internal/client"
	"github.com/devodev/kafkactl/internal/config"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/devodev/kafkactl/internal/serializers"
	"github.com/spf13/cobra"
)

type CLIOption func(c *CLI) error

type CLI struct {
	ignoreClusterIDUnset bool

	Context *config.Context

	Client     *kafkaclient.KafkaRest
	Serializer serializers.Serializer
}

func New(cmd *cobra.Command, args []string, opts ...CLIOption) (*CLI, error) {
	// parse flags
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return nil, err
	}
	cfgFilename, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return nil, err
	}
	headers, err := cmd.Flags().GetStringArray("header")
	if err != nil {
		return nil, err
	}

	// resolve default config filepath
	if cfgFilename == "" {
		var home string
		home, err = os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		cfgFilename = filepath.Join(home, ".kafkactl.yaml")
	}

	// load config
	var cfg *config.Config
	cfg, err = config.LoadFromFile(cfgFilename)
	if err != nil {
		return nil, err
	}

	// load context
	ctx, err := cfg.GetCurrentContext()
	if err != nil {
		return nil, err
	}

	if ctx.BaseURL == "" {
		return nil, fmt.Errorf("baseURL not set in current context")
	}

	// parse headers
	headerMap, err := util.KeyValueParse("=", headers)
	if err != nil {
		return nil, fmt.Errorf("could not parse headers: %s", err)
	}

	// create Kafka Rest client
	client, err := kafkaclient.New(ctx.BaseURL, kafkaclient.WithHeaders(headerMap))
	if err != nil {
		return nil, err
	}

	// create response serializer
	ser, err := serializers.NewSerializer(output)
	if err != nil {
		return nil, err
	}

	// create cli instance
	c := &CLI{
		Context:    ctx,
		Client:     client,
		Serializer: ser,
	}

	// apply options
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *CLI) Validate() error {
	if !c.ignoreClusterIDUnset && c.Context.ClusterID == "" {
		return fmt.Errorf("clusterID not set in current context")
	}
	return nil
}

func WithIgnoreClusterIDUnset() CLIOption {
	return func(c *CLI) error {
		c.ignoreClusterIDUnset = true
		return nil
	}
}
