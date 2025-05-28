package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetOutputFieldsCmd struct {
}

func (g *GetOutputFieldsCmd) Short() string {
	return "Get a list fo the current output fields definitions"
}

func (g *GetOutputFieldsCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Display a list of output definition names and their related fields from the configuration file
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get output-fields`))

	return longText
}
