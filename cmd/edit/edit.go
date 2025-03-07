package edit

import (
	"ktrouble/config"
	help "ktrouble/help/edit"

	"github.com/spf13/cobra"
)

const (
	chunksize int = 1024
)

var editHelp = help.EditCmd{}
var editConfigHelp = help.EditConfigCmd{}
var editTemplateHelp = help.EditTemplateCmd{}

var editCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.MinimumNArgs(1),
	Short: editHelp.Short(),
	Long: editHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return editCmd
}
