package combine

import (
	"fmt"
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type utilityParam struct {
	UtilityPod     objects.UtilityPod
	Combine        []string
	RemoveUpstream bool
}

// utiltyParam is used to store command line parameters
var utilityP utilityParam

// utilityCmd represents the utility command
var utilityCmd = &cobra.Command{
	Use:     "utility",
	Aliases: defaults.GetUtilitesAliases,
	Short:   combineUtilityHelp.Short(),
	Long:    combineUtilityHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if len(utilityP.Combine) < 2 {
			fmt.Println("The --combine parameter must have at least two utility names to combine")
			os.Exit(0)
		}
		if len(utilityP.UtilityPod.Name) > 0 {
			err := combineUtility()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to combine the utility definitions")
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
		} else {
			fmt.Println("No utility definition specified, need -u/--name to be specified")
		}
	},
}

func combineUtility() error {

	newUtility := objects.UtilityPod{
		Name:   utilityP.UtilityPod.Name,
		Hidden: false,
	}
	if combinedUtilitiesExist() {
		copiedSettings := false
		// all of these should exist, we just checked them
		for _, name := range utilityP.Combine {
			for _, u := range c.UtilDefs {
				if u.Name == name {
					if name == newUtility.Name {
						return fmt.Errorf("the utility definition '%s' specified in the --combine parameter is the same as the new utility definition", name)
					}
					newUtility.Tags = append(newUtility.Tags, u.Tags...)
					if !copiedSettings {
						newUtility.Image = u.Image
						newUtility.ExecCommand = u.ExecCommand
						newUtility.RequireConfigmaps = u.RequireConfigmaps
						newUtility.RequireSecrets = u.RequireSecrets
						newUtility.ExcludeFromShare = u.ExcludeFromShare
						newUtility.Hint = u.Hint
						newUtility.Environments = append(newUtility.Environments, u.Environments...)
					}
					c.UtilDefs.UpdateProperty(u.Name, "hidden", true)
					c.UtilDefs.UpdateProperty(u.Name, "removeupstream", utilityP.RemoveUpstream)
					copiedSettings = true
					break
				}
			}
		}

		// We should have ALL the utility names from the Combine list
		c.UtilDefs = append(c.UtilDefs, newUtility)
		viper.Set("utilityDefinitions", c.UtilDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			common.Logger.WithError(verr).Info("Failed to write config")
			return verr
		}
	} else {
		return fmt.Errorf("one or more of the utility names in the --combine parameter do not exist in your local 'config.yaml' file")
	}
	return nil
}

func combinedUtilitiesExist() bool {
	for _, name := range utilityP.Combine {
		found := false
		for _, u := range c.UtilDefs {
			if u.Name == name {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func init() {
	combineCmd.AddCommand(utilityCmd)
	utilityCmd.Flags().StringVarP(&utilityP.UtilityPod.Name, "name", "u", "", "Unique name of your utility pod")
	utilityCmd.Flags().StringSliceVar(&utilityP.Combine, "combine", []string{}, "A comma-separated list of utility pod names to combine, eg, --combine 'mysql5,mysql8'")
	utilityCmd.Flags().BoolVarP(&utilityP.RemoveUpstream, "remove-upstream", "r", false, "Remove the combined utility pods from the upstream repository on next push")
}
