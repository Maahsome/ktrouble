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

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update utility -u helm-kubectl --toggle-exclude`))

	longText += `EXAMPLE:
  Toggle the 'hidden' flag for an existing utility pod definition
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update utility -u alpine3 --toggle-hidden`))

	longText += `EXAMPLE:
  Change the 'command' the utility will run
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update utility -u helm-kubectl -c '/bin/sh'`))

	longText += `EXAMPLE:
  Set the image tags for a utility pod definition.  This is an overriding operation
  so make sure you specify all the tags you want to keep.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update utility -u helm-kubectl --tags "3.11,3.12"`))

	longText += `EXAMPLE:
  Change the image for a utility pod definition.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update utility -u redis -i "redis"`))

	return longText
}
