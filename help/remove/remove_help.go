package remove

import (
	"fmt"

	"github.com/fatih/color"
)

type RemoveCmd struct {
}

func (r *RemoveCmd) Short() string {
	return "Remove various objects for ktrouble"
}

func (r *RemoveCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble remove utility --help`))

	return longText
}
