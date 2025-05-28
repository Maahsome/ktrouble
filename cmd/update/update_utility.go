package update

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"
	"os"

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
	Short:   updateUtilityHelp.Short(),
	Long:    updateUtilityHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if checkUpdateUtilityParams() {
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
				AdditionalFields: []string{"ENVIRONMENTS", "HIDDEN", "EXCLUDED", "REQUIRECONFIGMAPS", "REQUIRESECRETS"},
				DefaultFields:    c.OutputFieldsMap["utility"],
			})
		} else {
			common.Logger.Error("Parameters passed in have failed checks.  Please review the warnings above")
		}
	},
}

func checkUpdateUtilityParams() bool {

	allParamsSet := true
	if len(utilityParam.Name) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --name/-u parameter must be set")
	}
	if len(utilityParam.Hint) > 0 {
		if !fileExists(utilityParam.Hint) {
			allParamsSet = false
			common.Logger.Warn("The file pointed to by the --hint-file must exist")
		}
	}
	if len(utilityParam.Environments) > 0 {
		common.Logger.Tracef("Environments: %s", utilityParam.Environments)
		if !objects.EnvironmentsExist(c.EnvMap, utilityParam.Environments) {
			allParamsSet = false
			common.Logger.Warnf("The --environments parameter must be set to valid environment names, %s", utilityParam.Environments)
			common.Logger.Warn("Please use 'get environments' to see the list of valid environment names")
		}
	}
	if allParamsSet {
		for _, v := range c.UtilDefs {
			showUtil := false
			u := objects.UtilityPodList{}
			common.Logger.Debugf("utilityParam.Name: %s, exising: %s", utilityParam.Name, v.Name)
			if utilityParam.Repository == v.Repository {
				allParamsSet = false
				showUtil = true
				u = append(u, v)
				common.Logger.Warnf("The --repository/-r parameter clashes with an existing utility: %s, please consider using that utility definition", v.Name)
			}
			if showUtil {
				if !c.FormatOverridden {
					c.OutputFormat = "text"
				}
				c.OutputData(&u, objects.TextOptions{
					NoHeaders:     c.NoHeaders,
					ShowHidden:    c.ShowHidden,
					Fields:        c.Fields,
					DefaultFields: c.OutputFieldsMap["utility"],
				})
			}
		}
	}
	return allParamsSet
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
			if utilityParam.RequireSecrets {
				c.UtilDefs[i].RequireSecrets = !c.UtilDefs[i].RequireSecrets
				updatedUtilty = true
			}
			if utilityParam.RequireConfigmaps {
				c.UtilDefs[i].RequireConfigmaps = !c.UtilDefs[i].RequireConfigmaps
				updatedUtilty = true
			}
			if len(utilityParam.Hint) > 0 {
				hintFile, err := os.ReadFile(utilityParam.Hint)
				if err != nil {
					common.Logger.WithError(err).Error("Error reading hint file")
					return updatedDef, err
				}
				c.UtilDefs[i].Hint = string(hintFile)
				updatedUtilty = true
			}
			if len(utilityParam.Environments) > 0 {
				c.UtilDefs[i].Environments = utilityParam.Environments
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
		common.Logger.Warnf("The utility, %s, was not updated, perhaps a mis-matched name", utilityParam.Name)
	}

	return updatedDef, nil
}

func init() {
	updateCmd.AddCommand(utilityCmd)
	utilityCmd.Flags().StringVarP(&utilityParam.Name, "name", "u", "", "Unique name for your utility pod")
	utilityCmd.Flags().StringVarP(&utilityParam.Repository, "repository", "r", "", "Repository and tag for your utility container, eg: cmaahs/basic-tools:latest")
	utilityCmd.Flags().StringVarP(&utilityParam.ExecCommand, "cmd", "c", "", "Default shell/command to use when 'exec'ing into the POD")
	utilityCmd.Flags().BoolVarP(&utilityParam.ExcludeFromShare, "toggle-exclude", "x", false, "Switch the current 'excludeFromShare' flag for the utility definition")
	utilityCmd.Flags().BoolVar(&utilityParam.Hidden, "toggle-hidden", false, "Switch the current 'hidden' flag for the utility definition")
	utilityCmd.Flags().BoolVar(&utilityParam.RequireSecrets, "require-secrets", false, "Set the Utilty to always prompt for secrets")
	utilityCmd.Flags().BoolVar(&utilityParam.RequireConfigmaps, "require-configmaps", false, "Set the Utilty to always prompt for configmaps")
	utilityCmd.Flags().StringVar(&utilityParam.Hint, "hint-file", "", "Specify a file containing the hint text")
	utilityCmd.Flags().StringSliceVarP(&utilityParam.Environments, "environments", "e", []string{}, "Specify an array of environment names: eg, --environments 'lowers,uppers'")
}

// fileExists checks if file exists
func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
