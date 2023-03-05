package kubernetes

import (
	"context"
	"strings"

	"ktrouble/ask"
	"ktrouble/common"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sYaml "k8s.io/apimachinery/pkg/util/yaml"
)

func (k *kubernetesClient) GetCreatedPods() *v1.PodList {
	listOptions := metav1.ListOptions{
		LabelSelector: "app=ktrouble",
	}
	podList, err := k.Client.CoreV1().Pods("").List(context.TODO(), listOptions)

	if err != nil {
		common.Logger.WithError(err).Error("could not get list of pods")
		return &v1.PodList{}
	}
	if len(podList.Items) == 0 {
		common.Logger.Error("no pods with label app=ktrouble were found on this cluster")
		return &v1.PodList{}
	}
	return podList
}

func (k *kubernetesClient) DeletePod(pod ask.PodDetail) error {
	podClient := k.Client.CoreV1().Pods(pod.Namespace)
	derr := podClient.Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
	if derr != nil {
		// common.Logger.WithError(derr).Fatal("Failed to delete pod")
		return derr
	}
	return nil
}

func (k *kubernetesClient) CreatePod(podJSON string, namespace string) {

	podClient := k.Client.CoreV1().Pods(namespace)
	newResource := &v1.Pod{}
	if err := k8sYaml.NewYAMLOrJSONDecoder(strings.NewReader(podJSON), 100).Decode(&newResource); err != nil {
		common.Logger.Errorf("Error converting to K8s: %s", podJSON)
	}
	// result, err := podClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	_, cerr := podClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	if cerr != nil {
		common.Logger.WithError(cerr).Fatal("Failed to create pod")
	}
}
