package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetNamespaceCmd struct {
}

func (g *GetNamespaceCmd) Short() string {
	return "Get a list of namespaces"
}

func (g *GetNamespaceCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Return a list of kubernetes namespaces for the current context cluster
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get ns`))

	return longText
}
