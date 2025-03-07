package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetCmd struct {
}

func (g *GetCmd) Short() string {
	return "Get various internal configuration and kubernetes resource listings"
}

func (g *GetCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  These are mostly utility commands to review things important to running ktrouble.
  Allowing a display of various items stored in config.yaml and listing various
  kubernetes resources.
`
	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get configs --help`))
	longText = fmt.Sprintf("%s    > %s\n", longText, yellow(`ktrouble get namespaces --help`))
	longText = fmt.Sprintf("%s    > %s\n", longText, yellow(`ktrouble get node --help`))
	longText = fmt.Sprintf("%s    > %s\n", longText, yellow(`ktrouble get nodelabels --help`))
	longText = fmt.Sprintf("%s    > %s\n", longText, yellow(`ktrouble get running --help`))
	longText = fmt.Sprintf("%s    > %s\n", longText, yellow(`ktrouble get serviceaccount --help`))
	longText = fmt.Sprintf("%s    > %s\n", longText, yellow(`ktrouble get sizes --help`))
	longText = fmt.Sprintf("%s    > %s\n", longText, yellow(`ktrouble get templates --help`))
	longText = fmt.Sprintf("%s    > %s\n\n", longText, yellow(`ktrouble get utilities --help`))

	longText += `EXAMPLE:
  Get a list of PODs that are currently running on the current context kubernetes
  cluster that were created with the ktrouble utility.  If the 'enableBashLinks'
  config.yaml setting is 'true', a '<bash: ... >' command will be displayed,
  otherwise the SHELL path will be displayed.
`
	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get pods`))

	longText += `
      NAME                NAMESPACE       STATUS   EXEC
      basic-tools-e1df2f  common-tooling  Running  <bash:kubectl -n common-tooling exec -it basic-tools-e1df2f -- /bin/bash>

      NAME                NAMESPACE       STATUS   SHELL
      basic-tools-e1df2f  common-tooling  Running  /bin/bash
`

	return longText
}
