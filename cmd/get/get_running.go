package get

import (
	"fmt"
	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

type GetRunningParam struct {
	All bool
}

var getRunningParam = GetRunningParam{}

// runningCmd represents the running command
var runningCmd = &cobra.Command{
	Use:     "running",
	Aliases: defaults.GetRunningAliases,
	Short:   getRunningHelp.Short(),
	Long:    getRunningHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {

		if c.Client != nil {
			podList := c.Client.GetCreatedPods(getRunningParam.All)

			podData := objects.PodList{}
			for _, v := range podList.Items {
				status := string(v.Status.Phase)
				if v.DeletionTimestamp != nil {
					status = "Terminating"
				}
				service := c.Client.GetAssociatedService(ask.PodDetail{
					Name:      v.Name,
					Namespace: v.Namespace,
				})
				serviceName := ""
				servicePort := ""
				if len(service.Items) > 0 {
					serviceName = service.Items[0].Name
					servicePort = fmt.Sprintf("%d", service.Items[0].Spec.Ports[0].TargetPort.IntVal)
				}
				podData = append(podData, objects.Pod{
					Name:          v.Name,
					Namespace:     v.Namespace,
					Status:        status,
					LaunchedBy:    v.Labels["launchedby"],
					Service:       serviceName,
					ServicePort:   servicePort,
					ContainerName: "",
				})
			}

			c.OutputData(&podData, objects.TextOptions{
				NoHeaders:    c.NoHeaders,
				BashLinks:    c.EnableBashLinks,
				UtilMap:      c.UtilMap,
				UniqIdLength: c.UniqIdLength,
				Fields:       c.Fields,
			})
		} else {
			common.Logger.Warn("Cannot fetch running pods, no valid kubernetes context")
		}

	},
}

func init() {
	getCmd.AddCommand(runningCmd)
	runningCmd.Flags().BoolVarP(&getRunningParam.All, "all", "a", false, "List running PODs from ALL users")
}
