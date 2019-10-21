package command

import "github.com/spf13/cobra"

func NewUpgradeCommand() *cobra.Command {
	upgradeCmd := &cobra.Command{
		Use:   "upgrade <name>",
		Short: "upgrade tidb version",
	}

	return upgradeCmd
}
