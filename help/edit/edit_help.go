package edit

import (
	"fmt"

	"github.com/fatih/color"
)

type EditCmd struct {
}

func (e *EditCmd) Short() string {
	return "Edit various objects for ktrouble"
}

func (e *EditCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
`
	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble edit config --help`))
	longText = fmt.Sprintf("%s    > %s\n\n", longText, yellow(`ktrouble edit template --help`))

	return longText
}
