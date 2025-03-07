package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fieldsCmd represents the fields command
var fieldsCmd = &cobra.Command{
	Use:   "fields",
	Short: fieldsHelp.Short(),
	Long:  fieldsHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		displayFieldHelp()
	},
}

func displayFieldHelp() {

	help := `A list of valid FIELDS that can be specified by command:

  COMMAND: get|add|update|remove utility

      FIELDS: NAME, REPOSITORY, EXEC, HIDDEN, EXCLUDED, SOURCE
`

	fmt.Println(help)
}

func init() {
	RootCmd.AddCommand(fieldsCmd)
}
