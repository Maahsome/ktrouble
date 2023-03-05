package set

import (
	"ktrouble/common"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type gitUserParam struct {
	Name     string
	TokenVar string
	Token    string
}

var p gitUserParam

// gitconfigCmd represents the gituser command
var gitconfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Set git configuration options",
	Long: `EXAMPLE:
  If you store your git personal access token in an ENV variable, you can specify
  the variable name.

  > ktrouble set config --user christopher.maahs --tokenvar GLA_TOKEN

EXAMPLE:
  If you don't store your personal access token in an ENV variable, it can be
  stored directly in the config.yaml file.  Don't forgot to add a 'space' in
  front of running this next command so the token doesn't end up in your
  history file, if you have that option set in your shell

  > ktrouble set config --user christopher.maahs --token <your token>
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := saveConfig()
		if err != nil {
			logrus.WithError(err).Error("Failed to save the gitUser property")
		}
	},
}

func saveConfig() error {

	if len(p.Name) > 0 {
		viper.Set("gitUser", p.Name)
	}
	if len(p.TokenVar) > 0 {
		viper.Set("gitTokenVar", p.TokenVar)
	}
	if len(p.Token) > 0 {
		viper.Set("gitToken", p.Token)
	}
	verr := viper.WriteConfig()
	if verr != nil {
		common.Logger.WithError(verr).Info("Failed to write config")
		return verr
	}
	return nil
}

func init() {
	setCmd.AddCommand(gitconfigCmd)

	gitconfigCmd.Flags().StringVarP(&p.Name, "user", "u", "", "Set your git username")
	gitconfigCmd.Flags().StringVar(&p.TokenVar, "tokenvar", "", "Set the name of the ENV VAR that contains your git personal token")
	gitconfigCmd.Flags().StringVar(&p.Token, "token", "", "Set your git personal token")
}
