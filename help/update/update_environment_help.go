package update

import (
	"fmt"

	"github.com/fatih/color"
)

type UpdateEnvironmentCmd struct {
}

func (u *UpdateEnvironmentCmd) Short() string {
	return "Update an existing environment definition in the local config.yaml file"
}

func (u *UpdateEnvironmentCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Toggle the 'exclude from push' flag for an environment definition.
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update environment -e lowers --toggle-exclude`))

	longText += `EXAMPLE:
  Change the 'repository' for the definition 'lowers' to '8675309.dkr.ecr.us-west-2.amazonaws.com'
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update environment -e lower -r '8675309.dkr.ecr.us-west-2.amazonaws.com'`))

	return longText
}
