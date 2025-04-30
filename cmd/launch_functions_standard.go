package cmd

import (
	"fmt"
	"os"

	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/objects"
	"ktrouble/template"

	"github.com/muesli/termenv"
)

func standardLaunch(utility string, sa string) {

	termFormatter := termenv.NewOutput(os.Stdout)
	if c.Client != nil {
		utilMap := objects.GetUtilityMap(c.UtilDefs, c.EnvMap)

		if utility == "" {
			utility = ask.PromptForUtility(c.UtilDefs, c.EnvMap, c.ShowHidden)
		}

		// Display the HINT
		if len(utilMap[utility].Hint) > 0 {
			fmt.Println(utilMap[utility].Hint)
		}

		namespace := c.Client.DetermineNamespace(c.Namespace)
		if sa == "" {
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
		// utilMap[utility].RequireSecrets is from the utility definitions
		if p.PromptForSecrets || c.PromptForSecrets || utilMap[utility].RequireSecrets {
			secrets := c.Client.GetSecrets(namespace)
			selectedSecrets = ask.PromptForSecrets(secrets)
		}

		selectedConfigMaps := []string{}
		// p.PromptForConfigMaps is the local command param --configs
		// c.PromptForConfigMaps is the config.yaml promptForConfigMaps setting
		// utilMap[utility].RequireConfigmaps is from the utility definitions
		if p.PromptForConfigMaps || c.PromptForConfigMaps || utilMap[utility].RequireConfigmaps {
			configmaps := c.Client.GetConfigMaps(namespace)
			selectedConfigMaps = ask.PromptForConfigMaps(configmaps)
		}

		shortUniq := randSeq(c.UniqIdLength)
		osUser := os.Getenv("USER")
		tc := &template.TemplateConfig{
			Parameters: map[string]string{
				"name":           fmt.Sprintf("%s-%s", utilMap[utility].Name, shortUniq),
				"serviceAccount": sa,
				"namespace":      namespace,
				"registry":       utilMap[utility].Repository,
				"limitsCpu":      c.SizeMap[resourceSize].LimitsCPU,
				"limitsMem":      c.SizeMap[resourceSize].LimitsMEM,
				"requestCpu":     c.SizeMap[resourceSize].RequestCPU,
				"requestMem":     c.SizeMap[resourceSize].RequestMEM,
				"hasSelector":    hasSelector,
				"selector":       selector,
				"launchedby":     osUser,
				"ingressEnabled": fmt.Sprintf("%t", p.CreateIngress),
				"host":           p.Host,
				"targetPort":     fmt.Sprintf("%d", p.Port),
				"path":           p.Path,
				"associatedPod":  fmt.Sprintf("%s-%s", utility, shortUniq),
			},
			Secrets:    selectedSecrets,
			ConfigMaps: selectedConfigMaps,
		}

		common.Logger.Debugf("Template file: %s", c.TemplateFile)
		tp := template.New(c.TemplateFile)
		podManifest := tp.RenderTemplate(tc)
		common.Logger.Debugf("Manifest: \n%s\n", podManifest)
		c.Client.CreatePod(podManifest, namespace)

		if p.CreateIngress {
			common.Logger.Debugf("Service Template file: %s", c.ServiceTemplateFile)
			tps := template.New(c.ServiceTemplateFile)
			serviceManifest := tps.RenderTemplate(tc)
			common.Logger.Debugf("Manifest: \n%s\n", serviceManifest)
			c.Client.CreateService(serviceManifest, namespace)

			common.Logger.Debugf("Ingress Template file: %s", c.IngressTemplateFile)
			tpi := template.New(c.IngressTemplateFile)
			ingressManifest := tpi.RenderTemplate(tc)
			common.Logger.Debugf("Manifest: \n%s\n", ingressManifest)
			c.Client.CreateIngress(ingressManifest, namespace)
		}

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
}
