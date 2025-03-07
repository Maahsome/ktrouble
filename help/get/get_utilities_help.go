package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetUtilitiesCmd struct {
}

func (g *GetUtilitiesCmd) Short() string {
	return "Get a list of supported utility container images"
}

func (g *GetUtilitiesCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Display a list of utilities defined in the configuration file
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get utilities`))

	return longText
}
