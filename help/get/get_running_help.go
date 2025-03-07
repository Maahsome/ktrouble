package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetRunningCmd struct {
}

func (g *GetRunningCmd) Short() string {
	return "Get a list of running pods"
}

func (g *GetRunningCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Get a list of PODs that are currently running on the current context kubernetes
  cluster that were created with the ktrouble utility.  If the 'enableBashLinks'
  config.yaml setting is 'true', a '<bash: ... >' command will be displayed,
  otherwise the SHELL path will be displayed.
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get running`))

	longText += `
      NAME                NAMESPACE       STATUS   EXEC
      basic-tools-e1df2f  common-tooling  Running  <bash:kubectl -n common-tooling exec -it basic-tools-e1df2f -- /bin/bash>

      NAME                NAMESPACE       STATUS   SHELL
      basic-tools-e1df2f  common-tooling  Running  /bin/bash

`

	longText += `EXAMPLE:
  You can use the subcommand 'pods' in place of 'running'
`
	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get pods`))

	return longText
}
