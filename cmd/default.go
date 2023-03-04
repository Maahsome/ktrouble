package cmd

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"strings"

	"ktrouble/template"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sYaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

type TemplateConfig struct {
	Parameters map[string]string
}

type UtilityPod struct {
	Name        string
	Repository  string
	ExecCommand string
}

var letters = []rune("abcdef0987654321")

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:   "launch",
	Short: "launch a kubernetes troubleshooting pod",
	Long: `EXAMPLE:
`,
	Run: func(cmd *cobra.Command, args []string) {
		utilDefs := buildUtilityDefinitions()
		// for i, v := range args {
		// 	fmt.Printf("ARG: %d: %s\n", i, v)
		// }
		utility := ""
		if len(args) > 0 && len(args[0]) > 0 {
			utility = args[0]
		} else {
			utility = askForUtility(utilDefs)
		}
		namespace := determineNamespace()
		sa := "default"
		if len(args) > 1 && len(args[1]) > 0 {
			sa = args[1]
		} else {
			sasList := getServiceAccounts(namespace)
			sa = askForServiceAccount(sasList)
		}

		resourceSize := askForResourceSize()

		nodeList := getNodes()
		selector := askForNodeLabels(nodeList)
		hasSelector := "true"
		if selector == "\"-none-\"" {
			hasSelector = "false"
		}
		shortUniq := randSeq(6)
		tc := &TemplateConfig{
			Parameters: map[string]string{
				"name":           fmt.Sprintf("%s-%s", utility, shortUniq),
				"serviceAccount": sa,
				"namespace":      namespace,
				"registry":       utilDefs[utility].Repository,
				"limitsCpu":      resourceSize["limitsCpu"],
				"limitsMem":      resourceSize["limitsMem"],
				"requestCpu":     resourceSize["requestCpu"],
				"requestMem":     resourceSize["requestMem"],
				"hasSelector":    hasSelector,
				"selector":       selector,
			},
		}
		// err := template.ApplicationsTemplate.Execute(os.Stdout, tc)
		// if err != nil {
		// 	logrus.WithError(err).Error("unable to generate the template data")
		// }

		var tpl bytes.Buffer
		if err := template.ApplicationsTemplate.Execute(&tpl, tc); err != nil {
			logrus.WithError(err).Error("unable to generate the template data")
		}

		podManifest := tpl.String()

		createPod(podManifest, namespace)

		fmt.Printf("kubectl -n %s exec -it %s -- %s\n", namespace, fmt.Sprintf("%s-%s", utility, shortUniq), utilDefs[utility].ExecCommand)

	},
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createPod(podJSON string, namespace string) {

	cfg, err := restConfig()
	if err != nil {
		logrus.WithError(err).Fatal("could not get config")
	}
	if cfg == nil {
		logrus.Fatal("failed to determine kubernetes config")
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logrus.WithError(err).Fatal("could not create client from config")
	}

	podClient := client.CoreV1().Pods(namespace)
	newResource := &v1.Pod{}
	if err := k8sYaml.NewYAMLOrJSONDecoder(strings.NewReader(podJSON), 100).Decode(&newResource); err != nil {
		logrus.Errorf("Error converting to K8s: %s", podJSON)
	}
	// result, err := podClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	_, cerr := podClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	if cerr != nil {
		logrus.WithError(cerr).Fatal("Failed to create pod")
	}
	// fmt.Printf("Created pod %q.\n", result.GetObjectMeta().GetName())
}
func buildUtilityDefinitions() map[string]UtilityPod {

	utilityDefinitions := make(map[string]UtilityPod, 0)
	utilityDefinitions["dnsutils"] = UtilityPod{
		Name:        "dnsutils",
		Repository:  "gcr.io/kubernetes-e2e-test-images/dnsutils:1.3",
		ExecCommand: "/bin/sh",
	}
	utilityDefinitions["psql-curl"] = UtilityPod{
		Name:        "psql-curl",
		Repository:  "barrypiccinni/psql-curl:latest",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["psqlutils15"] = UtilityPod{
		Name:        "psqlutils15",
		Repository:  "postgres:15-bullseye",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["psqlutils14"] = UtilityPod{
		Name:        "psqlutils14",
		Repository:  "postgres:14-bullseye",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["awscli"] = UtilityPod{
		Name:        "awscli",
		Repository:  "amazon/aws-cli:latest",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["gcloudutils"] = UtilityPod{
		Name:        "gcloudutils",
		Repository:  "google/cloud-sdk:latest",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["azutils"] = UtilityPod{
		Name:        "azutils",
		Repository:  "mcr.microsoft.com/azure-cli",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["mysqlutils5"] = UtilityPod{
		Name:        "mysqlutils5",
		Repository:  "mysql:5.7.40-debian",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["mysqlutils8"] = UtilityPod{
		Name:        "mysqlutils8",
		Repository:  "mysql:8-debian",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["redis6"] = UtilityPod{
		Name:        "redis6",
		Repository:  "cmaahs/redis-cli:6.2",
		ExecCommand: "/bin/bash",
	}
	utilityDefinitions["curl"] = UtilityPod{
		Name:        "curl",
		Repository:  "curlimages/curl:latest",
		ExecCommand: "/bin/sh",
	}

	return utilityDefinitions
}

func init() {
	rootCmd.AddCommand(defaultCmd)
}
