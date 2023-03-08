package kubernetes

import (
	"context"
	"ktrouble/common"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *kubernetesClient) GetSecrets(namespace string) *v1.SecretList {
	ss := k.Client.CoreV1().Secrets(namespace)
	secretList, err := ss.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		common.Logger.WithError(err).Error("could not get list of secrets")
		return &v1.SecretList{}
	}
	if len(secretList.Items) == 0 {
		common.Logger.Errorf("no secrets were found in namespace: %s", namespace)
		return &v1.SecretList{}
	}
	return secretList
}
