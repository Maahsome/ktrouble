package get

import (
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// utilitiesCmd represents the utilities command
var utilitiesCmd = &cobra.Command{
	Use:     "utilities",
	Aliases: defaults.GetUtilitesAliases,
	Short:   "Get a list of supported utility container images",
	Long: `EXAMPLE:
  Display a list of utilities defined in the configuration file

  > ktrouble get utilities
`,
	Run: func(cmd *cobra.Command, args []string) {

		c.OutputData(&c.UtilDefs, objects.TextOptions{
			NoHeaders:  c.NoHeaders,
			ShowHidden: c.ShowHidden,
			Fields:     c.Fields,
		})

	},
}

func init() {
	getCmd.AddCommand(utilitiesCmd)
}
