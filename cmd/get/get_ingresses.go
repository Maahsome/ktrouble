package get

import (
	"fmt"
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

type GetIngressesParam struct {
	All bool
}

var getIngressesParam = GetIngressesParam{}

// ingressesCmd represents the running command
var ingressesCmd = &cobra.Command{
	Use:     "ingresses",
	Aliases: defaults.GetIngressesAliases,
	Short:   "Get a list of ktrouble installed ingresses",
	Long: `EXAMPLE
Get a list of ingresses that are installed by ktrouble

  > ktrouble get ingresses

`,
	Run: func(cmd *cobra.Command, args []string) {

		if c.Client != nil {
			ingressList := c.Client.GetCreatedIngresses(getIngressesParam.All)

			ingressData := objects.IngressList{}
			for _, v := range ingressList.Items {
				ingressData = append(ingressData, objects.Ingress{
					Name:       v.Name,
					Namespace:  v.Namespace,
					Class:      "nginx", // *v.Spec.IngressClassName,
					Hosts:      fmt.Sprintf("https://%s%s", v.Spec.Rules[0].Host, v.Spec.Rules[0].HTTP.Paths[0].Path),
					Address:    v.Status.LoadBalancer.Ingress[0].IP,
					Ports:      fmt.Sprintf("%d", v.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Port.Number),
					LaunchedBy: v.Labels["launchedby"],
				})
			}

			c.OutputData(&ingressData, objects.TextOptions{
				NoHeaders:     c.NoHeaders,
				BashLinks:     c.EnableBashLinks,
				UtilMap:       c.UtilMap,
				UniqIdLength:  c.UniqIdLength,
				Fields:        c.Fields,
				DefaultFields: c.OutputFieldsMap["ingress"],
			})
		} else {
			common.Logger.Warn("Cannot fetch installed ingresses, no valid kubernetes context")
		}
	},
}

func init() {
	getCmd.AddCommand(ingressesCmd)
	ingressesCmd.Flags().BoolVarP(&getIngressesParam.All, "all", "a", false, "List installed ingreses from ALL users")
}
