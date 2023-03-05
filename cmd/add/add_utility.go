package add

import (
	"ktrouble/common"
	"ktrouble/objects"

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
		if checkAddUtilityParams() {
			err := addUtility()
			if err != nil {
				common.Logger.WithError(err).Error("Failed to express the version")
			}
			if !c.FormatOverridden {
				c.OutputFormat = "text"
			}
			c.OutputData(&c.UtilDefs, objects.TextOptions{
				NoHeaders:    c.NoHeaders,
				ShowExec:     c.EnableBashLinks,
				UtilMap:      c.UtilMap,
				UniqIdLength: c.UniqIdLength,
			})
		}
	},
}

func checkAddUtilityParams() bool {

	allParamsSet := true
	if len(utilityParam.Name) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --name/-n parameter must be set")
	}
	if len(utilityParam.Repository) == 0 {
		allParamsSet = false
		common.Logger.Warn("The --repository/-r repository parameter must be set")
	}
	if allParamsSet {
		for _, v := range c.UtilDefs {
			showUtil := false
			u := objects.UtilityPodList{}
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
					NoHeaders:    c.NoHeaders,
					ShowExec:     c.EnableBashLinks,
					UtilMap:      c.UtilMap,
					UniqIdLength: c.UniqIdLength,
				})
			}
		}
	}
	return allParamsSet
}

func addUtility() error {

	c.UtilDefs = append(c.UtilDefs, objects.UtilityPod{
		Name:             utilityParam.Name,
		Repository:       utilityParam.Repository,
		ExecCommand:      utilityParam.ExecCommand,
		Source:           "local",
		ExcludeFromShare: utilityParam.ExcludeFromShare,
	})
	viper.Set("utilityDefinitions", c.UtilDefs)
	verr := viper.WriteConfig()
	if verr != nil {
		common.Logger.WithError(verr).Info("Failed to write config")
		return verr
	}
	return nil
}

func init() {
	addCmd.AddCommand(utilityCmd)

	utilityCmd.Flags().StringVarP(&utilityParam.Name, "name", "u", "", "Unique name for your utility pod")
	utilityCmd.Flags().StringVarP(&utilityParam.Repository, "repository", "r", "", "Repository and tag for your utility container, eg: cmaahs/basic-tools:latest")
	utilityCmd.Flags().StringVarP(&utilityParam.ExecCommand, "cmd", "c", "/bin/sh", "Default shell/command to use when 'exec'ing into the POD")
	utilityCmd.Flags().BoolVarP(&utilityParam.ExcludeFromShare, "exclude", "e", false, "Exclude from 'push' to central repository")
}
