package kubernetes

import (
	"context"

	"ktrouble/common"

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
