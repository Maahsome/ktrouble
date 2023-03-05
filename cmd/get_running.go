package cmd

import (
	"fmt"
	"ktrouble/objects"
	"strings"

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

		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		fmt.Println(podDataToString(podData, fmt.Sprintf("%#v", podList.Items)))

	},
}

func podDataToString(podData objects.PodList, raw string) string {

	switch strings.ToLower(c.OutputFormat) {
	case "raw":
		return raw
	case "json":
		return podData.ToJSON()
	case "gron":
		return podData.ToGRON()
	case "yaml":
		return podData.ToYAML()
	case "text", "table":
		return podData.ToTEXT(c.NoHeaders, c.EnableBashLinks, utilMap, uniqIdLength)
	default:
		return podData.ToTEXT(c.NoHeaders, c.EnableBashLinks, utilMap, uniqIdLength)
	}
}

func init() {
	getCmd.AddCommand(runningCmd)
}
