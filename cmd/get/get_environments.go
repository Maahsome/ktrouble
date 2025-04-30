package get

import (
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// environmentsCmd represents the utilities command
var environmentsCmd = &cobra.Command{
	Use:     "environments",
	Aliases: defaults.EnvironmentAliases,
	Short:   getEnvironmentsHelp.Short(),
	Long:    getEnvironmentsHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {

		additionalFields := []string{}
		if c.ShowHidden {
			additionalFields = append(additionalFields, []string{"HIDDEN", "REMOVE_UPSTREAM"}...)
		}
		c.OutputData(&c.EnvDefs, objects.TextOptions{
			NoHeaders:        c.NoHeaders,
			ShowHidden:       c.ShowHidden,
			Fields:           c.Fields,
			AdditionalFields: additionalFields,
		})
	},
}

func init() {
	getCmd.AddCommand(environmentsCmd)
}
