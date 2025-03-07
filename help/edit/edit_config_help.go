package edit

import (
	"fmt"

	"github.com/fatih/color"
)

type EditConfigCmd struct {
}

func (e *EditConfigCmd) Short() string {
	return "Edit the default config, or specified in KTROUBLE_CONFIG"
}

func (e *EditConfigCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The default config can be hand edited
  `

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble edit config`))

	longText += `EXAMPLE:
  Edit a secondary NON default config file
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`export KTROUBLE_CONFIG=cmaahs-config.yaml`))
	longText = fmt.Sprintf("%s    > %s\n\n", longText, yellow(`ktrouble edit config`))

	return longText
}
