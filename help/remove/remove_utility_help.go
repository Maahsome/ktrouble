package remove

import (
	"fmt"

	"github.com/fatih/color"
)

type RemoveUtilityCmd struct {
}

func (r *RemoveUtilityCmd) Short() string {
	return "Remove a utility from the config file, or HIDE it if it is an upstream definition"
}

func (r *RemoveUtilityCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The 'remove utility' command will prompt to select a utility definition to
  remove from your local 'config.yaml' file
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble remove utility`))

	return longText
}
