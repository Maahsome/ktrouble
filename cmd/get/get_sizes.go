package get

import (
	"ktrouble/defaults"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// sizesCmd represents the sizes command
var sizesCmd = &cobra.Command{
	Use:     "sizes",
	Aliases: defaults.GetSizesAliases,
	Short:   getSizesHelp.Short(),
	Long:    getSizesHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {

		c.OutputData(&c.SizeDefs, objects.TextOptions{NoHeaders: c.NoHeaders})

	},
}

func init() {
	getCmd.AddCommand(sizesCmd)
}
