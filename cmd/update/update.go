package update

import (
	"ktrouble/config"
	"ktrouble/defaults"
	help "ktrouble/help/update"

	"github.com/spf13/cobra"
)

var updateHelp = help.UpdateCmd{}
var updateUtilityHelp = help.UpdateUtilityCmd{}

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: defaults.UpdateAliases,
	Args:    cobra.MinimumNArgs(1),
	Short:   updateHelp.Short(),
	Long:    updateHelp.Long(),
	Run:     func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return updateCmd
}
