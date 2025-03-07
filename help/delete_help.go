package help

import (
	"fmt"

	"github.com/fatih/color"
)

type DeleteCmd struct {
}

func (d *DeleteCmd) Short() string {
	return "Delete PODs that have been created by ktrouble"
}

func (d *DeleteCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Delete a running POD.  This will prompt with a list of PODs that are running
  and were launched using ktrouble.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble delete`))

	return longText
}
