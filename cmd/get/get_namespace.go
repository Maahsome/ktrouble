package get

import (
	"ktrouble/common"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"namespaces", "ns"},
	Short:   "Get a list of namespaces",
	Long: `EXAMPLE:
  Return a list of kubernetes namespaces for the current context cluster

  > ktrouble get ns
`,
	Run: func(cmd *cobra.Command, args []string) {

		if c.Client != nil {
			nssList := c.Client.GetNamespaces()

			nsData := objects.NamespaceList{}
			for _, v := range nssList.Items {
				nsData.Namespace = append(nsData.Namespace, v.Name)
			}

			c.OutputData(&nsData, objects.TextOptions{NoHeaders: c.NoHeaders})
		} else {
			common.Logger.Warn("Cannot fetch namespaces, no valid kubernetes context")
		}
	},
}

func init() {
	getCmd.AddCommand(namespaceCmd)
}
