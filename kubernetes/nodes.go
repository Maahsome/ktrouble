package kubernetes

import (
	"context"
	"fmt"

	"ktrouble/common"
	"ktrouble/defaults"

	"github.com/spf13/viper"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNodes returns the nodes of the current context kubernetes cluster
// Similar to "kubectl get nodes"
func (k *kubernetesClient) GetNodes() *v1.NodeList {
	node := k.Client.CoreV1().Nodes()
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

func (k *kubernetesClient) IsValidNodeSelector(selector string) bool {
	var labelList []string
	err := viper.UnmarshalKey("nodeSelectorLabels", &labelList)
	if err != nil {
		common.Logger.Fatal("Error unmarshalling...")
	}
	if len(labelList) == 0 {
		labelList = defaults.Labels()
	}
	labelMap := make(map[string]string, len(labelList))
	for _, v := range labelList {
		labelMap[v] = v
	}

	nodeList := k.GetNodes()

	nodeArray := make(map[string]string, len(nodeList.Items))

	for _, v := range nodeList.Items {
		for k, l := range v.Labels {
			if _, mok := labelMap[k]; mok {
				if _, ok := nodeArray[fmt.Sprintf("\"%s\": \"%s\"", k, l)]; !ok {
					nodeArray[fmt.Sprintf("\"%s\": \"%s\"", k, l)] = fmt.Sprintf("\"%s\": \"%s\"", k, l)
				}
			}
		}
	}

	if _, ok := nodeArray[selector]; ok {
		return true
	}
	return false
}
