package remove

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// environmentParam is used to store command line parameters
var environmentParam objects.Environment

// environmentCmd represents the utility command
var environmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: defaults.EnvironmentAliases,
	Short:   removeEnvironmentHelp.Short(),
	Long:    removeEnvironmentHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if len(environmentParam.Name) > 0 {
			err := removeOrHideEnvironment()
			if err != nil {
				logrus.WithError(err).Error("Failed to remove the environment definition")
			}
			if !c.FormatOverridden {
				c.OutputFormat = "text"
			}
			c.OutputData(&c.EnvDefs, objects.TextOptions{
				NoHeaders:        c.NoHeaders,
				ShowHidden:       true,
				Fields:           c.Fields,
				AdditionalFields: []string{"HIDDEN", "REMOVE_UPSTREAM"},
				DefaultFields:    c.OutputFieldsMap["environments"],
			})
		} else {
			logrus.Warn("--name/-e environment name must be specified")
		}
	},
}

func removeOrHideEnvironment() error {

	updatedDefs := false
	for i, v := range c.EnvDefs {
		if environmentParam.Name == v.Name {
			updatedDefs = true
			common.Logger.WithField("name", v.Name).Tracef("Hiding environment definition")
			if environmentParam.RemoveUpstream {
				c.EnvDefs[i].RemoveUpstream = true
				c.EnvDefs[i].Hidden = true
			} else {
				common.Logger.WithField("name", v.Name).Tracef("Removing environment definition")
				c.EnvDefs = objects.RemoveEnvIndex(c.EnvDefs, i)
			}
			break
		}
	}
	if updatedDefs {
		viper.Set("environments", c.EnvDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			common.Logger.WithError(verr).Info("Failed to write config")
			return verr
		}
	}

	return nil
}

func init() {
	removeCmd.AddCommand(environmentCmd)
	environmentCmd.Flags().StringVarP(&environmentParam.Name, "name", "e", "", "Unique name of your environment")
	environmentCmd.Flags().BoolVarP(&environmentParam.RemoveUpstream, "remove-upstream", "r", false, "Remove the environment from the upstream repository on next push")
}
