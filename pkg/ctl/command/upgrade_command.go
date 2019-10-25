package command

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/tidbops/tim/pkg/diff"
	"github.com/tidbops/tim/pkg/models"
)

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

	name := args[0]
	cli, err := genClient(cmd)
	if err != nil {
		cmd.Printf("init client failed, %v\n", err)
	}

	tc, err := cli.GetTiDBClusterByName(name)
	if err != nil {
		cmd.Printf("%s tidb cluster not exist\n", name)
		return
	}

	tmpID := time.Now().Unix()
	tmpPath := fmt.Sprintf("/tmp/tim/%s/%d", tc.Name, tmpID)

	// just prepare tikv config fot demo
	// TODO: support prepare pd / tidb config
	oldTiKVConfig, targetTiKVConfig, err := prepareConfigFile(tc, upgradeCmdFlags.TargetVersion, tmpPath)
	if err != nil {
		cmd.Println("prepare config file failed, %v", err)
		return
	}

	diffStr, err := diff.DiffYaml(oldTiKVConfig, targetTiKVConfig, true)
	if err != nil {
		cmd.Printf("compare %s %s failed, %v\n", oldTiKVConfig, targetTiKVConfig, err)
		return
	}
	cmd.Println(diffStr)
}

const (
	tikvRawConfigURL = "https://raw.githubusercontent.com/pingcap/tidb-ansible/%s/conf/tikv.yml"
)

func prepareConfigFile(tc *models.TiDBCluster, targetVersion string, path string) (string, string, error) {
	// cloneOldCmd := exec.Command("sh", "-c",
	// fmt.Sprintf("git clone -b %s %s", tc.Version, TiDBAnsibleURL))

	// tmpPath := fmt.Sprintf("/tmp/tim/%s/%d", tc.Name, id)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return "", "", err
	}

	oldRawTiKVConfigURL := fmt.Sprintf(tikvRawConfigURL, tc.Version)
	oldTiKVConfigPath := filepath.Join(path, fmt.Sprintf("%s-tikv.yml", tc.Version))
	if err := DownloadFile(oldRawTiKVConfigURL, oldTiKVConfigPath); err != nil {
		return "", "", err
	}

	targetRawTiKVConfigURL := fmt.Sprintf(tikvRawConfigURL, targetVersion)
	targetTiKVConfigPath := filepath.Join(path, fmt.Sprintf("%s-tikv.yml", targetVersion))
	if err := DownloadFile(targetRawTiKVConfigURL, targetTiKVConfigPath); err != nil {
		return "", "", err
	}

	return oldTiKVConfigPath, targetTiKVConfigPath, nil
}
