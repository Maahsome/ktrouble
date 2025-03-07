package set

import (
	"fmt"

	"github.com/fatih/color"
)

type SetCmd struct {
}

func (s *SetCmd) Short() string {
	return "Set various objects for ktrouble"
}

func (s *SetCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set config --help`))

	return longText
}
