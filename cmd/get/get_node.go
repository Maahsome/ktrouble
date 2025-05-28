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
	Short:   getNodeHelp.Short(),
	Long:    getNodeHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {

		if c.Client != nil {
			nodeList := c.Client.GetNodes()

			nodeData := objects.NodeList{}
			for _, v := range nodeList.Items {
				nodeData.Node = append(nodeData.Node, v.Name)
			}

			c.OutputData(&nodeData, objects.TextOptions{
				NoHeaders:     c.NoHeaders,
				Fields:        c.Fields,
				DefaultFields: c.OutputFieldsMap["node"],
			})
		} else {
			common.Logger.Warn("Cannot fetch nodes, no valid kubernetes context")
		}

	},
}

func init() {
	getCmd.AddCommand(nodeCmd)
}
