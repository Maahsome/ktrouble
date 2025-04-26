package kubernetes

import (
	"context"
	"fmt"
	"os"
	"strings"

	"ktrouble/ask"
	"ktrouble/common"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sYaml "k8s.io/apimachinery/pkg/util/yaml"
)

func (k *kubernetesClient) GetNamespacePods(namespace string) *corev1.PodList {
	listOptions := metav1.ListOptions{}
	podList, err := k.Client.CoreV1().Pods(namespace).List(context.TODO(), listOptions)

	if err != nil {
		common.Logger.WithError(err).Error("could not get list of pods")
		return &corev1.PodList{}
	}
	if len(podList.Items) == 0 {
		common.Logger.Errorf("no pods in this namespace: %s", namespace)
		return &corev1.PodList{}
	}
	return podList
}

func (k *kubernetesClient) GetCreatedPods(all bool) *corev1.PodList {
	labelSelector := fmt.Sprintf("app=ktrouble,launchedby=%s", os.Getenv("USER"))
	if all {
		labelSelector = "app=ktrouble"
	}
	listOptions := metav1.ListOptions{
		LabelSelector: labelSelector,
	}
	podList, err := k.Client.CoreV1().Pods("").List(context.TODO(), listOptions)

	if err != nil {
		common.Logger.WithError(err).Error("could not get list of pods")
		return &corev1.PodList{}
	}
	if len(podList.Items) == 0 {
		common.Logger.Error("no pods with label app=ktrouble were found on this cluster")
		return &corev1.PodList{}
	}
	return podList
}

func (k *kubernetesClient) GetAttachedContainers(all bool) *corev1.PodList {
	labelSelector := fmt.Sprintf("ktrouble=container,ktrouble.launchedby=%s", os.Getenv("USER"))
	if all {
		labelSelector = "ktrouble=container"
	}
	listOptions := metav1.ListOptions{
		LabelSelector: labelSelector,
	}
	podList, err := k.Client.CoreV1().Pods("").List(context.TODO(), listOptions)

	if err != nil {
		common.Logger.WithError(err).Error("could not get list of pods")
		return &corev1.PodList{}
	}
	if len(podList.Items) == 0 {
		common.Logger.Error("no pods with label ktrouble=container were found on this cluster")
		return &corev1.PodList{}
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
	newResource := &corev1.Pod{}
	if err := k8sYaml.NewYAMLOrJSONDecoder(strings.NewReader(podJSON), 100).Decode(&newResource); err != nil {
		common.Logger.Errorf("Error converting to K8s: %s", podJSON)
	}
	// result, err := podClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	_, cerr := podClient.Create(context.TODO(), newResource, metav1.CreateOptions{})
	if cerr != nil {
		common.Logger.WithError(cerr).Fatal("Failed to create pod")
	}
}

func (k *kubernetesClient) AttachContainerToPod(
	namespace string,
	podName string,
	containerName string,
	image string,
	sleep string,
	mounts []corev1.VolumeMount,
) error {

	podClient := k.Client.CoreV1().Pods(namespace)
	pod, err := podClient.Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	ephemeralContainer := corev1.EphemeralContainer{
		EphemeralContainerCommon: corev1.EphemeralContainerCommon{
			Name:         containerName,
			Image:        image,
			Command:      []string{"sleep", sleep},
			TTY:          false,
			VolumeMounts: mounts,
		},
	}

	// Add the ephemeral container to the pod
	pod.Spec.EphemeralContainers = append(pod.Spec.EphemeralContainers, ephemeralContainer)

	// Update the pod with the ephemeral container and return the pod so we can update the labels
	_, err = podClient.UpdateEphemeralContainers(context.TODO(), podName, pod, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	pod, err = podClient.Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	// Add a ktrouble label to the pod
	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
		pod.Labels["ktrouble"] = "container"
		pod.Labels["ktrouble.launchedby"] = os.Getenv("USER")
	} else {
		pod.Labels["ktrouble"] = "container"
		pod.Labels["ktrouble.launchedby"] = os.Getenv("USER")
	}

	_, err = podClient.Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (k *kubernetesClient) GetPodMounts(namespace string, podName string) []corev1.VolumeMount {

	validMounts := []corev1.VolumeMount{}
	podClient := k.Client.CoreV1().Pods(namespace)
	pod, err := podClient.Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return validMounts
	}

	// Ephemeral Container do NOT support subPath based mounts
	for _, c := range pod.Spec.Containers {
		for _, m := range c.VolumeMounts {
			if m.SubPath == "" {
				validMounts = append(validMounts, m)
			}
		}
	}
	return validMounts
}
