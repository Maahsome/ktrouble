package help

import (
	"fmt"

	"github.com/fatih/color"
)

type PullCmd struct {
}

func (p *PullCmd) Short() string {
	return "Pull utility definitions from git"
}

func (p *PullCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The 'pull' command will prompt to choose from a list of utilities that are
  missing from your local 'config.yaml' utility defintions.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble pull`))

	longText += `EXAMPLE:
  Items that you have that are local, but have different setting, can be pulled,
  and overwritten by adding a '-a' switch to the command.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble pull -a`))
	return longText
}
