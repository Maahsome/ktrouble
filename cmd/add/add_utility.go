package add

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// utilityParam is used to store command line parameters
var utilityParam objects.UtilityPod

// utilityCmd represents the utility command
var utilityCmd = &cobra.Command{
	Use:     "utility",
	Aliases: defaults.GetUtilitesAliases,
	Short:   addUtilityHelp.Short(),
	Long:    addUtilityHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if checkAddUtilityParams() {
			u, err := addUtility()
			if err != nil {
				common.Logger.WithError(err).Error("Failed to add the utility definition")
			}
			if !c.FormatOverridden {
				c.OutputFormat = "text"
			}
			added := objects.UtilityPodList{}
			added = append(added, u)
			c.OutputData(&added, objects.TextOptions{
				NoHeaders:        c.NoHeaders,
				ShowHidden:       c.ShowHidden,
				Fields:           c.Fields,
				AdditionalFields: []string{"EXCLUDED"},
			})
		} else {
			common.Logger.Error("Parameters passed in have failed checks.  Please review the warnings above")
		}
	},
}

func checkAddUtilityParams() bool {

	allParamsSet := true
	if len(utilityParam.Name) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --name/-u parameter must be set")
	}
	if len(utilityParam.Repository) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --repository/-r repository parameter must be set")
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
			if utilityParam.Name == v.Name {
				allParamsSet = false
				showUtil = true
				u = append(u, v)
				common.Logger.Warn("The --name/-n utility name clashes with an existing utility name, please choose another, or use 'update utility'")
			}
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
					NoHeaders:  c.NoHeaders,
					ShowHidden: c.ShowHidden,
					Fields:     c.Fields,
				})
			}
		}
	}
	return allParamsSet
}

func addUtility() (objects.UtilityPod, error) {

	newUtil := objects.UtilityPod{
		Name:              utilityParam.Name,
		Repository:        utilityParam.Repository,
		ExecCommand:       utilityParam.ExecCommand,
		Source:            "local",
		ExcludeFromShare:  utilityParam.ExcludeFromShare,
		RequireSecrets:    utilityParam.RequireSecrets,
		RequireConfigmaps: utilityParam.RequireConfigmaps,
		RemoveUpstream:    false,
		Environments:      utilityParam.Environments,
	}

	hintData := []byte{}

	if len(utilityParam.Hint) > 0 {
		// read the file into hintData
		var rerr error
		hintData, rerr = os.ReadFile(utilityParam.Hint)
		if rerr != nil {
			return objects.UtilityPod{}, rerr
		}

	}

	newUtil.Hint = string(hintData)

	c.UtilDefs = append(c.UtilDefs, newUtil)
	viper.Set("utilityDefinitions", c.UtilDefs)
	verr := viper.WriteConfig()
	if verr != nil {
		common.Logger.WithError(verr).Info("Failed to write config")
		return newUtil, verr
	}
	return newUtil, nil
}

func init() {
	addCmd.AddCommand(utilityCmd)

	utilityCmd.Flags().StringVarP(&utilityParam.Name, "name", "u", "", "Unique name for your utility pod")
	utilityCmd.Flags().StringVarP(&utilityParam.Repository, "repository", "r", "", "Repository and tag for your utility container, eg: cmaahs/basic-tools:latest")
	utilityCmd.Flags().StringVarP(&utilityParam.ExecCommand, "cmd", "c", "/bin/sh", "Default shell/command to use when 'exec'ing into the POD")
	utilityCmd.Flags().BoolVarP(&utilityParam.ExcludeFromShare, "exclude", "x", false, "Exclude from 'push' to central repository")
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
