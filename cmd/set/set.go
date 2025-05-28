package set

import (
	"ktrouble/config"
	help "ktrouble/help/set"

	"github.com/spf13/cobra"
)

var setHelp = help.SetCmd{}
var setConfigHelp = help.SetConfigCmd{}
var setOutputFieldsHelp = help.SetOutputFieldsCmd{}

var setCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.MinimumNArgs(1),
	Short: setHelp.Short(),
	Long:  setHelp.Long(),
	Run:   func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return setCmd
}
