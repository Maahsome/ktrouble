package cmd

import (
	"context"
	"fmt"
	"ktrouble/objects"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// serviceaccountCmd represents the serviceaccount command
var serviceaccountCmd = &cobra.Command{
	Use:     "serviceaccount",
	Aliases: []string{"serviceaccounts", "sa"},
	Short:   "Get a list of K8s ServiceAccount(s) in a Namespace",
	Long: `EXAMPLE:
	> ktrouble get serviceaccount -n myspace
`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace := determineNamespace()

		sasList := getServiceAccounts(namespace)

		saData := objects.ServiceAccountList{}
		rawData := []string{}
		for _, v := range sasList.Items {
			saData.ServiceAccount = append(saData.ServiceAccount, v.Name)
			rawData = append(rawData, v.Name)
		}

		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		fmt.Println(saDataToString(saData, strings.Join(rawData, ",")))

	},
}

func getServiceAccounts(namespace string) *v1.ServiceAccountList {

	cfg, err := restConfig()
	if err != nil {
		logrus.WithError(err).Error("could not get config")
		return &v1.ServiceAccountList{}
	}
	if cfg == nil {
		logrus.Error("failed to determine kubernetes config")
		return &v1.ServiceAccountList{}
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logrus.WithError(err).Error("could not create client from config")
		return &v1.ServiceAccountList{}
	}

	sas := client.CoreV1().ServiceAccounts(namespace)
	sasList, err := sas.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logrus.WithError(err).Error("could not get list of service accounts")
		return &v1.ServiceAccountList{}
	}
	if len(sasList.Items) == 0 {
		logrus.Errorf("no serviceaccounts were found in namespace: %s", namespace)
		return &v1.ServiceAccountList{}
	}
	return sasList
}

func saDataToString(saData objects.ServiceAccountList, raw string) string {

	switch strings.ToLower(c.OutputFormat) {
	case "raw":
		return raw
	case "json":
		return saData.ToJSON()
	case "gron":
		return saData.ToGRON()
	case "yaml":
		return saData.ToYAML()
	case "text", "table":
		return saData.ToTEXT(c.NoHeaders)
	default:
		return saData.ToTEXT(c.NoHeaders)
	}
}

func init() {
	getCmd.AddCommand(serviceaccountCmd)
}
