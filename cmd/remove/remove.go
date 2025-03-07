package remove

import (
	"ktrouble/config"
	help "ktrouble/help/remove"

	"github.com/spf13/cobra"
)

var removeHelp = help.RemoveCmd{}
var removeUtilityHelp = help.RemoveUtilityCmd{}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Args:  cobra.MinimumNArgs(1),
	Short: removeHelp.Short(),
	Long:  removeHelp.Long(),
	Run:   func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return removeCmd
}
