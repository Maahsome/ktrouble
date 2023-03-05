package cmd

import (
	"bytes"
	"fmt"
	"math/rand"

	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/objects"
	"ktrouble/template"

	"github.com/spf13/cobra"
)

type TemplateConfig struct {
	Parameters map[string]string
}

var letters = []rune("abcdef0987654321")

// launchCmd represents the default command
var launchCmd = &cobra.Command{
	Use:     "launch",
	Aliases: []string{"create", "apply", "pod", "l"},
	Short:   "launch a kubernetes troubleshooting pod",
	Long: `EXAMPLE:
  Just running kubectl launch will prompt for all the things required to run

  > kubectl launch

EXAMPLE:
  TODO: add command line parameters that can be used to set all the options
  for launching a POD

  > kubectl launch (...)
`,
	Run: func(cmd *cobra.Command, args []string) {

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
			utility = ask.PromptForUtility(c.UtilDefs)
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
		shortUniq := randSeq(c.UniqIdLength)
		tc := &TemplateConfig{
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
		}

		var tpl bytes.Buffer
		if err := template.ApplicationsTemplate.Execute(&tpl, tc); err != nil {
			common.Logger.WithError(err).Error("unable to generate the template data")
		}

		podManifest := tpl.String()

		c.Client.CreatePod(podManifest, namespace)

		fmt.Printf("kubectl -n %s exec -it %s -- %s\n", namespace, fmt.Sprintf("%s-%s", utility, shortUniq), utilMap[utility].ExecCommand)

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
}
