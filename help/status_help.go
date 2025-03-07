package help

import (
	"fmt"

	"github.com/fatih/color"
)

type StatusCmd struct {
}

func (s *StatusCmd) Short() string {
	return "Get a comparison of the local utility definitions with the upstream one"
}

func (s *StatusCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The 'status' command will list the disposition of your local 'config.yaml'
  file 'utilities' definitions against the 'futurama/farnsworth/tools/ktrouble-utils'
  repostory.
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble status`))

	return longText
}
