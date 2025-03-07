package set

import (
	"fmt"

	"github.com/fatih/color"
)

type SetConfigCmd struct {
}

func (s *SetConfigCmd) Short() string {
	return "Set configuration options for ktrouble"
}

func (s *SetConfigCmd) Long() string {
	longText := ""
	yellow := color.New(color.FgYellow).SprintFunc()

	longText += `EXAMPLE:
  If you store your git personal access token in an ENV variable, you can specify
  the variable name.
`

	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set config --user christopher.maahs --tokenvar GLA_TOKEN`))

	longText += `EXAMPLE:
  If you don't store your personal access token in an ENV variable, it can be
  stored directly in the config.yaml file.  Don't forgot to add a 'space' in
  front of running this next command so the token doesn't end up in your
  history file, if you have that option set in your shell
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set config --user christopher.maahs --token <your token>`))

	longText += `EXAMPLE:
  If you want to point 'ktrouble' to a different git repository for upstream
  storage of utility pod definitions
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set config --giturl "https://github.com/cmaahs/ktrouble-utils.git"`))

	longText += `EXAMPLE:
  If you would like 'ktrouble launch' to prompt for secret selection on every
  run, rather than just when you use the '--secrets' switch
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set config --secrets`))

	longText += `EXAMPLE:
  If you would like 'ktrouble launch' to prompt for configmap selection on every
  run, rather than just when you use the '--configs' switch
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set config --configs`))

	longText += `EXAMPLE:
  If you use dynamic hyperlinking in your terminal software, you can enable output
  with '<bash: >' decorations
`
	longText = fmt.Sprintf("%s\n    > %s\n\n", longText, yellow(`ktrouble set config --bashlinks`))

	return longText

}
