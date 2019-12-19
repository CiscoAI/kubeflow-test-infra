package create

import (
	"github.com/spf13/cobra"
)

// Flag holds the CLI flags in a consolidated format for easier re-use
type Flag struct {
	Verbose bool
}

// NewCommand returns a new cobra.Command for create
func NewCommand() *cobra.Command {
	flags := &Flag{Verbose: false}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "create",
		Short: "Creates a ProwJob on local or in-cluster",
		Long:  "Creates a ProwJob on local or in-cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			// set flags here
			return create(flags, cmd, args)
		},
	}
	return cmd
}

func create(flags *Flag, cmd *cobra.Command, args []string) error {
	return nil
}
