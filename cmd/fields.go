package cmd

import (
	"fmt"
	"ktrouble/defaults"

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

	validFields := defaults.ValidOutputFields()
	help := `A list of valid FIELDS that can be specified by command:

  COMMAND: get|add|update|remove environment
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["environment"])
	help += `

  COMMAND: get ingress
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["ingress"])
	help += `

  COMMAND: get|add|update|remove sleep
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["ephemeral_sleep"])
	help += `

  COMMAND: get|add|update|remove utility
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["utility"])
	help += `

  COMMAND: get namespace
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["namespace"])
	help += `

  COMMAND: get nodes
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["node_labels"])
	help += `

  COMMAND: get running
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["pod"])
	help += `

  COMMAND: get service
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["service"])
	help += `

  COMMAND: get serviceaccount
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["service_account"])
	help += `

  COMMAND: get sizes
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["size"])
	help += `

  COMMAND: status
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["status"])
	help += `

  COMMAND version
`
	help += fmt.Sprintf("      FIELDS: %s", validFields["version"])

	fmt.Println(help)
}

func init() {
	RootCmd.AddCommand(fieldsCmd)
}
