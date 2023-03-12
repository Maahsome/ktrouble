package update

import (
	"ktrouble/config"
	"ktrouble/defaults"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: defaults.UpdateAliases,
	Args:    cobra.MinimumNArgs(1),
	Short:   "",
	Long: `EXAMPLES
	ktrouble update utility`,
	Run: func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return updateCmd
}
