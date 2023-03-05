package set

import (
	"ktrouble/config"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.MinimumNArgs(1),
	Short: "",
	Long: `EXAMPLES
	ktrouble set gituser`,
	Run: func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return setCmd
}
