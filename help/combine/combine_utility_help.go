package remove

import (
	"fmt"

	"github.com/fatih/color"
)

type CombineUtilityCmd struct {
}

func (r *CombineUtilityCmd) Short() string {
	return "Combine a utility from the config file"
}

func (r *CombineUtilityCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  The 'combine utility' allows you to collapse several utility definitions that
  share the same 'image' and differing 'tags' into a single definition.  This
  will create the new utility definition, adding the 'tags' from each to the
  'tags' list of the new definition.  It will also mark the utilities defined in
  the '--combine' parameter as "hidden", this way it can be removed from the
  upstream git repository with a 'ktrouble remove utility' with the
  '--remove-upstream' switch
`

	longText = fmt.Sprintf("%s\n    > %s\n", longText, yellow(`ktrouble combine utility --name 'mysql' --combine 'mysql5,mysql8'`))

	return longText
}
