package get

import (
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// nodelabelsCmd represents the nodelabels command
var nodelabelsCmd = &cobra.Command{
	Use:     "nodelabels",
	Aliases: []string{"nodelabel", "nl", "labels"},
	Short:   "Get a list of defined node labels in config.yaml",
	Long: `EXAMPLE:
  Show the list of node labels in the configuration file

  > ktrouble get nodelabels
`,
	Run: func(cmd *cobra.Command, args []string) {
		nodeLabels := objects.NodeLabels{}
		nodeLabels = c.NodeSelectorLabels
		c.OutputData(&nodeLabels, objects.TextOptions{NoHeaders: c.NoHeaders})
	},
}

func init() {
	getCmd.AddCommand(nodelabelsCmd)
}
