package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	kafkaclient "github.com/devodev/kafkactl/internal/client"
	"github.com/devodev/kafkactl/internal/config"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/devodev/kafkactl/internal/serializers"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type cliOption func(c *CLI) error

func WithLimitedOutputs(outputs ...string) cliOption {
	return func(c *CLI) error {
		if len(outputs) == 0 {
			return fmt.Errorf("supported outputs is empty")
		}
		for _, o := range outputs {
			if !serializers.IsSupported(o) {
				return fmt.Errorf("output not supported: %s", o)
			}
		}
		c.supportedOutputs = outputs
		return nil
	}
}

type cliInitOption cliOption

func WithIgnoreClusterIDUnset() cliInitOption {
	return func(c *CLI) error {
		c.ignoreClusterIDUnset = true
		return nil
	}
}

type CLI struct {
	flagset *pflag.FlagSet

	ignoreClusterIDUnset bool
	supportedOutputs     []string

	Context    *config.Context
	Client     *kafkaclient.KafkaRest
	Serializer serializers.Serializer
}

func New(opts ...cliOption) (*CLI, error) {
	var supportedOutputs []string
	for _, s := range serializers.Types {
		supportedOutputs = append(supportedOutputs, s.String())
	}

	c := &CLI{
		supportedOutputs: supportedOutputs,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	flagset := pflag.NewFlagSet("", pflag.ContinueOnError)
	flagset.StringP("output", "o", c.supportedOutputs[0], fmt.Sprintf("Output format (%s)", strings.Join(c.supportedOutputs, ",")))
	flagset.StringP("config-file", "f", "", "Configuration file path")
	flagset.StringArrayP("header", "H", []string{}, "Additional HTTP header(s)")

	c.flagset = flagset

	return c, nil
}

func (c *CLI) Flags() *pflag.FlagSet {
	return c.flagset
}

func (c *CLI) Init(cmd *cobra.Command, args []string, opts ...cliInitOption) error {
	// parse flags
	output, err := c.flagset.GetString("output")
	if err != nil {
		return err
	}
	cfgFilename, err := c.flagset.GetString("config-file")
	if err != nil {
		return err
	}
	headers, err := c.flagset.GetStringArray("header")
	if err != nil {
		return err
	}

	// resolve default config filepath
	if cfgFilename == "" {
		var home string
		home, err = os.UserHomeDir()
		if err != nil {
			return err
		}
		cfgFilename = filepath.Join(home, ".kafkactl.yaml")
	}

	// load config
	var cfg *config.Config
	cfg, err = config.LoadFromFile(cfgFilename)
	if err != nil {
		return err
	}

	// load context
	ctx, err := cfg.GetCurrentContext()
	if err != nil {
		return err
	}

	if ctx.BaseURL == "" {
		return fmt.Errorf("baseURL not set in current context")
	}

	// parse headers
	headerMap, err := util.KeyValueParse(":", headers)
	if err != nil {
		return fmt.Errorf("could not parse headers: %s", err)
	}

	// create Kafka Rest client
	client, err := kafkaclient.New(ctx.BaseURL, kafkaclient.WithHeaders(headerMap))
	if err != nil {
		return err
	}

	// create response serializer
	ser, err := serializers.NewSerializer(output)
	if err != nil {
		return err
	}

	// create cli instance
	c.Context = ctx
	c.Client = client
	c.Serializer = ser

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}

	return nil
}

func (c *CLI) Validate() error {
	if !c.ignoreClusterIDUnset && c.Context.ClusterID == "" {
		return fmt.Errorf("clusterID not set in current context")
	}
	return nil
}
