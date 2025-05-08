package kubernetes

import (
	"context"
	"ktrouble/common"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *kubernetesClient) GetConfigMaps(namespace string) *v1.ConfigMapList {
	cm := k.Client.CoreV1().ConfigMaps(namespace)
	configmapList, err := cm.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		common.Logger.WithError(err).Error("could not get list of secrets")
		return &v1.ConfigMapList{}
	}
	if len(configmapList.Items) == 0 {
		common.Logger.Errorf("no configmaps were found in namespace: %s", namespace)
		return &v1.ConfigMapList{}
	}
	return configmapList
}

func (k *kubernetesClient) IsValidConfigmaps(namespace string, configmaps []string) bool {
	configmapList := k.GetConfigMaps(namespace)
	if len(configmapList.Items) == 0 {
		return false
	}

	configmapMap := make(map[string]string, len(configmapList.Items))
	for _, v := range configmapList.Items {
		configmapMap[v.Name] = v.Name
	}
	for _, v := range configmaps {
		if _, ok := configmapMap[v]; !ok {
			return false
		}
	}
	return true
}
