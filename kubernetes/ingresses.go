package kubernetes

import (
	"context"
	"fmt"
	"os"
	"strings"

	"ktrouble/ask"
	"ktrouble/common"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sYaml "k8s.io/apimachinery/pkg/util/yaml"
)

func (k *kubernetesClient) GetAssociatedIngress(pod ask.PodDetail) *networkingv1.IngressList {
	labelSelector := fmt.Sprintf("app=ktrouble,launchedby=%s,associatedPod=%s", os.Getenv("USER"), pod.Name)
	listOptions := metav1.ListOptions{
		LabelSelector: labelSelector,
	}
	ingressList, err := k.Client.NetworkingV1().Ingresses("").List(context.TODO(), listOptions)

	if err != nil {
		common.Logger.WithError(err).Error("could not get list of services")
		return &networkingv1.IngressList{}
	}
	if len(ingressList.Items) == 0 {
		common.Logger.Tracef("no ingresses with labels %s were found on this cluster", labelSelector)
		return &networkingv1.IngressList{}
	}
	return ingressList
}

func (k *kubernetesClient) DeleteAssociatedIngress(pod ask.PodDetail) error {

	common.Logger.Tracef("Deleting associated ingresses for pod %s", pod.Name)
	associatedIngress := k.GetAssociatedIngress(pod)
	if len(associatedIngress.Items) == 0 {
		common.Logger.Tracef("no ingresses associated with pod %s were found on this cluster", pod.Name)
		return nil
	}
	ingressClient := k.Client.NetworkingV1().Ingresses(pod.Namespace)
	for _, ingress := range associatedIngress.Items {
		common.Logger.Tracef("Deleting ingress: %s", ingress.Name)
		derr := ingressClient.Delete(context.TODO(), ingress.Name, metav1.DeleteOptions{})
		if derr != nil {
			common.Logger.WithError(derr).Errorf("Failed to delete ingress: %s", ingress.Name)
		}
	}
	return nil
}

func (k *kubernetesClient) CreateIngress(serviceJSON string, namespace string) {

	ingressClient := k.Client.NetworkingV1().Ingresses(namespace)
	newResource := &networkingv1.Ingress{}
	if err := k8sYaml.NewYAMLOrJSONDecoder(strings.NewReader(serviceJSON), 100).Decode(&newResource); err != nil {
		common.Logger.Errorf("Error converting to K8s: %s", serviceJSON)
	}
	_, cerr := ingressClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	if cerr != nil {
		common.Logger.WithError(cerr).Fatal("Failed to create ingress")
	}
}
