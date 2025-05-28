package get

import (
	"fmt"
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

type GetServicesParam struct {
	All bool
}

var getServicesParam = GetServicesParam{}

// servicesCmd represents the running command
var servicesCmd = &cobra.Command{
	Use:     "services",
	Aliases: defaults.GetServicesAliases,
	Short:   "Get a list of ktrouble installed services",
	Long: `EXAMPLE
Get a list of services that are installed by ktrouble

  > ktrouble get services

`,
	Run: func(cmd *cobra.Command, args []string) {

		if c.Client != nil {
			serviceList := c.Client.GetCreatedServices(getServicesParam.All)

			serviceData := objects.ServiceList{}
			for _, v := range serviceList.Items {
				externalIP := "<none>"
				if len(v.Spec.ExternalIPs) > 0 {
					externalIP = v.Spec.ExternalIPs[0]
				}
				serviceData = append(serviceData, objects.Service{
					Name:        v.Name,
					Namespace:   v.Namespace,
					ServiceType: string(v.Spec.Type),
					ClusterIP:   v.Spec.ClusterIP,
					ExternalIP:  externalIP,
					Ports:       fmt.Sprintf("%d/%s -> %s", v.Spec.Ports[0].Port, v.Spec.Ports[0].Protocol, &v.Spec.Ports[0].TargetPort),
					LaunchedBy:  v.Labels["launchedby"],
				})
			}

			c.OutputData(&serviceData, objects.TextOptions{
				NoHeaders:     c.NoHeaders,
				BashLinks:     c.EnableBashLinks,
				UtilMap:       c.UtilMap,
				UniqIdLength:  c.UniqIdLength,
				Fields:        c.Fields,
				DefaultFields: c.OutputFieldsMap["service"],
			})
		} else {
			common.Logger.Warn("Cannot fetch installed services, no valid kubernetes context")
		}
	},
}

func init() {
	getCmd.AddCommand(servicesCmd)
	servicesCmd.Flags().BoolVarP(&getServicesParam.All, "all", "a", false, "List installed services from ALL users")
}
