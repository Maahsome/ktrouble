package get

import (
	"fmt"
	"ktrouble/common"
	"ktrouble/defaults"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// configsCmd represents the templates command
var configsCmd = &cobra.Command{
	Use:     "configs",
	Aliases: defaults.GetSizesAliases,
	Short:   "Get a list of configs",
	Long: `EXAMPLE:
  > ktrouble get configs
`,
	Run: func(cmd *cobra.Command, args []string) {
		home, herr := os.UserHomeDir()
		if herr != nil {
			common.Logger.WithError(herr).Error("failed to determine the HOME directory")
		}
		confDir := fmt.Sprintf("%s/.config/ktrouble", home)
		files, _ := os.ReadDir(confDir)
		if !c.NoHeaders {
			fmt.Println("CONFIG")
		}
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".yaml") {
				fmt.Printf("%s\n", f.Name())
			}
		}
	},
}

func init() {
	getCmd.AddCommand(configsCmd)
}
