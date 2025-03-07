package help

import (
	"fmt"

	"github.com/fatih/color"
)

type LaunchCmd struct {
}

func (l *LaunchCmd) Short() string {
	return "Launch a kubernetes troubleshooting pod"
}

func (l *LaunchCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Just running ktrouble launch will prompt for all the things required to run
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble launch`))

	longText += `EXAMPLE:
  As an option, ktrouble can prompt with a list of kubernetes secrets from the
  namespace selected.  The chosen secrets will all be mounted under '/secrets/'
  directory, with each key as a file inside a directory named for the secret.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble launch --secrets`))

	longText += `EXAMPLE:
  TODO: add command line parameters that can be used to set all the options
  for launching a POD
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble launch (...)`))

	return longText
}
