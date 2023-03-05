package remove

import (
	"ktrouble/config"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Args:  cobra.MinimumNArgs(1),
	Short: "",
	Long: `EXAMPLES
	ktrouble remove utility`,
	Run: func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return removeCmd
}
