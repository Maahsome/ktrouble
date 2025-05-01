package cmd

import (
	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/defaults"

	"github.com/spf13/cobra"
)

type DeleteParam struct {
	All bool
}

var deleteParam = DeleteParam{}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: defaults.DeleteAliases,
	Short:   deleteHelp.Short(),
	Long:    deleteHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if c.Client != nil {
			podList := c.Client.GetCreatedPods(deleteParam.All)

			switch count := len(podList.Items); {
			case count == 1:
				selectedPod := ask.PromptForPod(podList, "Choose a pod to delete:")

				c.Client.DeletePod(selectedPod)
				c.Client.DeleteAssociatedService(selectedPod)
				c.Client.DeleteAssociatedIngress(selectedPod)
			case count > 1:
				selectedPods := ask.PromptForPodList(podList, "Choose a pod to delete:")

				for _, p := range selectedPods {
					c.Client.DeletePod(p)
					c.Client.DeleteAssociatedService(p)
					c.Client.DeleteAssociatedIngress(p)
				}
			}
		} else {
			common.Logger.Warn("Cannot delete a pod, no valid kubernetes context")
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&deleteParam.All, "all", "a", false, "Choose from a list of running PODs from ALL users")

}
