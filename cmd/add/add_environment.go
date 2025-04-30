package add

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// environmentParam is used to store command line parameters
var environmentParam objects.Environment

// environmentCmd represents the utility command
var environmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: defaults.EnvironmentAliases,
	Short:   addEnvironmentHelp.Short(),
	Long:    addEnvironmentHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if checkAddEnvironmentParams() {
			u, err := addEnvironment()
			if err != nil {
				common.Logger.WithError(err).Error("Failed to add the environment definition")
			}
			if !c.FormatOverridden {
				c.OutputFormat = "text"
			}
			added := objects.EnvironmentList{}
			added = append(added, u)
			c.OutputData(&added, objects.TextOptions{
				NoHeaders: c.NoHeaders,
				Fields:    c.Fields,
			})
		} else {
			common.Logger.Error("Parameters passed in have failed checks.  Please review the warnings above")
		}
	},
}

func checkAddEnvironmentParams() bool {

	allParamsSet := true
	if len(environmentParam.Name) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --name/-e parameter must be set")
	}
	if len(environmentParam.Repository) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --repository/-r repository parameter must be set")
	}
	if allParamsSet {
		for _, v := range c.EnvDefs {
			showUtil := false
			u := objects.EnvironmentList{}
			common.Logger.Debugf("environmentParam.Name: %s, exising: %s", utilityParam.Name, v.Name)
			if environmentParam.Name == v.Name {
				allParamsSet = false
				showUtil = true
				u = append(u, v)
				common.Logger.Warn("The --name/-e environment name clashes with an existing environment name, please choose another, or use 'update environment' to update the existing environment")
			}
			if environmentParam.Repository == v.Repository {
				allParamsSet = false
				showUtil = true
				u = append(u, v)
				common.Logger.Warnf("The --repository/-r parameter clashes with an existing environment: %s, please consider using that utility definition", v.Name)
			}
			if showUtil {
				if !c.FormatOverridden {
					c.OutputFormat = "text"
				}
				c.OutputData(&u, objects.TextOptions{
					NoHeaders:  c.NoHeaders,
					ShowHidden: c.ShowHidden,
					Fields:     c.Fields,
				})
			}
		}
	}
	return allParamsSet
}

func addEnvironment() (objects.Environment, error) {

	newEnv := objects.Environment{
		Name:             environmentParam.Name,
		Repository:       environmentParam.Repository,
		Source:           "local",
		ExcludeFromShare: environmentParam.ExcludeFromShare,
	}

	c.EnvDefs = append(c.EnvDefs, newEnv)
	viper.Set("environments", c.EnvDefs)
	verr := viper.WriteConfig()
	if verr != nil {
		common.Logger.WithError(verr).Info("Failed to write config")
		return newEnv, verr
	}
	return newEnv, nil
}

func init() {
	addCmd.AddCommand(environmentCmd)

	environmentCmd.Flags().StringVarP(&environmentParam.Name, "name", "e", "", "Unique name for your environment definition")
	environmentCmd.Flags().StringVarP(&environmentParam.Repository, "repository", "r", "", "Repository for your environment, eg: us-docker.pkg.dev/my-lower-repo")
	environmentCmd.Flags().BoolVarP(&environmentParam.ExcludeFromShare, "exclude", "x", false, "Exclude from 'push' to central repository")
}
