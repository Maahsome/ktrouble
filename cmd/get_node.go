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

// nodeCmd represents the namespace command
var nodeCmd = &cobra.Command{
	Use:     "node",
	Aliases: []string{"nodes"},
	Short:   "Get a list of node labels",
	Long: `EXAMPLE:
	> ktrouble get node
`,
	Run: func(cmd *cobra.Command, args []string) {

		nodeList := getNodes()

		fmt.Println("NODES")
		fmt.Println("---------------")
		for _, v := range nodeList.Items {
			fmt.Println(v.Name)
		}
	},
}

func getNodes() *v1.NodeList {
	cfg, err := restConfig()
	if err != nil {
		logrus.WithError(err).Error("could not get config")
		return &v1.NodeList{}
	}
	if cfg == nil {
		logrus.Error("failed to determine kubernetes config")
		return &v1.NodeList{}
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logrus.WithError(err).Error("could not create client from config")
		return &v1.NodeList{}
	}

	node := client.CoreV1().Nodes()
	nodeList, err := node.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logrus.WithError(err).Error("could not get list of namespaces")
		return &v1.NodeList{}
	}
	if len(nodeList.Items) == 0 {
		return &v1.NodeList{}

	}
	return nodeList

}

func init() {
	getCmd.AddCommand(nodeCmd)
}
