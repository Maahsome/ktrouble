package kubernetes

import (
	"context"
	"fmt"
	"os"
	"strings"

	"ktrouble/ask"
	"ktrouble/common"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sYaml "k8s.io/apimachinery/pkg/util/yaml"
)

func (k *kubernetesClient) GetAssociatedService(pod ask.PodDetail) *v1.ServiceList {
	labelSelector := fmt.Sprintf("app=ktrouble,launchedby=%s,associatedPod=%s", os.Getenv("USER"), pod.Name)
	listOptions := metav1.ListOptions{
		LabelSelector: labelSelector,
	}
	serviceList, err := k.Client.CoreV1().Services("").List(context.TODO(), listOptions)

	if err != nil {
		common.Logger.WithError(err).Error("could not get list of services")
		return &v1.ServiceList{}
	}
	if len(serviceList.Items) == 0 {
		common.Logger.Tracef("no services with labels %s were found on this cluster", labelSelector)
		return &v1.ServiceList{}
	}
	return serviceList
}

func (k *kubernetesClient) DeleteAssociatedService(pod ask.PodDetail) error {

	common.Logger.Tracef("Deleting associated services for pod %s", pod.Name)
	associatedServices := k.GetAssociatedService(pod)
	if len(associatedServices.Items) == 0 {
		common.Logger.Tracef("no services associated with pod %s were found on this cluster", pod.Name)
		return nil
	}
	serviceClient := k.Client.CoreV1().Services(pod.Namespace)
	for _, service := range associatedServices.Items {
		common.Logger.Tracef("Deleting service: %s", service.Name)
		derr := serviceClient.Delete(context.TODO(), service.Name, metav1.DeleteOptions{})
		if derr != nil {
			common.Logger.WithError(derr).Errorf("Failed to delete service: %s", service.Name)
		}
	}
	return nil
}

func (k *kubernetesClient) CreateService(serviceJSON string, namespace string) {

	serviceClient := k.Client.CoreV1().Services(namespace)
	newResource := &v1.Service{}
	if err := k8sYaml.NewYAMLOrJSONDecoder(strings.NewReader(serviceJSON), 100).Decode(&newResource); err != nil {
		common.Logger.Errorf("Error converting to K8s: %s", serviceJSON)
	}
	// result, err := podClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	_, cerr := serviceClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	if cerr != nil {
		common.Logger.WithError(cerr).Fatal("Failed to create service")
	}
}
