package update

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
	Short:   updateEnvironmentHelp.Short(),
	Long:    updateEnvironmentHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		u, err := updateEnvironment()
		if err != nil {
			logrus.WithError(err).Error("Failed to update the environment definition")
		}
		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		updated := objects.EnvironmentList{}
		updated = append(updated, u)
		c.OutputData(&updated, objects.TextOptions{
			NoHeaders:     c.NoHeaders,
			Fields:        c.Fields,
			DefaultFields: c.OutputFieldsMap["environments"],
		})
	},
}

func updateEnvironment() (objects.Environment, error) {

	updatedEnvironment := false
	updatedDef := objects.Environment{}
	for i, v := range c.EnvDefs {
		common.Logger.Tracef("Checking environment: %s", v.Name)
		if environmentParam.Name == v.Name {
			common.Logger.Tracef("Updating environment: %s", v.Name)
			if len(environmentParam.Repository) > 0 {
				common.Logger.Tracef("Updating environment repository: %s", environmentParam.Repository)
				c.EnvDefs[i].Repository = environmentParam.Repository
				updatedEnvironment = true
			}
			if environmentParam.ExcludeFromShare {
				c.EnvDefs[i].ExcludeFromShare = !c.EnvDefs[i].ExcludeFromShare
				updatedEnvironment = true
			}
			if environmentParam.Hidden {
				c.EnvDefs[i].Hidden = !c.EnvDefs[i].Hidden
				updatedEnvironment = true
			}
			updatedDef = c.EnvDefs[i]
		}
	}
	if updatedEnvironment {
		viper.Set("environments", c.EnvDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			common.Logger.WithError(verr).Info("Failed to write config")
			return updatedDef, verr
		}
	} else {
		common.Logger.Warnf("The environment, %s, was not updated, perhaps a mis-matched name", environmentParam.Name)
	}

	return updatedDef, nil
}

func init() {
	updateCmd.AddCommand(environmentCmd)
	environmentCmd.Flags().StringVarP(&environmentParam.Name, "name", "e", "", "Unique name for your environment, eg: uppers or lowers")
	environmentCmd.Flags().StringVarP(&environmentParam.Repository, "repository", "r", "", "Repository for your environment, eg: us-docker.pkg.dev/my-lower-repo")
	environmentCmd.Flags().BoolVarP(&environmentParam.ExcludeFromShare, "toggle-exclude", "x", false, "Switch the current 'excludeFromShare' flag for the environment definition")
	environmentCmd.Flags().BoolVar(&environmentParam.Hidden, "toggle-hidden", false, "Switch the current 'Hidden' flag for the environment definition")
}
