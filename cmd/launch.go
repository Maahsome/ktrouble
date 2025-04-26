package cmd

import (
	"ktrouble/defaults"
	"math/rand"

	"github.com/spf13/cobra"
)

type launchParam struct {
	PromptForSecrets    bool
	PromptForConfigMaps bool
	PromptForVolumes    bool
	CreateIngress       bool
	Port                int
	Host                string
	Path                string
}

var p launchParam

var letters = []rune("abcdef0987654321")

// launchCmd represents the default command
var launchCmd = &cobra.Command{
	Use:     "launch",
	Aliases: defaults.LaunchAliases,
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
	launchCmd.Flags().BoolVar(&p.CreateIngress, "ingress", false, "Use this switch to enable creating a service and ingress for the POD")
	launchCmd.Flags().IntVar(&p.Port, "port", 8080, "Specify the port that the POD listens on, used in the service and ingress settings")
	launchCmd.Flags().StringVar(&p.Host, "host", "flexo.bender.rocks", "Specify the host that the ingress will listen on, for configuration of ingress-nginx")
	launchCmd.Flags().StringVar(&p.Path, "path", "service-futurama", "Specify the PATH that the ingress will listen on, for configuration of ingress-nginx, sans the enclosing slashes")
}
