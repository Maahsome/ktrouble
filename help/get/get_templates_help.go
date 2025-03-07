package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetTemplatesCmd struct {
}

func (g *GetTemplatesCmd) Short() string {
	return "Get a list of templates"
}

func (g *GetTemplatesCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The 'get templates' command will output a list of templates in the templates
  directory
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get templates`))

	return longText
}
