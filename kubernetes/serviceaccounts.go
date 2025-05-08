package kubernetes

import (
	"context"
	"ktrouble/common"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *kubernetesClient) GetServiceAccounts(namespace string) *v1.ServiceAccountList {
	sas := k.Client.CoreV1().ServiceAccounts(namespace)
	sasList, err := sas.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		common.Logger.WithError(err).Error("could not get list of service accounts")
		return &v1.ServiceAccountList{}
	}
	if len(sasList.Items) == 0 {
		common.Logger.Errorf("no serviceaccounts were found in namespace: %s", namespace)
		return &v1.ServiceAccountList{}
	}
	return sasList
}

func (k *kubernetesClient) IsServiceAccountValid(namespace string, sa string) bool {
	sasList := k.GetServiceAccounts(namespace)
	for _, saItem := range sasList.Items {
		if saItem.Name == sa {
			return true
		}
	}
	return false
}
