package help

import (
	"fmt"

	"github.com/fatih/color"
)

type GenHelpCmd struct {
}

func (g *GenHelpCmd) Short() string {
	return "Output help from all the sub commands"
}

func (g *GenHelpCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  This command will generate markdown for all of the cobra commands ktrouble
  supports.
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble genhelp > HELP.md`))

	longText += `EXAMPLE:
  This command will generate a wiki compatible file that can be submitted to
  confluence via the REST api.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble genhelp --format confluence > HELP.cf`))

	return longText
}
