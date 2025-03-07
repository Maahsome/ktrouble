package update

import (
	"fmt"

	"github.com/fatih/color"
)

type UpdateUtilityCmd struct {
}

func (u *UpdateUtilityCmd) Short() string {
	return "Update an existing utility pod definition in the local config.yaml file"
}

func (u *UpdateUtilityCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Toggle the 'exclude from push' flag for a utility definition.
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update utility -u helm-kubectl311 --toggle-exclude`))

	longText += `EXAMPLE:
  Toggle the 'hidden' flag for an existing utility pod definition
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update utility -u alpine3 --toggle-hidden`))

	longText += `EXAMPLE:
  Change the 'command' the utility will run
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update utility -u helm-kubectl311 -c '/bin/sh'`))

	return longText
}
