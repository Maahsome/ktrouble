package add

import (
	"fmt"

	"github.com/fatih/color"
)

type AddCmd struct {
}

func (a *AddCmd) Short() string {
	return "Add various objects for ktrouble"
}

func (a *AddCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Use the "add utility" command to add a new utility definition to your 'config.yaml'`

	longText = fmt.Sprintf("%s\n\n    > %s\n\n", longText, yellow(`ktrouble add utility --help`))

	return longText
}
