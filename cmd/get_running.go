package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runningCmd represents the namespace command
var runningCmd = &cobra.Command{
	Use:     "running",
	Aliases: []string{"pods", "pod"},
	Short:   "Get a list of running pods",
	Long: `EXAMPLE:
	> ktrouble get running
`,
	Run: func(cmd *cobra.Command, args []string) {

		podList := getCreatedPods()

		fmt.Printf("%-50s %s\n", "POD", "NS")
		fmt.Printf("%-50s %s\n", "----------------------------------------------", "---------------------")
		for _, v := range podList.Items {
			fmt.Printf("%-50s %s\n", v.Name, v.Namespace)
		}
	},
}

func init() {
	getCmd.AddCommand(runningCmd)
}
