package cmd

import (
	"fmt"
	"ktrouble/objects"
	"strings"

	"github.com/spf13/cobra"
)

// sizesCmd represents the namespace command
var sizesCmd = &cobra.Command{
	Use:     "sizes",
	Aliases: []string{"size", "requests", "request", "limit", "limits"},
	Short:   "Get a list of defined sizes",
	Long: `EXAMPLE:
	> ktrouble get sizes
`,
	Run: func(cmd *cobra.Command, args []string) {

		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		fmt.Println(sizeDataToString(sizeDefs, fmt.Sprintf("%#v", sizeDefs)))

	},
}

func sizeDataToString(sizeData objects.ResourceSizeList, raw string) string {

	switch strings.ToLower(c.OutputFormat) {
	case "raw":
		return raw
	case "json":
		return sizeData.ToJSON()
	case "gron":
		return sizeData.ToGRON()
	case "yaml":
		return sizeData.ToYAML()
	case "text", "table":
		return sizeData.ToTEXT(c.NoHeaders)
	default:
		return sizeData.ToTEXT(c.NoHeaders)
	}
}
func init() {
	getCmd.AddCommand(sizesCmd)
}
