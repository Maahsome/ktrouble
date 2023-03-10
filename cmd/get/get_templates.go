package get

import (
	"fmt"
	"ktrouble/common"
	"ktrouble/defaults"
	"os"

	"github.com/spf13/cobra"
)

// templatesCmd represents the templates command
var templatesCmd = &cobra.Command{
	Use:     "templates",
	Aliases: defaults.GetSizesAliases,
	Short:   "Get a list of templates",
	Long: `EXAMPLE:
  > ktrouble get templates
`,
	Run: func(cmd *cobra.Command, args []string) {
		home, herr := os.UserHomeDir()
		if herr != nil {
			common.Logger.WithError(herr).Error("failed to determine the HOME directory")
		}
		tmplDir := fmt.Sprintf("%s/.config/ktrouble/templates", home)
		files, _ := os.ReadDir(tmplDir)
		if !c.NoHeaders {
			fmt.Println("TEMPLATE")
		}
		for _, f := range files {
			fmt.Printf("%s\n", f.Name())
		}
	},
}

func init() {
	getCmd.AddCommand(templatesCmd)
}
