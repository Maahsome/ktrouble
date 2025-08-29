package combine

import (
	"ktrouble/config"
	"ktrouble/defaults"
	help "ktrouble/help/combine"

	"github.com/spf13/cobra"
)

var combineHelp = help.CombineCmd{}
var combineUtilityHelp = help.CombineUtilityCmd{}

var combineCmd = &cobra.Command{
	Use:     "combine",
	Aliases: defaults.CombineAliases,
	Args:    cobra.MinimumNArgs(1),
	Short:   combineHelp.Short(),
	Long:    combineHelp.Long(),
	Run:     func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return combineCmd
}
