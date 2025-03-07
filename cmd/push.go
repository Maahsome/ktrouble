package cmd

import (
	"fmt"
	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/objects"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: pushHelp.Short(),
	Long:  pushHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		pushLocalDefinitions()
	},
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
}
