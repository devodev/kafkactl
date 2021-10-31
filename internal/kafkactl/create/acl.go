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

type createAclOptions struct {
	cli *cli.CLI

	params *v3.AclParams
}

// TODO: improve UX, something like:
//
//       kafkactl create acl cluster PRINCIPAL
//       kafkactl create acl group   PRINCIPAL GROUP_NAME [--pattern PATTERN (prefixed,literal,match)]
//       kafkactl create acl topic   PRINCIPAL TOPIC_NAME [--pattern PATTERN (prefixed,literal,match)]
//
func newCmdCreateAcl(c *cli.CLI) *cobra.Command {
	o := createAclOptions{cli: c, params: &v3.AclParams{}}

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

	cmd.Flags().StringVar(&o.params.ResourceType, "resource-type", "", "Filter by Resource Type")
	cmd.Flags().StringVar(&o.params.ResourceName, "resource-name", "", "Filter by Resource Name")
	cmd.Flags().StringVar(&o.params.PatternType, "pattern-type", "", "Filter by Pattern Type")
	cmd.Flags().StringVar(&o.params.Principal, "principal", "", "Filter by Principal")
	cmd.Flags().StringVar(&o.params.Host, "host", "", "Filter by Host")
	cmd.Flags().StringVar(&o.params.Operation, "operation", "", "Filter by Operation")
	cmd.Flags().StringVar(&o.params.Permission, "permission", "", "Filter by Permission")

	return cmd
}

func (o *createAclOptions) setup(cmd *cobra.Command, args []string) error {
	if err := o.cli.Init(cmd, args); err != nil {
		return err
	}

	return nil
}

func (o *createAclOptions) validate() error {
	return o.cli.Validate()
}

func (o *createAclOptions) run() error {
	req, err := o.params.Request()
	if err != nil {
		return err
	}
	status, err := o.cli.Client.Acl.Create(context.Background(), o.cli.Context.ClusterID, req)
	if err != nil {
		return util.MakeCLIError("create", "acl", err)
	}
	fmt.Fprintln(os.Stdout, status)
	return nil
}
