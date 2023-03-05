package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// utilitiesCmd represents the namespace command
var utilitiesCmd = &cobra.Command{
	Use:     "utilities",
	Aliases: []string{"utility", "util", "container", "containers", "image", "images"},
	Short:   "Get a list of supported utility container images",
	Long: `EXAMPLE:
	> ktrouble get utilities
`,
	Run: func(cmd *cobra.Command, args []string) {

		utilDefs := []UtilityPod{}
		err := viper.UnmarshalKey("utilityDefinitions", &utilDefs)
		if err != nil {
			logrus.Fatal("Error unmarshalling utility defs...")
		}
		if len(utilDefs) == 0 {
			utilDefs = defaultUtilityDefinitions()
		}

		fmt.Printf("%-15s %-50s %s\n", "UTILITY", "REGISTRY", "EXEC_CMD")
		fmt.Printf("%-15s %-50s %s\n", "---------------", "----------------------------------------------", "---------------------")
		for _, v := range utilDefs {
			fmt.Printf("%-15s %-50s %s\n", v.Name, v.Repository, v.ExecCommand)
		}
	},
}

func init() {
	getCmd.AddCommand(utilitiesCmd)
}
