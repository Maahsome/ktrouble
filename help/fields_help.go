package help

import (
	"fmt"

	"github.com/fatih/color"
)

type FieldsCmd struct {
}

func (f *FieldsCmd) Short() string {
	return "Display a list of valid fields to use with the --fields/-f parameter for each command"
}

func (f *FieldsCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The 'fields' command will list the valid fields that can be used with various
  commands that accept the --fields/-f parameter.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble fields`))

	return longText
}
