package cmd

import (
	"fmt"
	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/objects"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type PushParam struct {
	Environments bool
}

var pushParam = PushParam{}

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: pushHelp.Short(),
	Long:  pushHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if !c.GitUpstream.VersionDirectoryExists(fmt.Sprintf("v%d", c.Semver.Major)) {
			if c.Semver.Major == 0 {
				common.Logger.Error("The repository is not initialized, please create the repository usign git before running this command")
				return
			} else {
				common.Logger.Error("The version directory in the repository does not exist.  Please run 'ktrouble migrate' to migrate data to the new version")
				common.Logger.Error("The existing data will remain in the old version directory")
				return
			}
		}
		if pushParam.Environments {
			pushLocalEnvDefinitions()
		} else {
			pushLocalDefinitions()
		}
	},
}

func pushLocalEnvDefinitions() {

	// Get a list of "locals"
	localDefs := objects.EnvironmentList{}

	status := EnvironmentDefinitionStatus()

	for _, v := range status {
		def := c.EnvMap[v.Name]
		if def.Source == "local" && !def.ExcludeFromShare {
			def.Source = "ktrouble-utils"
			localDefs = append(localDefs, def)
		}
		if def.Source != "local" && !def.ExcludeFromShare && v.Status == "different" {
			localDefs = append(localDefs, def)
		}
	}

	if len(localDefs) > 0 {
		// Call gitupstream.PushLocals
		uploadConfig := objects.EnvironmentConfig{}
		uploadConfig.Environments = localDefs

		common.Logger.Tracef("uploadConfig: \n%#v", localDefs)
		if c.GitUpstream.PushEnvLocals(uploadConfig) {
			for _, v := range localDefs {
				for i, u := range c.EnvDefs {
					if v.Name == u.Name {
						c.EnvDefs[i].Source = "ktrouble-utils"
						if v.RemoveUpstream {
							c.EnvDefs = objects.RemoveEnvIndex(c.EnvDefs, i)
						}
						break
					}
				}
			}
			viper.Set("environments", c.EnvDefs)
			verr := viper.WriteConfig()
			if verr != nil {
				common.Logger.WithError(verr).Info("Failed to write config")
			}
		} else {
			common.Logger.Error("failed to push to repository")
		}
	} else {
		fmt.Println("No changed definitions to push")
	}
}

func pushLocalDefinitions() {

	// Get a list of "locals"
	localDefs := objects.UtilityPodList{}

	status := UtilityDefinitionStatus()

	for _, v := range status {
		def := c.UtilMap[v.Name]
		if def.Source == "local" && !def.ExcludeFromShare {
			localDefs = append(localDefs, def)
		}
		if def.Source != "local" && !def.ExcludeFromShare && v.Status == "different" {
			localDefs = append(localDefs, def)
		}
	}

	prompt := "Choose utilities to submit to the remote repository:"
	pushList := ask.PromptForUtilityList(localDefs, prompt)

	if len(pushList) > 0 {
		// Call gitupstream.PushLocals
		uploadDefs := objects.UtilityPodList{}
		for _, v := range pushList {
			uploadDefs = append(uploadDefs, c.UtilMap[v])
		}
		if c.GitUpstream.PushLocals(uploadDefs) {
			for _, v := range pushList {
				for i, u := range c.UtilDefs {
					if v == u.Name {
						c.UtilDefs[i].Source = "ktrouble-utils"
						if c.UtilMap[v].RemoveUpstream {
							c.UtilDefs = objects.RemoveUtilIndex(c.UtilDefs, i)
						}
						break
					}
				}
			}
			viper.Set("utilityDefinitions", c.UtilDefs)
			verr := viper.WriteConfig()
			if verr != nil {
				common.Logger.WithError(verr).Info("Failed to write config")
			}
		} else {
			common.Logger.Error("failed to push to repository")
		}
	} else {
		fmt.Println("No definitions selected")
	}

}

func init() {
	RootCmd.AddCommand(pushCmd)
	pushCmd.Flags().BoolVar(&pushParam.Environments, "env", false, "Use this switch to operate on the environment definitions")
}
