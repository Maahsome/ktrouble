package get

import (
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// sleepCmd represents the sizes command
var sleepCmd = &cobra.Command{
	Use:     "sleep",
	Aliases: defaults.GetSleepAliases,
	Short:   "Get a list of sleep times for ephemeral containers",
	Long: `EXAMPLE:
  Display a list of sleep times for ephemeral containers

  > ktrouble get sleep
`,
	Run: func(cmd *cobra.Command, args []string) {

		c.OutputData(&c.EphemeralSleepDefs, objects.TextOptions{
			NoHeaders:  c.NoHeaders,
			ShowHidden: c.ShowHidden,
			Fields:     c.Fields,
		})
	},
}

func init() {
	getCmd.AddCommand(sleepCmd)
}
