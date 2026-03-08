package add

import (
	"fmt"

	"github.com/fatih/color"
)

type AddSizeCmd struct {
}

func (a *AddSizeCmd) Short() string {
	return "Add a resource size definition to the ktrouble configuration"
}

func (a *AddSizeCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  Use 'add size' command to add a new resource sizing definition to your 'config.yaml'
  file`

	longText = fmt.Sprintf("%s\n\n    > %s\n\n", longText, yellow(`ktrouble add size --name small-plus --limitscpu 300m --limitsmem 3Gi --requestcpu 150m --requestmem 768Mi`))

	return longText
}
