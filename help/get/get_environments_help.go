package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetEnvironmentsCmd struct {
}

func (g *GetEnvironmentsCmd) Short() string {
	return "Get a list of defined environments"
}

func (g *GetEnvironmentsCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Display a list of environments defined in the configuration file
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get environments`))

	return longText
}
