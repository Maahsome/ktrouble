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

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"namespaces", "ns"},
	Short:   "Get a list of namespaces",
	Long: `EXAMPLE:
`,
	Run: func(cmd *cobra.Command, args []string) {

		nssList := getNamespaces()

		nsData := objects.NamespaceList{}
		rawData := []string{}
		for _, v := range nssList.Items {
			nsData.Namespace = append(nsData.Namespace, v.Name)
			rawData = append(rawData, v.Name)
		}

		if !c.FormatOverridden {
			c.OutputFormat = "text"
		}
		fmt.Println(namespaceDataToString(nsData, strings.Join(rawData, ",")))
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

func namespaceDataToString(nsData objects.NamespaceList, raw string) string {

	switch strings.ToLower(c.OutputFormat) {
	case "raw":
		return raw
	case "json":
		return nsData.ToJSON()
	case "gron":
		return nsData.ToGRON()
	case "yaml":
		return nsData.ToYAML()
	case "text", "table":
		return nsData.ToTEXT(c.NoHeaders)
	default:
		return nsData.ToTEXT(c.NoHeaders)
	}
}

func init() {
	getCmd.AddCommand(namespaceCmd)
}
