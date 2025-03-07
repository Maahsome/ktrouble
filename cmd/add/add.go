package add

import (
	"ktrouble/config"
	help "ktrouble/help/add"

	"github.com/spf13/cobra"
)

var addHelp = help.AddCmd{}
var addUtilityHelp = help.AddUtilityCmd{}

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MinimumNArgs(1),
	Short: addHelp.Short(),
	Long:  addHelp.Long(),
	Run:   func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return addCmd
}
