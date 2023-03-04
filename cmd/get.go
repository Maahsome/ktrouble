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
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("get called")
	// },
}

func init() {
	rootCmd.AddCommand(getCmd)
}
