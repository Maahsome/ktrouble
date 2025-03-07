package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/objects"
	"ktrouble/template"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

type launchParam struct {
	PromptForSecrets    bool
	PromptForConfigMaps bool
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

		termFormatter := termenv.NewOutput(os.Stdout)
		if c.Client != nil {
			utilMap := make(map[string]objects.UtilityPod)
			for _, v := range c.UtilDefs {
				utilMap[v.Name] = objects.UtilityPod{
					Name:        v.Name,
					Repository:  v.Repository,
					ExecCommand: v.ExecCommand,
				}
			}

			utility := ""
			if len(args) > 0 && len(args[0]) > 0 {
				utility = args[0]
			} else {
				utility = ask.PromptForUtility(c.UtilDefs, c.ShowHidden)
			}
			namespace := c.Client.DetermineNamespace(c.Namespace)
			sa := "default"
			if len(args) > 1 && len(args[1]) > 0 {
				sa = args[1]
			} else {
				sasList := c.Client.GetServiceAccounts(namespace)
				sa = ask.PromptForServiceAccount(sasList)
			}

			resourceSize := ask.PromptForResourceSize(c.SizeDefs)

			nodeList := c.Client.GetNodes()
			selector := ask.PromptForNodeLabels(nodeList)
			hasSelector := "true"
			if selector == "\"-none-\"" {
				hasSelector = "false"
			}

			selectedSecrets := []string{}
			// p.PromptForSecrets is the local command param --secrets
			// c.PromptForSecrets is the config.yaml promptForSecrets setting
			if p.PromptForSecrets || c.PromptForSecrets {
				secrets := c.Client.GetSecrets(namespace)
				selectedSecrets = ask.PromptForSecrets(secrets)
			}

			selectedConfigMaps := []string{}
			// p.PromptForConfigMaps is the local command param --configs
			// c.PromptForConfigMaps is the config.yaml promptForConfigMaps setting
			if p.PromptForConfigMaps || c.PromptForConfigMaps {
				configmaps := c.Client.GetConfigMaps(namespace)
				selectedConfigMaps = ask.PromptForConfigMaps(configmaps)
			}

			shortUniq := randSeq(c.UniqIdLength)
			tc := &template.TemplateConfig{
				Parameters: map[string]string{
					"name":           fmt.Sprintf("%s-%s", utility, shortUniq),
					"serviceAccount": sa,
					"namespace":      namespace,
					"registry":       utilMap[utility].Repository,
					"limitsCpu":      c.SizeMap[resourceSize].LimitsCPU,
					"limitsMem":      c.SizeMap[resourceSize].LimitsMEM,
					"requestCpu":     c.SizeMap[resourceSize].RequestCPU,
					"requestMem":     c.SizeMap[resourceSize].RequestMEM,
					"hasSelector":    hasSelector,
					"selector":       selector,
				},
				Secrets:    selectedSecrets,
				ConfigMaps: selectedConfigMaps,
			}

			common.Logger.Debugf("Template file: %s", c.TemplateFile)
			tp := template.New(c.TemplateFile)
			podManifest := tp.RenderTemplate(tc)
			common.Logger.Debugf("Manifest: \n%s\n", podManifest)
			c.Client.CreatePod(podManifest, namespace)

			if c.EnableBashLinks {
				hl := fmt.Sprintf("<bash:kubectl -n %s exec -it %s -- %s>", namespace, fmt.Sprintf("%s-%s", utility, shortUniq), utilMap[utility].ExecCommand)
				tx := fmt.Sprintf("kubectl -n %s exec -it %s -- %s", namespace, fmt.Sprintf("%s-%s", utility, shortUniq), utilMap[utility].ExecCommand)
				fmt.Println(termFormatter.Hyperlink(hl, tx))
			} else {
				fmt.Printf("kubectl -n %s exec -it %s -- %s\n", namespace, fmt.Sprintf("%s-%s", utility, shortUniq), utilMap[utility].ExecCommand)
			}
		} else {
			common.Logger.Warn("Cannot launch a pod, no valid kubernetes context")
		}

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
}
