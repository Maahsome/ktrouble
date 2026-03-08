package remove

import (
	"fmt"

	"github.com/fatih/color"
)

type RemoveSizeCmd struct {
}

func (r *RemoveSizeCmd) Short() string {
	return "Remove a resource size from the config file"
}

func (r *RemoveSizeCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Remove a size definition from your local 'config.yaml' file.
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble remove size --name Small`))

	longText += `EXAMPLE:
  If '--name' is omitted, you will be prompted to choose a size definition to remove.
`
	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble remove size`))

	return longText
}
