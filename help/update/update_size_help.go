package update

import (
	"fmt"

	"github.com/fatih/color"
)

type UpdateSizeCmd struct {
}

func (u *UpdateSizeCmd) Short() string {
	return "Update an existing resource size definition in the local config.yaml file"
}

func (u *UpdateSizeCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Update the cpu/memory limits for an existing size definition.
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update size --name Small --limitscpu 300m --limitsmem 3Gi`))

	longText += `EXAMPLE:
  If '--name' is omitted, you will be prompted to choose a size definition to update.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble update size --requestcpu 150m --requestmem 768Mi`))

	return longText
}
