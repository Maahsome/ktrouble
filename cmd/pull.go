package cmd

import (
	"fmt"
	"ktrouble/ask"
	"ktrouble/common"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type PullParam struct {
	All bool
}

var pullParam = PullParam{}

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: pullHelp.Short(),
	Long:  pullHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		pullUtilityDefinitions()
	},
}

func pullUtilityDefinitions() {

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
		prompt := "Choose utility definitions to add/update to your local configuration:"
		addUtils := ask.PromptForUtilityList(remoteDefs, prompt)

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
						foundExisting = true
						break
					}
				}
				if !foundExisting {
					def := remoteDefsMap[v]
					def.Hidden = false
					c.UtilDefs = append(c.UtilDefs, def)
				}
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

}

func init() {
	RootCmd.AddCommand(pullCmd)
	pullCmd.Flags().BoolVarP(&pullParam.All, "all", "a", false, "Specify --all to list locally modified definitions as pull selections")
}
