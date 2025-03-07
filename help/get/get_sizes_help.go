package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetSizesCmd struct {
}

func (g *GetSizesCmd) Short() string {
	return "Get a list of defined sizes"
}

func (g *GetSizesCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Display a list of POD size options from the configuration file
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get sizes`))

	return longText
}
