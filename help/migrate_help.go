package help

import (
	"fmt"

	"github.com/fatih/color"
)

type MigrateCmd struct {
}

func (f *MigrateCmd) Short() string {
	return "Migrate a repostiory to the current version"
}

func (f *MigrateCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  ktrouble uses a versioned directory structure to store the data.  This command
  will create a new version directory in the repository.  The existing data will
  remain in the old version directory.  The new version directory will be
  initialized with the current version data.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble migrate`))

	return longText
}
