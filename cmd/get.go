package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get various resource lists",
	Long: `EXAMPLE:
`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
