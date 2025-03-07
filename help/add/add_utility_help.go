package add

import (
	"fmt"

	"github.com/fatih/color"
)

type AddUtilityCmd struct {
}

func (a *AddUtilityCmd) Short() string {
	return "Add a utility definition to the ktrouble configuration"
}

func (a *AddUtilityCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Use 'add utility' command to add a new utility definition to your 'config.yaml'
  file`

	longText = fmt.Sprintf("%s\n\n    > %s\n\n", longText, yellow(`ktrouble add utility -u helm-kubectl311 -c "/bin/bash" -r "dtzar/helm-kubectl:3.11"`))

	return longText
}
