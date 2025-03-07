package get

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// runningCmd represents the running command
var runningCmd = &cobra.Command{
	Use:     "running",
	Aliases: defaults.GetRunningAliases,
	Short:   getRunningHelp.Short(),
	Long:    getRunningHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {

		if c.Client != nil {
			podList := c.Client.GetCreatedPods()

			podData := objects.PodList{}
			for _, v := range podList.Items {
				status := string(v.Status.Phase)
				if v.DeletionTimestamp != nil {
					status = "Terminating"
				}
				podData = append(podData, objects.Pod{
					Name:      v.Name,
					Namespace: v.Namespace,
					Status:    status,
				})
			}

			c.OutputData(&podData, objects.TextOptions{
				NoHeaders:    c.NoHeaders,
				BashLinks:    c.EnableBashLinks,
				UtilMap:      c.UtilMap,
				UniqIdLength: c.UniqIdLength,
			})
		} else {
			common.Logger.Warn("Cannot fetch running pods, no valid kubernetes context")
		}

	},
}

func init() {
	getCmd.AddCommand(runningCmd)
}
