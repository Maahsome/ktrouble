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
	// TODO: finish parames
	Utility        string
	ServiceAccount string
	Namespace      string
	NodeSelector   string
	Secrets        []string
	ConfigMaps     []string
	Size           string
	OutputName     bool
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
		standardLaunch(p)
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
	launchCmd.Flags().BoolVar(&p.PromptForSecrets, "prompt-secrets", false, "Use this switch to prompt to mount secrets in the POD")
	launchCmd.Flags().BoolVar(&p.PromptForConfigMaps, "prompt-configmaps", false, "Use this switch to prompt to mount configmaps in the POD")
	launchCmd.Flags().BoolVar(&p.PromptForVolumes, "volumes", false, "Use this switch to prompt to mount volumes in the POD")
	launchCmd.Flags().BoolVar(&p.CreateIngress, "ingress", false, "Use this switch to enable creating a service and ingress for the POD")
	launchCmd.Flags().IntVar(&p.Port, "port", 8080, "Specify the port that the POD listens on, used in the service and ingress settings")
	launchCmd.Flags().StringVar(&p.Host, "host", "flexo.bender.rocks", "Specify the host that the ingress will listen on, for configuration of ingress-nginx")
	launchCmd.Flags().StringVar(&p.Path, "path", "service-futurama", "Specify the PATH that the ingress will listen on, for configuration of ingress-nginx, sans the enclosing slashes")
	launchCmd.Flags().StringVarP(&p.Utility, "utility", "u", "", "Specify the name of the utility to launch")
	launchCmd.Flags().StringVar(&p.ServiceAccount, "service-account", "", "Specify the name of the service account to use")
	launchCmd.Flags().StringVarP(&p.Namespace, "namespace", "n", "", "Specify the namespace to use")
	launchCmd.Flags().StringVar(&p.NodeSelector, "node-selector", "", "Specify the node selector to use")
	launchCmd.Flags().StringSliceVar(&p.Secrets, "secrets", []string{}, "Specify an array of secret names to mount, eg --secrets 'secret1,secret2'")
	launchCmd.Flags().StringSliceVar(&p.ConfigMaps, "configmaps", []string{}, "Specify an array of configmap names to mount, eg --secrets 'cm1,cm2'")
	launchCmd.Flags().StringVar(&p.Size, "size", "", "Specify the size of the POD, eg --size 'small,medium,large'")
	launchCmd.Flags().BoolVar(&p.OutputName, "output-name", false, "Use this switch to only output the name of the POD")
}
