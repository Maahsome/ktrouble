package add

import (
	"fmt"

	"github.com/fatih/color"
)

type AddEnvironmentCmd struct {
}

func (a *AddEnvironmentCmd) Short() string {
	return "Add an environment definition to the ktrouble configuration"
}

func (a *AddEnvironmentCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Use 'add environment' command to add a new environment definition to your 'config.yaml'
  file`

	longText = fmt.Sprintf("%s\n\n    > %s\n\n", longText, yellow(`ktrouble add environment -e 'lowers' -r 'us-docker.pkg.dev/my-lower-repo'`))

	return longText
}
