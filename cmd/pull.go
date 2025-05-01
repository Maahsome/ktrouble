package cmd

import (
	"fmt"
	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/objects"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type PullParam struct {
	All          bool
	Environments bool
	Utilities    []string
}

var pullParam = PullParam{}

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: pullHelp.Short(),
	Long:  pullHelp.Long(),
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
		if pullParam.Environments {
			status := pullEnvironmentDefinitions()
			if len(status) > 0 {
				c.OutputData(&status, objects.TextOptions{
					NoHeaders: c.NoHeaders,
					Fields:    c.Fields,
				})
			}
		} else {
			status := pullUtilityDefinitions()
			if len(status) > 0 {
				c.OutputData(&status, objects.TextOptions{
					NoHeaders: c.NoHeaders,
					Fields:    c.Fields,
				})
			}
		}
	},
}

func pullEnvironmentDefinitions() objects.StatusList {

	status := objects.StatusList{}
	// TODO: duplicate pullUtilityDefinitions but for environments
	remoteDefs, remoteDefsMap := c.GitUpstream.GetNewUpstreamEnvDefs(c.EnvDefs)

	if pullParam.All {
		status := EnvironmentDefinitionStatus()
		for _, v := range status {
			def := c.EnvMap[v.Name]
			if def.Source != "local" && !def.ExcludeFromShare && v.Status == "different" {
				remoteDefs = append(remoteDefs, def)
			}
		}
	}

	if len(remoteDefs) > 0 {
		for _, v := range remoteDefs {
			foundExisting := false
			for i, u := range c.EnvDefs {
				if v.Name == u.Name {
					common.Logger.Tracef("Found existing environment definition %s", v.Name)
					c.EnvDefs[i].Repository = remoteDefsMap[v.Name].Repository
					c.EnvDefs[i].ExcludeFromShare = remoteDefsMap[v.Name].ExcludeFromShare
					foundExisting = true
					break
				}
			}
			if !foundExisting {
				def := remoteDefsMap[v.Name]
				c.EnvDefs = append(c.EnvDefs, def)
				status = append(status, objects.Status{
					Name:    def.Name,
					Status:  "added",
					Exclude: fmt.Sprintf("%t", def.ExcludeFromShare),
				})
			}
		}
		viper.Set("environments", c.EnvDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			common.Logger.WithError(verr).Info("Failed to write config")
		}
	} else {
		fmt.Println("Up to date")
	}

	return status
}

func pullUtilityDefinitions() objects.StatusList {

	status := objects.StatusList{}
	addUtils := []string{}

	remoteDefs, remoteDefsMap := c.GitUpstream.GetNewUpstreamDefs(c.UtilDefs)

	if pullParam.All {
		status := UtilityDefinitionStatus()
		for _, v := range status {
			def := c.UtilMap[v.Name]
			if def.Source != "local" && !def.ExcludeFromShare && v.Status == "different" {
				remoteDefs = append(remoteDefs, def)
			}
		}
	}

	if len(remoteDefs) > 0 {
		if len(pullParam.Utilities) == 0 {
			prompt := "Choose utility definitions to add/update to your local configuration:"
			addUtils = ask.PromptForUtilityList(remoteDefs, prompt)
		} else {
			if paramUtilitiesExist(remoteDefs) {
				addUtils = pullParam.Utilities
			} else {
				common.Logger.Fatal("Some of the utility definition names you specified do not exist in the remote repository, or are aleady the same locally, use 'ktrouble status' to get a list of 'different' utilities")
			}
		}

		if len(addUtils) > 0 {
			for _, v := range addUtils {
				foundExisting := false
				for i, u := range c.UtilDefs {
					if v == u.Name {
						c.UtilDefs[i].Repository = remoteDefsMap[v].Repository
						c.UtilDefs[i].Source = remoteDefsMap[v].Source
						c.UtilDefs[i].ExecCommand = remoteDefsMap[v].ExecCommand
						c.UtilDefs[i].ExcludeFromShare = remoteDefsMap[v].ExcludeFromShare
						c.UtilDefs[i].Hidden = remoteDefsMap[v].Hidden
						c.UtilDefs[i].Environments = remoteDefsMap[v].Environments
						foundExisting = true
						break
					}
				}
				pullStatus := "added"
				def := remoteDefsMap[v]
				def.Hidden = false
				if foundExisting {
					pullStatus = "updated"
				} else {
					c.UtilDefs = append(c.UtilDefs, def)
				}
				status = append(status, objects.Status{
					Name:    def.Name,
					Status:  pullStatus,
					Exclude: fmt.Sprintf("%t", def.ExcludeFromShare),
				})
			}
			viper.Set("utilityDefinitions", c.UtilDefs)
			verr := viper.WriteConfig()
			if verr != nil {
				common.Logger.WithError(verr).Info("Failed to write config")
			}
		} else {
			fmt.Println("No definitions selected")
		}
	} else {
		fmt.Println("Up to date")
	}

	return status

}

func paramUtilitiesExist(remoteDefs objects.UtilityPodList) bool {
	for _, v := range pullParam.Utilities {
		// check v.Name exists in remoteDefs
		found := false
		for _, u := range remoteDefs {
			if v == u.Name {
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
	RootCmd.AddCommand(pullCmd)
	pullCmd.Flags().BoolVarP(&pullParam.All, "all", "a", false, "Specify --all to list locally modified definitions as pull selections")
	pullCmd.Flags().BoolVar(&pullParam.Environments, "env", false, "Use this switch to operate on the environment definitions")
	pullCmd.Flags().StringSliceVarP(&pullParam.Utilities, "utilities", "u", []string{}, "Specify an array of utility names to pull: eg, --utilities 'basic-tools,dns-tools', default is to prompt")
}
