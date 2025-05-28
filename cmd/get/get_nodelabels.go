package get

import (
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// nodelabelsCmd represents the nodelabels command
var nodelabelsCmd = &cobra.Command{
	Use:     "nodelabels",
	Aliases: defaults.GetNodeLabelsAliases,
	Short:   getNodeLabelsHelp.Short(),
	Long:    getNodeLabelsHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		nodeLabels := objects.NodeLabels{}
		nodeLabels = c.NodeSelectorLabels
		c.OutputData(&nodeLabels, objects.TextOptions{
			NoHeaders:     c.NoHeaders,
			Fields:        c.Fields,
			DefaultFields: c.OutputFieldsMap["node_labels"],
		})
	},
}

func init() {
	getCmd.AddCommand(nodelabelsCmd)
}
