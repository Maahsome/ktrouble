package remove

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
	Short: "Remove a utility from the config file, or HIDE it if it is an upstream definition",
	Run: func(cmd *cobra.Command, args []string) {
		if len(utilityParam.Name) > 0 {
			err := removeOrHideUtility()
			if err != nil {
				logrus.WithError(err).Error("Failed to express the version")
			}
			if !c.FormatOverridden {
				c.OutputFormat = "text"
			}
			c.OutputData(&c.UtilDefs, objects.TextOptions{
				NoHeaders:    c.NoHeaders,
				ShowExec:     c.EnableBashLinks,
				UtilMap:      c.UtilMap,
				UniqIdLength: c.UniqIdLength,
				ShowHidden:   c.ShowHidden,
			})
		}
	},
}

func removeOrHideUtility() error {

	updatedDefs := false
	for i, v := range c.UtilDefs {
		if utilityParam.Name == v.Name {
			updatedDefs = true
			// this is the one to delete
			if v.Source == "ktrouble-utils" {
				// hide it
				c.UtilDefs[i].Hidden = true
			} else {
				// remove it from the list
				c.UtilDefs = removeIndex(c.UtilDefs, i)
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

func removeIndex(s objects.UtilityPodList, index int) objects.UtilityPodList {
	ret := make(objects.UtilityPodList, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func init() {
	removeCmd.AddCommand(utilityCmd)
	utilityCmd.Flags().StringVarP(&utilityParam.Name, "name", "u", "", "Unique name for your utility pod")
}
