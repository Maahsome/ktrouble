package cmd

import (
	"context"
	"fmt"
	"strings"

	"ktrouble/common"
	"ktrouble/objects"

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

		nodeData := objects.NodeList{}
		rawData := []string{}
		for _, v := range nodeList.Items {
			nodeData.Node = append(nodeData.Node, v.Name)
			rawData = append(rawData, v.Name)
		}

		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		fmt.Println(nodeDataToString(nodeData, strings.Join(rawData, ",")))

	},
}

func getNodes() *v1.NodeList {
	cfg, err := restConfig()
	if err != nil {
		common.Logger.WithError(err).Error("could not get config")
		return &v1.NodeList{}
	}
	if cfg == nil {
		common.Logger.Error("failed to determine kubernetes config")
		return &v1.NodeList{}
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		common.Logger.WithError(err).Error("could not create client from config")
		return &v1.NodeList{}
	}

	node := client.CoreV1().Nodes()
	nodeList, err := node.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		common.Logger.WithError(err).Error("could not get list of namespaces")
		return &v1.NodeList{}
	}
	if len(nodeList.Items) == 0 {
		return &v1.NodeList{}

	}
	return nodeList

}

func nodeDataToString(nodeData objects.NodeList, raw string) string {

	switch strings.ToLower(c.OutputFormat) {
	case "raw":
		return raw
	case "json":
		return nodeData.ToJSON()
	case "gron":
		return nodeData.ToGRON()
	case "yaml":
		return nodeData.ToYAML()
	case "text", "table":
		return nodeData.ToTEXT(c.NoHeaders)
	default:
		return nodeData.ToTEXT(c.NoHeaders)
	}
}

func init() {
	getCmd.AddCommand(nodeCmd)
}
