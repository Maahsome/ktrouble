package cmd

import (
	"math/rand"

	"github.com/spf13/cobra"
)

type launchParam struct {
	PromptForSecrets    bool
	PromptForConfigMaps bool
	PromptForVolumes    bool
}

var p launchParam

var letters = []rune("abcdef0987654321")

// launchCmd represents the default command
var launchCmd = &cobra.Command{
	Use:     "launch",
	Aliases: []string{"create", "apply", "pod", "l"},
	Short:   launchHelp.Short(),
	Long:    launchHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		utility := ""
		sa := ""
		if len(args) > 0 && len(args[0]) > 0 {
			utility = args[0]
		}
		if len(args) > 1 && len(args[1]) > 0 {
			sa = args[1]
		}

		standardLaunch(utility, sa)
	},
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func init() {
	RootCmd.AddCommand(launchCmd)
	launchCmd.Flags().BoolVar(&p.PromptForSecrets, "secrets", false, "Use this switch to prompt to mount secrets in the POD")
	launchCmd.Flags().BoolVar(&p.PromptForConfigMaps, "configs", false, "Use this switch to prompt to mount configmaps in the POD")
	launchCmd.Flags().BoolVar(&p.PromptForVolumes, "volumes", false, "Use this switch to prompt to mount volumes in the POD")
}
