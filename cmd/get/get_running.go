package get

import (
	"ktrouble/common"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// runningCmd represents the running command
var runningCmd = &cobra.Command{
	Use:     "running",
	Aliases: []string{"pods", "pod"},
	Short:   "Get a list of running pods",
	Long: `EXAMPLE:
  Get a list of PODs that are currently running on the current context kubernetes
  cluster that were created with the ktrouble utility.  If the 'enableBashLinks'
  config.yaml setting is 'true', a '<bash: ... >' command will be displayed,
  otherwise the SHELL path will be displayed.

  > ktrouble get running

    NAME                NAMESPACE       STATUS   EXEC
    basic-tools-e1df2f  common-tooling  Running  <bash:kubectl -n common-tooling exec -it basic-tools-e1df2f -- /bin/bash>

    NAME                NAMESPACE       STATUS   SHELL
    basic-tools-e1df2f  common-tooling  Running  /bin/bash
`,
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
				ShowExec:     c.EnableBashLinks,
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
