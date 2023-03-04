package cmd

import (
	"context"
	"fmt"

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

		fmt.Println("SERVICE ACCOUNT")
		fmt.Println("---------------")
		for _, v := range sasList.Items {
			fmt.Println(v.Name)
		}

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

func init() {
	getCmd.AddCommand(serviceaccountCmd)
}
