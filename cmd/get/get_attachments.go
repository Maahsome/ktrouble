package get

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

type GetAttachmentsParam struct {
	All bool
}

var getAttachmentsParam = GetAttachmentsParam{}

// attachmentsCmd represents the running command
var attachmentsCmd = &cobra.Command{
	Use:     "attachments",
	Aliases: defaults.GetAttachmentsAliases,
	Short:   "Get a list of attached containers",
	Long: `EXAMPLE:
  Get a list of utility containers that are attached to existing PODS that are
  currently running on the current context kubernetes cluster that were attached
  with the ktrouble utility.  If the 'enableBashLinks' config.yaml setting is
  'true', a '<bash: ... >' command will be displayed, otherwise the SHELL path
  will be displayed.

  > ktrouble get attachments

    NAME                NAMESPACE       STATUS   EXEC
    basic-tools-e1df2f  common-tooling  Running  <bash:kubectl -n common-tooling exec -it basic-tools-e1df2f -- /bin/bash>

    NAME                NAMESPACE       STATUS   SHELL
    basic-tools-e1df2f  common-tooling  Running  /bin/bash
`,
	Run: func(cmd *cobra.Command, args []string) {

		if c.Client != nil {
			podList := c.Client.GetAttachedContainers(getAttachmentsParam.All)

			podData := objects.PodList{}
			for _, v := range podList.Items {
				for _, es := range v.Status.EphemeralContainerStatuses {
					if es.State.Running != nil {
						podData = append(podData, objects.Pod{
							Name:          v.Name,
							Namespace:     v.Namespace,
							Status:        "Running",
							LaunchedBy:    v.Labels["ktrouble.launchedby"],
							Service:       "",
							ServicePort:   "",
							ContainerName: es.Name,
						})
					}
				}
			}

			c.OutputData(&podData, objects.TextOptions{
				NoHeaders:    c.NoHeaders,
				BashLinks:    c.EnableBashLinks,
				UtilMap:      c.UtilMap,
				UniqIdLength: c.UniqIdLength,
			})
		} else {
			common.Logger.Warn("Cannot fetch attached containers, no valid kubernetes context")
		}

	},
}

func init() {
	getCmd.AddCommand(attachmentsCmd)
	attachmentsCmd.Flags().BoolVarP(&getAttachmentsParam.All, "all", "a", false, "List attached containers for ALL users")
}
