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
	Short:   getUtilitiesHelp.Short(),
	Long:    getUtilitiesHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		additionalFields := []string{}
		if c.ShowHidden {
			additionalFields = append(additionalFields, []string{"HIDDEN", "REMOVE_UPSTREAM"}...)
		}
		c.OutputData(&c.UtilDefs, objects.TextOptions{
			NoHeaders:        c.NoHeaders,
			ShowHidden:       c.ShowHidden,
			Fields:           c.Fields,
			AdditionalFields: additionalFields,
			DefaultFields:    c.OutputFieldsMap["utility"],
		})
	},
}

func init() {
	getCmd.AddCommand(utilitiesCmd)
}
