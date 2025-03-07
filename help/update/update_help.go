package update

import (
	"fmt"

	"github.com/fatih/color"
)

type UpdateCmd struct {
}

func (u *UpdateCmd) Short() string {
	return "Update various objects for ktrouble"
}

func (u *UpdateCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:`
	longText = fmt.Sprintf("%s\n\n    > %s\n\n", longText, yellow(`ktrouble update utility --help`))

	return longText
}
