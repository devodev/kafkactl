package config

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/devodev/kafkactl/internal/config"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/devodev/kafkactl/internal/presentation"
	"github.com/devodev/kafkactl/internal/serializers"
	"github.com/spf13/cobra"
)

type configGetContextOptions struct {
	cfgFilename string

	serializer serializers.Serializer
}

func newCmdConfigGetContext() *cobra.Command {
	o := configGetContextOptions{}

	cmd := &cobra.Command{
		Use:     "get-context",
		Aliases: []string{"get-contexts"},
		Short:   "List contexts",
		Args:    cobra.NoArgs,
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

	cmd.Flags().StringP("output", "o", "table", "How to format the output (table, json)")

	return cmd
}

func (o *configGetContextOptions) setup(cmd *cobra.Command, args []string) error {
	cfgFilename, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	// if flag not set, load config using default filepath
	if cfgFilename == "" {
		var home string
		home, err = os.UserHomeDir()
		if err != nil {
			return err
		}
		cfgFilename = filepath.Join(home, ".kafkactl.yaml")
	}

	// create response serializer
	serializer, err := serializers.NewSerializer(output)
	if err != nil {
		return err
	}

	o.serializer = serializer

	o.cfgFilename = cfgFilename

	return nil
}

func (o *configGetContextOptions) validate() error { return nil }

func (o *configGetContextOptions) run() error {
	cfg, err := config.LoadFromFile(o.cfgFilename)
	if err != nil {
		return err
	}

	currentCtx := cfg.GetCurrentContextName()
	contexts := cfg.GetContexts()

	configs := make(presentation.ConfigList, 0, len(contexts))
	for _, ctx := range contexts {
		isCurrent := ctx.Name == currentCtx
		configs = append(configs, *presentation.MapConfig(ctx, isCurrent))
	}
	sort.Sort(presentation.ConfigAlphabeticSort(configs))

	if err := o.serializer.Serialize(configs, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}

	return nil
}
