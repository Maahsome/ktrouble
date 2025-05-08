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

func standardLaunch(p launchParam) {

	utility := p.Utility
	namespace := p.Namespace
	sa := p.ServiceAccount
	resourceSize := p.Size
	selector := p.NodeSelector
	hasSelector := "false"

	termFormatter := termenv.NewOutput(os.Stdout)
	if c.Client != nil {
		utilMap := objects.GetUtilityMap(c.UtilDefs, c.EnvMap)

		if p.Utility == "" {
			utility = ask.PromptForUtility(c.UtilDefs, c.EnvMap, c.ShowHidden)
			if len(utilMap[utility].Hint) > 0 {
				fmt.Println(utilMap[utility].Hint)
			}
		} else {
			if _, ok := utilMap[p.Utility]; !ok {
				fmt.Printf("Invalid utility: %s\n", p.Utility)
				return
			}
			utility = p.Utility
		}

		if p.Namespace == "" {
			namespace = c.Client.DetermineNamespace(c.Namespace)
		} else {
			if !c.Client.IsNamespaceValid(namespace) {
				fmt.Printf("Invalid namespace: %s\n", namespace)
				return
			}
		}

		if p.ServiceAccount == "" {
			sasList := c.Client.GetServiceAccounts(namespace)
			sa = ask.PromptForServiceAccount(sasList)
		} else {
			if !c.Client.IsServiceAccountValid(namespace, sa) {
				fmt.Printf("Invalid service account: %s\n", sa)
				return
			}
		}

		if p.Size == "" {
			resourceSize = ask.PromptForResourceSize(c.SizeDefs)
		} else {
			if _, ok := c.SizeMap[p.Size]; !ok {
				fmt.Printf("Invalid resource size: %s\n", p.Size)
				return
			}
		}

		nodeList := c.Client.GetNodes()
		if p.NodeSelector == "" {
			selector := ask.PromptForNodeLabels(nodeList)
			hasSelector = "true"
			if selector == "\"-none-\"" {
				hasSelector = "false"
			}
		} else {
			hasSelector = "true"
			if selector == "-none-" {
				hasSelector = "false"
			} else {
				if !c.Client.IsValidNodeSelector(selector) {
					fmt.Printf("Invalid node selector: %s\n", selector)
					return
				}
			}
		}

		selectedSecrets := []string{}
		// p.PromptForSecrets is the local command param --secrets
		// c.PromptForSecrets is the config.yaml promptForSecrets setting
		// utilMap[utility].RequireSecrets is from the utility definitions
		if len(p.Secrets) == 0 {
			if p.PromptForSecrets || c.PromptForSecrets || utilMap[utility].RequireSecrets {
				secrets := c.Client.GetSecrets(namespace)
				selectedSecrets = ask.PromptForSecrets(secrets)
			}
		} else {
			// If the user provided secrets, use them
			if c.Client.IsValidSecrets(namespace, p.Secrets) {
				selectedSecrets = append(selectedSecrets, p.Secrets...)
			} else {
				fmt.Printf("One of the secrets provided is invalid: %s\n", p.Secrets)
				return
			}
		}

		selectedConfigMaps := []string{}
		// p.PromptForConfigMaps is the local command param --configs
		// c.PromptForConfigMaps is the config.yaml promptForConfigMaps setting
		// utilMap[utility].RequireConfigmaps is from the utility definitions
		if len(p.ConfigMaps) == 0 {
			if p.PromptForConfigMaps || c.PromptForConfigMaps || utilMap[utility].RequireConfigmaps {
				configmaps := c.Client.GetConfigMaps(namespace)
				selectedConfigMaps = ask.PromptForConfigMaps(configmaps)
			}
		} else {
			// If the user provided configmaps, use them
			if c.Client.IsValidConfigmaps(namespace, p.ConfigMaps) {
				selectedConfigMaps = append(selectedConfigMaps, p.ConfigMaps...)
			} else {
				fmt.Printf("One of the configmaps provided is invalid: %s\n", p.ConfigMaps)
				return
			}
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
				"associatedPod":  fmt.Sprintf("%s-%s", utilMap[utility].Name, shortUniq),
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

		if p.OutputName {
			fmt.Printf("%s-%s\n", utilMap[utility].Name, shortUniq)
		} else {
			if c.EnableBashLinks {
				hl := fmt.Sprintf("<bash:kubectl -n %s exec -it %s -- %s>", namespace, fmt.Sprintf("%s-%s", utilMap[utility].Name, shortUniq), utilMap[utility].ExecCommand)
				tx := fmt.Sprintf("kubectl -n %s exec -it %s -- %s", namespace, fmt.Sprintf("%s-%s", utilMap[utility].Name, shortUniq), utilMap[utility].ExecCommand)
				fmt.Println(termFormatter.Hyperlink(hl, tx))
			} else {
				fmt.Printf("kubectl -n %s exec -it %s -- %s\n", namespace, fmt.Sprintf("%s-%s", utilMap[utility].Name, shortUniq), utilMap[utility].ExecCommand)
			}
		}
	} else {
		common.Logger.Warn("Cannot launch a pod, no valid kubernetes context")
	}
}
