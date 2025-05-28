package remove

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// utilityParam is used to store command line parameters
var utilityParam objects.UtilityPod

// utilityCmd represents the utility command
var utilityCmd = &cobra.Command{
	Use:     "utility",
	Aliases: defaults.GetUtilitesAliases,
	Short:   removeUtilityHelp.Short(),
	Long:    removeUtilityHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if len(utilityParam.Name) > 0 {
			err := removeOrHideUtility()
			if err != nil {
				logrus.WithError(err).Error("Failed to remove the utility definition")
			}
			if !c.FormatOverridden {
				c.OutputFormat = "text"
			}
			c.OutputData(&c.UtilDefs, objects.TextOptions{
				NoHeaders:        c.NoHeaders,
				ShowHidden:       true,
				Fields:           c.Fields,
				AdditionalFields: []string{"HIDDEN", "REMOVE_UPSTREAM"},
				DefaultFields:    c.OutputFieldsMap["utility"],
			})
		}
	},
}

func removeOrHideUtility() error {

	updatedDefs := false
	for i, v := range c.UtilDefs {
		if utilityParam.Name == v.Name {
			updatedDefs = true
			if v.Source == "ktrouble-utils" {
				c.UtilDefs[i].Hidden = true
				if utilityParam.RemoveUpstream {
					c.UtilDefs[i].RemoveUpstream = true
				}
			} else {
				c.UtilDefs = objects.RemoveUtilIndex(c.UtilDefs, i)
			}
			break
		}
	}
	if updatedDefs {
		viper.Set("utilityDefinitions", c.UtilDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			common.Logger.WithError(verr).Info("Failed to write config")
			return verr
		}
	}

	return nil
}

func init() {
	removeCmd.AddCommand(utilityCmd)
	utilityCmd.Flags().StringVarP(&utilityParam.Name, "name", "u", "", "Unique name of your utility pod")
	utilityCmd.Flags().BoolVarP(&utilityParam.RemoveUpstream, "remove-upstream", "r", false, "Remove the utility pod from the upstream repository on next push")
}
