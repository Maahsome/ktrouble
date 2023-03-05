package cmd

import (
	"fmt"
	"ktrouble/ask"
	"ktrouble/common"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull utility definitions from git",
	Long: `EXAMPLE:
  > ktrouble pull
`,
	Run: func(cmd *cobra.Command, args []string) {
		pullUtilityDefinitions()
	},
}

func pullUtilityDefinitions() {

	remoteDefs, remoteDefsMap := c.GitUpstream.GetNewUpstreamDefs(c.UtilDefs)

	if len(remoteDefs) > 0 {
		prompt := "Choose utility definitions to add to your local configuration:"
		addUtils := ask.PromptForUtilityList(remoteDefs, prompt)

		if len(addUtils) > 0 {
			for _, v := range addUtils {
				c.UtilDefs = append(c.UtilDefs, remoteDefsMap[v])
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
}
