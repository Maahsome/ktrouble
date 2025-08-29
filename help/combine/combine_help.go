package remove

import (
	"fmt"

	"github.com/fatih/color"
)

type CombineCmd struct {
}

func (r *CombineCmd) Short() string {
	return "Combine utilty definitions for ktrouble"
}

func (r *CombineCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble combine utility --help`))

	return longText
}
