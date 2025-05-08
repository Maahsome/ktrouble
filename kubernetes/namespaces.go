package kubernetes

import (
	"context"
	"ktrouble/ask"
	"ktrouble/common"
	"os"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNamespaces returns a list of namespaces configured on the current context
// kubernetes cluster.  Similar to "kubectl get namespaces"
func (k *kubernetesClient) GetNamespaces() *v1.NamespaceList {
	nss := k.Client.CoreV1().Namespaces()
	nssList, err := nss.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		common.Logger.WithError(err).Error("could not get list of namespaces")
		return &v1.NamespaceList{}
	}
	if len(nssList.Items) == 0 {
		return &v1.NamespaceList{}

	}
	return nssList

}

func (k *kubernetesClient) DetermineNamespace(nsParam string) string {

	namespace := ""
	if len(os.Getenv("NAMESPACE")) > 0 {
		namespace = os.Getenv("NAMESPACE")
		if len(nsParam) > 0 {
			namespace = nsParam
		}
	} else {
		if len(nsParam) > 0 {
			namespace = nsParam
		}
	}

	if namespace == "" {
		nssList := k.GetNamespaces()
		namespace = ask.PromptForNamespace(nssList)
	}
	return namespace
}

func (k *kubernetesClient) IsNamespaceValid(namespace string) bool {
	nssList := k.GetNamespaces()
	for _, ns := range nssList.Items {
		if ns.Name == namespace {
			return true
		}
	}
	return false
}
