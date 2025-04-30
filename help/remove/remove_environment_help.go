package remove

import (
	"fmt"

	"github.com/fatih/color"
)

type RemoveEnvironmentCmd struct {
}

func (r *RemoveEnvironmentCmd) Short() string {
	return "Remove an environment from the config file, or HIDE it if it is an upstream definition, use --remove-upstream to remove it from the upstream repository"
}

func (r *RemoveEnvironmentCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The 'remove environment' command will prompt to select an environment definition to
  remove from your local 'config.yaml' file
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble remove environment`))

	return longText
}
