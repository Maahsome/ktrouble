package get

import (
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// utilitiesCmd represents the utilities command
var utilitiesCmd = &cobra.Command{
	Use:     "utilities",
	Aliases: []string{"utility", "utils", "util", "container", "containers", "image", "images"},
	Short:   "Get a list of supported utility container images",
	Long: `EXAMPLE:
  Display a list of utilities defined in the configuration file

  > ktrouble get utilities
`,
	Run: func(cmd *cobra.Command, args []string) {

		c.OutputData(&c.UtilDefs, objects.TextOptions{
			NoHeaders:    c.NoHeaders,
			ShowExec:     c.EnableBashLinks,
			UtilMap:      c.UtilMap,
			UniqIdLength: c.UniqIdLength,
			ShowHidden:   c.ShowHidden,
		})

	},
}

func init() {
	getCmd.AddCommand(utilitiesCmd)
}
