package get

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:     "node",
	Aliases: defaults.GetNodesAliases,
	Short:   "Get a list of node labels",
	Long: `EXAMPLE:
  Get a list of nodes for the current context cluster

  > ktrouble get node
`,
	Run: func(cmd *cobra.Command, args []string) {

		if c.Client != nil {
			nodeList := c.Client.GetNodes()

			nodeData := objects.NodeList{}
			for _, v := range nodeList.Items {
				nodeData.Node = append(nodeData.Node, v.Name)
			}

			c.OutputData(&nodeData, objects.TextOptions{NoHeaders: c.NoHeaders})
		} else {
			common.Logger.Warn("Cannot fetch nodes, no valid kubernetes context")
		}

	},
}

func init() {
	getCmd.AddCommand(nodeCmd)
}
