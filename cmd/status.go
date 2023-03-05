package cmd

import (
	"fmt"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get a comparison of the local utility definitions with the upstream one",
	Long: `EXAMPLE:
  > ktrouble status
`,
	Run: func(cmd *cobra.Command, args []string) {
		status := utilityDefinitionStatus()
		c.OutputData(&status, objects.TextOptions{
			NoHeaders: c.NoHeaders,
			Fields:    c.Fields,
		})
	},
}

func utilityDefinitionStatus() objects.StatusList {
	status := objects.StatusList{}

	remoteDefs, remoteDefsMap := c.GitUpstream.GetUpstreamDefs()

	for _, l := range c.UtilDefs {
		if r, ok := remoteDefsMap[l.Name]; !ok {
			status = append(status, objects.Status{
				Name:    l.Name,
				Status:  "only local",
				Exclude: fmt.Sprintf("%t", l.ExcludeFromShare),
			})
		} else {
			s := compareDefs(l, r)
			status = append(status, objects.Status{
				Name:    l.Name,
				Status:  s,
				Exclude: fmt.Sprintf("%t", l.ExcludeFromShare),
			})
		}
	}
	for _, r := range remoteDefs {
		if _, ok := c.UtilMap[r.Name]; !ok {
			status = append(status, objects.Status{
				Name:    r.Name,
				Status:  "only remote",
				Exclude: "",
			})
		}
	}
	return status
}

func compareDefs(local objects.UtilityPod, remote objects.UtilityPod) string {
	if local.ExecCommand == remote.ExecCommand {
		if local.Repository == remote.Repository {
			return "same"
		}
	}
	return "different"
}
func init() {
	RootCmd.AddCommand(statusCmd)
}
