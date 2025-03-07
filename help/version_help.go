package help

import (
	"fmt"

	"github.com/fatih/color"
)

type VersionCmd struct {
}

func (v *VersionCmd) Short() string {
	return "Express the 'version' of ktrouble"
}

func (v *VersionCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE: `
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble version`))

	return longText
}
