package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetEnvVarsCmd struct {
}

func (g *GetEnvVarsCmd) Short() string {
	return "Get a list of environment variables used by ktrouble"
}

func (g *GetEnvVarsCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Display a list of environment variables recognized by ktrouble
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get envvars`))

	return longText
}
