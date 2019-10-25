package command

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/tidbops/tim/pkg/models"
	"github.com/tidbops/tim/pkg/utils"
	"github.com/tidbops/tim/pkg/yaml"
)

const (
	tikvRawConfigURL = "https://raw.githubusercontent.com/pingcap/tidb-ansible/%s/conf/tikv.yml"
)

const (
	InputNew     = "Input a new config file"
	UseOrigin    = "Use the origin config file"
	UseRuleFiles = "Use the configuration rules file to generate a new configuration file?"
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

	diffStr, err := yaml.Diff(oldTiKVConfig, targetTiKVConfig, true)
	if err != nil {
		cmd.Printf("compare %s %s failed, %v\n", oldTiKVConfig, targetTiKVConfig, err)
		return
	}

	if len(diffStr) > 0 {
		cmd.Println("Default tikv config has changed!")
		cmd.Println(diffStr)
	}

	prompt := promptui.Select{
		Label: "Select to init Config",
		Items: []string{
			InputNew,
			UseOrigin,
			UseRuleFiles,
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		cmd.Println(err)
		return
	}
	// cmd.Println(result)

	// var genConfig string

	switch result {
	case InputNew:
	case UseOrigin:
	case UseRuleFiles:
	default:
		cmd.Printf("%s is invalid\n", result)
		return
	}
}

func generateConfigByRuleFile(cmd cobra.Command) (string, error) {
	validate := func(input string) error {
		if exist := utils.FileExists(input); !exist {
			return fmt.Errorf("file %s not exist", input)
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Rule File",
		Validate: validate,
	}

	_, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return "", nil
}

func prepareConfigFile(tc *models.TiDBCluster, targetVersion string, path string) (string, string, error) {
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
