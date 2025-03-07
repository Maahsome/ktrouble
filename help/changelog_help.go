package help

import (
	"fmt"

	"github.com/fatih/color"
)

type ChangelogCmd struct {
}

func (c *ChangelogCmd) Short() string {
	return "Get changelog information"
}

func (c *ChangelogCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Get just the latest changelog entry
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble changelog`))

	longText += `EXAMPLE:
  Get all the changelog entries
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble changelog --all`))

	return longText
}
