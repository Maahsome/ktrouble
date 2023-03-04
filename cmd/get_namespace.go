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

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"namespaces", "ns"},
	Short:   "Get a list of namespaces",
	Long: `EXAMPLE:
`,
	Run: func(cmd *cobra.Command, args []string) {

		nssList := getNamespaces()

		fmt.Println("NAMESPACE")
		fmt.Println("---------------")
		for _, v := range nssList.Items {
			fmt.Println(v.Name)
		}
	},
}

func getNamespaces() *v1.NamespaceList {
	cfg, err := restConfig()
	if err != nil {
		logrus.WithError(err).Error("could not get config")
		return &v1.NamespaceList{}
	}
	if cfg == nil {
		logrus.Error("failed to determine kubernetes config")
		return &v1.NamespaceList{}
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logrus.WithError(err).Error("could not create client from config")
		return &v1.NamespaceList{}
	}

	nss := client.CoreV1().Namespaces()
	nssList, err := nss.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logrus.WithError(err).Error("could not get list of namespaces")
		return &v1.NamespaceList{}
	}
	if len(nssList.Items) == 0 {
		return &v1.NamespaceList{}

	}
	return nssList

}

func init() {
	getCmd.AddCommand(namespaceCmd)
}
