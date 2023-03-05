package cmd

import (
	"fmt"
	"ktrouble/objects"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// utilitiesCmd represents the namespace command
var utilitiesCmd = &cobra.Command{
	Use:     "utilities",
	Aliases: []string{"utility", "util", "container", "containers", "image", "images"},
	Short:   "Get a list of supported utility container images",
	Long: `EXAMPLE:
	> ktrouble get utilities
`,
	Run: func(cmd *cobra.Command, args []string) {

		utilDefs := []objects.UtilityPod{}
		err := viper.UnmarshalKey("utilityDefinitions", &utilDefs)
		if err != nil {
			logrus.Fatal("Error unmarshalling utility defs...")
		}
		if len(utilDefs) == 0 {
			utilDefs = defaultUtilityDefinitions()
		}

		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		fmt.Println(utilityPodDataToString(utilDefs, fmt.Sprintf("%#v", utilDefs)))

	},
}

func utilityPodDataToString(podData objects.UtilityPodList, raw string) string {

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
		return podData.ToTEXT(c.NoHeaders)
	default:
		return podData.ToTEXT(c.NoHeaders)
	}
}
func init() {
	getCmd.AddCommand(utilitiesCmd)
}
