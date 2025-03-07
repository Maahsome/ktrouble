package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetNodeCmd struct {
}

func (g *GetNodeCmd) Short() string {
	return "Get a list of node labels"
}

func (g *GetNodeCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Get a list of nodes for the current context cluster
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get node`))

	return longText
}
