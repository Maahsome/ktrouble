package update

import (
	"ktrouble/common"
	"ktrouble/objects"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// utilityParam is used to store command line parameters
var utilityParam objects.UtilityPod

// utilityCmd represents the utility command
var utilityCmd = &cobra.Command{
	Use:   "utility",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		u, err := updateUtility()
		if err != nil {
			logrus.WithError(err).Error("Failed to update the utility definition")
		}
		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		updated := objects.UtilityPodList{}
		updated = append(updated, u)
		c.OutputData(&updated, objects.TextOptions{
			NoHeaders:        c.NoHeaders,
			ShowHidden:       true,
			Fields:           c.Fields,
			AdditionalFields: []string{"HIDDEN", "EXCLUDED"},
		})
	},
}

func updateUtility() (objects.UtilityPod, error) {

	updatedUtilty := false
	updatedDef := objects.UtilityPod{}
	for i, v := range c.UtilDefs {
		if utilityParam.Name == v.Name {
			if len(utilityParam.Repository) > 0 {
				c.UtilDefs[i].Repository = utilityParam.Repository
				updatedUtilty = true
			}
			if len(utilityParam.ExecCommand) > 0 {
				c.UtilDefs[i].ExecCommand = utilityParam.ExecCommand
				updatedUtilty = true
			}
			if utilityParam.ExcludeFromShare {
				c.UtilDefs[i].ExcludeFromShare = !c.UtilDefs[i].ExcludeFromShare
				updatedUtilty = true
			}
			if utilityParam.Hidden {
				c.UtilDefs[i].Hidden = !c.UtilDefs[i].Hidden
				updatedUtilty = true
			}
			updatedDef = c.UtilDefs[i]
		}
	}
	if updatedUtilty {
		viper.Set("utilityDefinitions", c.UtilDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			common.Logger.WithError(verr).Info("Failed to write config")
			return updatedDef, verr
		}
	} else {
		common.Logger.Warn("The utility, %s, was not updated, perhaps a mis-matched name")
	}

	return updatedDef, nil
}

func init() {
	updateCmd.AddCommand(utilityCmd)
	utilityCmd.Flags().StringVarP(&utilityParam.Name, "name", "u", "", "Unique name for your utility pod")
	utilityCmd.Flags().StringVarP(&utilityParam.Repository, "repository", "r", "", "Repository and tag for your utility container, eg: cmaahs/basic-tools:latest")
	utilityCmd.Flags().StringVarP(&utilityParam.ExecCommand, "cmd", "c", "", "Default shell/command to use when 'exec'ing into the POD")
	utilityCmd.Flags().BoolVarP(&utilityParam.ExcludeFromShare, "toggle-exclude", "e", false, "Switch the current 'excludeFromShare' flag for the utility definition")
	utilityCmd.Flags().BoolVar(&utilityParam.Hidden, "toggle-hidden", false, "Switch the current 'hidden' flag for the utility definition")

}
