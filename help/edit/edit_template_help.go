package edit

import (
	"fmt"

	"github.com/fatih/color"
)

type EditTemplateCmd struct {
}

func (e *EditTemplateCmd) Short() string {
	return "Edit the default template, or specified one via --template/-t"
}

func (e *EditTemplateCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  `

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble edit template --template christmas`))

	return longText
}
