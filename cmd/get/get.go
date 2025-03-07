package get

import (
	"ktrouble/config"
	help "ktrouble/help/get"

	"github.com/spf13/cobra"
)

var getHelp = help.GetCmd{}
var getConfigsHelp = help.GetConfigsCmd{}
var getRunningHelp = help.GetRunningCmd{}
var getTemplatesHelp = help.GetTemplatesCmd{}
var getNamespaceHelp = help.GetNamespaceCmd{}
var getNodeHelp = help.GetNodeCmd{}
var getNodeLabelsHelp = help.GetNodeLabelsCmd{}
var getServiceAccountHelp = help.GetServiceAccountCmd{}
var getSizesHelp = help.GetSizesCmd{}
var getUtilitiesHelp = help.GetUtilitiesCmd{}

var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.MinimumNArgs(1),
	Short: getHelp.Short(),
	Long:  getHelp.Long(),
	Run:   func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return getCmd
}
