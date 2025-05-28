package get

import (
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// outputFieldsCmd
var outputFieldsCmd = &cobra.Command{
	Use:     "output-fields",
	Aliases: defaults.OutputFieldsAliases,
	Short:   getOutputFieldsHelp.Short(),
	Long:    getOutputFieldsHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {

		c.OutputData(&c.OutputFieldsDefs, objects.TextOptions{
			NoHeaders:     c.NoHeaders,
			Fields:        c.Fields,
			DefaultFields: c.OutputFieldsMap["output_fields"],
		})
	},
}

func init() {
	getCmd.AddCommand(outputFieldsCmd)
}
