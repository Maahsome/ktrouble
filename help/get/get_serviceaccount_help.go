package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetServiceAccountCmd struct {
}

func (g *GetServiceAccountCmd) Short() string {
	return "Get a list of K8s ServiceAccount(s) in a Namespace"
}

func (g *GetServiceAccountCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Return a list of kubernetes service accounts for a namespace
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble get serviceaccount -n myspace`))

	longText += `EXAMPLE:
  If you do not specify a namespace with '-n <namespace>', you will be prompted
  to select one
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble get sa`))

	return longText
}
