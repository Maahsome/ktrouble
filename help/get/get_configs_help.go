package get

import (
	"fmt"

	"github.com/fatih/color"
)

type GetConfigsCmd struct {
}

func (g *GetConfigsCmd) Short() string {
	return "Get a list of configs"
}

func (g *GetConfigsCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The ktrouble utility can support multiple config files, either with the
  '--config <config path>' or by setting the environment variable
  'KTROUBLE_CONFIG' to just the name of the config file, which will need to
  reside in the '$HOME/.config/ktrouble' directory
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble get configs`))

	longText += `
      CONFIG
      alteryx-config.yaml
      cmaahs-config.yaml
      config.yaml
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`export KTROUBLE_CONFIG=cmaahs-config.yaml`))

	return longText
}
