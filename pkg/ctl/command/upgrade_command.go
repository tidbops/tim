package command

import "github.com/spf13/cobra"

type UpgradeCommandFlags struct {
	TargetVersion string
	RuleFile      string
}

var (
	upgradeCmdFlags = &UpgradeCommandFlags{}
)

func NewUpgradeCommand() *cobra.Command {
	upgradeCmd := &cobra.Command{
		Use:   "upgrade <name>",
		Short: "upgrade tidb version, just generate the new version tidb-ansible files",
		Run:   upgradeCommandFunc,
	}

	upgradeCmd.Flags().StringVar(&upgradeCmdFlags.TargetVersion, "target-version", "", "the version that ready to upgrade to")
	upgradeCmd.Flags().StringVar(&upgradeCmdFlags.RuleFile, "rule-file", "",
		"rule files for different version of configuration conversion")

	return upgradeCmd
}

func upgradeCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) < 0 {
		cmd.Println("name is required")
		cmd.Println(cmd.UsageString())
		return
	}

	if upgradeCmdFlags.TargetVersion == "" {
		cmd.Println("target-version flag is required")
		cmd.Println(cmd.UsageString())
		return
	}
}
