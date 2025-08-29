package help

import (
	"fmt"

	"github.com/fatih/color"
)

type PushCmd struct {
}

func (p *PushCmd) Short() string {
	return "Push local objects to upstream git repository"
}

func (p *PushCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The 'push' command allows you to push your local utility definitions into a
  common repository in the repository defined in the config file.  The command
  will prompt you to choose a list of utilities to push to the repository.
  Utilities marked 'exclude from push' will not appear on the selection list.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble push`))

	return longText
}
