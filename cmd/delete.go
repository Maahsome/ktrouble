package cmd

import (
	"ktrouble/ask"
	"ktrouble/common"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete PODs that have been created by ktrouble",
	Long: `EXAMPLE:
  Delete a running POD.  This will prompt with a list of PODs that are running
  and were launched using ktrouble

  > ktrouble delete
`,
	Run: func(cmd *cobra.Command, args []string) {
		if c.Client != nil {
			podList := c.Client.GetCreatedPods()

			switch count := len(podList.Items); {
			case count == 1:
				selectedPod := ask.PromptForPod(podList)

				c.Client.DeletePod(selectedPod)
			case count > 1:
				selectedPods := ask.PromptForPodList(podList)

				for _, p := range selectedPods {
					c.Client.DeletePod(p)
				}
			}
		} else {
			common.Logger.Warn("Cannot delete a pod, no valid kubernetes context")
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
