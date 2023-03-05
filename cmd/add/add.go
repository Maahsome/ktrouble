package add

import (
	"ktrouble/config"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MinimumNArgs(1),
	Short: "",
	Long: `EXAMPLES
	ktrouble add utility`,
	Run: func(cmd *cobra.Command, args []string) {},
}

var c *config.Config

func InitSubCommands(conf *config.Config) *cobra.Command {
	c = conf
	return addCmd
}
