package set

import (
	"ktrouble/common"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type configUserParam struct {
	Name                string
	TokenVar            string
	Token               string
	GitURL              string
	PromptForSecrets    bool
	PromptForConfigMaps bool
	EnableBashLinks     bool
}

var p configUserParam

// gitconfigCmd represents the gituser command
var gitconfigCmd = &cobra.Command{
	Use:   "config",
	Short: setConfigHelp.Short(),
	Long:  setConfigHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		err := saveConfig()
		if err != nil {
			logrus.WithError(err).Error("Failed to save the gitUser property")
		}
	},
}

func saveConfig() error {

	itemsChanged := false
	if len(p.Name) > 0 {
		viper.Set("gitUser", p.Name)
		common.Logger.Info("The gitUser has been set")
		itemsChanged = true
	}
	if len(p.TokenVar) > 0 {
		viper.Set("gitTokenVar", p.TokenVar)
		common.Logger.Info("The gitTokenVar has been set")
		itemsChanged = true
	}
	if len(p.Token) > 0 {
		viper.Set("gitToken", p.Token)
		common.Logger.Info("The gitToken has been set")
		itemsChanged = true
	}
	if len(p.GitURL) > 0 {
		viper.Set("gitURL", p.GitURL)
		common.Logger.Info("The gitUrl has been set")
		itemsChanged = true
	}
	if p.PromptForSecrets {
		itemsChanged = true
		if c.PromptForSecrets {
			viper.Set("promptForSecrets", false)
			common.Logger.Info("The promptForSecrets default has been set to false")
		} else {
			viper.Set("promptForSecrets", true)
			common.Logger.Info("The promptForSecrets default has been set to true")
		}
	}
	if p.PromptForConfigMaps {
		itemsChanged = true
		if c.PromptForConfigMaps {
			viper.Set("promptForConfigMaps", false)
			common.Logger.Info("The promptForConfigMaps default has been set to false")
		} else {
			viper.Set("promptForConfigMaps", true)
			common.Logger.Info("The promptForConfigMaps default has been set to true")
		}
	}
	if p.EnableBashLinks {
		itemsChanged = true
		if c.EnableBashLinks {
			viper.Set("enableBashLinks", false)
			common.Logger.Info("The enableBashLinks has been set to false")
		} else {
			viper.Set("enableBashLinks", true)
			common.Logger.Info("The enableBashLinks has been set to true")
		}
	}
	if itemsChanged {
		verr := viper.WriteConfig()
		if verr != nil {
			common.Logger.WithError(verr).Info("Failed to write config")
			return verr
		}
	}
	return nil
}

func init() {
	setCmd.AddCommand(gitconfigCmd)

	gitconfigCmd.Flags().StringVarP(&p.Name, "user", "u", "", "Set your git username")
	gitconfigCmd.Flags().StringVar(&p.TokenVar, "tokenvar", "", "Set the name of the ENV VAR that contains your git personal token")
	gitconfigCmd.Flags().StringVar(&p.Token, "token", "", "Set your git personal token")
	gitconfigCmd.Flags().StringVar(&p.GitURL, "giturl", "", "Set the URL for the repository for upstream utils")
	gitconfigCmd.Flags().BoolVar(&p.PromptForSecrets, "secrets", false, "Toggle the Prompt for Secrets default")
	gitconfigCmd.Flags().BoolVar(&p.PromptForConfigMaps, "configs", false, "Toggle the Prompt for ConfigMaps default")
	gitconfigCmd.Flags().BoolVar(&p.EnableBashLinks, "bashlinks", false, "Toggle the use of Bash Links for iTerm2")
}
