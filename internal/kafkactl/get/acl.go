package get

import (
	"context"
	"os"

	v3 "github.com/devodev/kafkactl/internal/api/v3"
	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/spf13/cobra"
)

type getAclOptions struct {
	cli *cli.CLI

	queryParams *v3.AclParams
}

func newCmdGetAcl(c *cli.CLI) *cobra.Command {
	o := getAclOptions{cli: c, queryParams: &v3.AclParams{}}

	cmd := &cobra.Command{
		Use:     "acl",
		Aliases: []string{"acls"},
		Short:   "ACL API",
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

	cmd.Flags().StringVar(&o.queryParams.ResourceType, "resource-type", "", "Filter by Resource Type")
	cmd.Flags().StringVar(&o.queryParams.ResourceName, "resource-name", "", "Filter by Resource Name")
	cmd.Flags().StringVar(&o.queryParams.PatternType, "pattern-type", "", "Filter by Pattern Type")
	cmd.Flags().StringVar(&o.queryParams.Principal, "principal", "", "Filter by Principal")
	cmd.Flags().StringVar(&o.queryParams.Host, "host", "", "Filter by Host")
	cmd.Flags().StringVar(&o.queryParams.Operation, "operation", "", "Filter by Operation")
	cmd.Flags().StringVar(&o.queryParams.Permission, "permission", "", "Filter by Permission")

	return cmd
}

func (o *getAclOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	return nil
}

func (o *getAclOptions) validate() error {
	return o.cli.Validate()
}

func (o *getAclOptions) run() error {
	resp, err := o.cli.Client.Acl.ListWide(context.Background(), o.cli.Context.ClusterID, o.queryParams)
	if err != nil {
		return util.MakeCLIError("get", "acl", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}
