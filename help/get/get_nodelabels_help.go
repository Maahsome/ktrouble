package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetNodeLabelsCmd struct {
}

func (g *GetNodeLabelsCmd) Short() string {
	return "Get a list of defined node labels in config.yaml"
}

func (g *GetNodeLabelsCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Show the list of node labels in the configuration file
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get nodelabels`))

	return longText
}
