package get

import (
	"context"
	"os"

	"github.com/devodev/kafkactl/internal/cli"
	"github.com/devodev/kafkactl/internal/kafkactl/util"
	"github.com/devodev/kafkactl/internal/presentation"
	"github.com/spf13/cobra"
)

type getAclOptions struct {
	cli *cli.CLI

	resourceType string
	resourceName string
	patternType  string
	principal    string
	host         string
	operation    string
	permission   string
}

func newCmdGetAcl(c *cli.CLI) *cobra.Command {
	o := getAclOptions{cli: c}

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

	cmd.Flags().StringVar(&o.resourceType, "resource-type", "", "Filter by Resource Type")
	cmd.Flags().StringVar(&o.resourceName, "resource-name", "", "Filter by Resource Name")
	cmd.Flags().StringVar(&o.patternType, "pattern-type", "", "Filter by Pattern Type")
	cmd.Flags().StringVar(&o.principal, "principal", "", "Filter by Principal")
	cmd.Flags().StringVar(&o.host, "host", "", "Filter by Host")
	cmd.Flags().StringVar(&o.operation, "operation", "", "Filter by Operation")
	cmd.Flags().StringVar(&o.permission, "permission", "", "Filter by Permission")

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
	resp, err := o.cli.Client.Acl.ListWide(context.Background(), o.cli.Context.ClusterID, o.queryParams())
	if err != nil {
		return util.MakeCLIError("get", "acl", err)
	}
	if err := o.cli.Serializer.Serialize(resp, os.Stdout); err != nil {
		return util.MakeSerializationError(err)
	}
	return nil
}

func (o *getAclOptions) queryParams() *presentation.AclListQueryParams {
	return &presentation.AclListQueryParams{
		ResourceType: o.resourceType,
		ResourceName: o.resourceName,
		PatternType:  o.patternType,
		Principal:    o.principal,
		Host:         o.host,
		Operation:    o.operation,
		Permission:   o.permission,
	}
}
