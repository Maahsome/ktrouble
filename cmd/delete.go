package cmd

import (
	"context"

	"ktrouble/common"

	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete PODs that have been created by ktrouble",
	Long: `EXAMPLE:
	> ktrouble delete
`,
	Run: func(cmd *cobra.Command, args []string) {
		podList := getCreatedPods()

		if len(podList.Items) > 0 {
			selectedPod := askForPod(podList)

			deletePod(selectedPod)
		}
	},
}

func getCreatedPods() *v1.PodList {

	cfg, err := restConfig()
	if err != nil {
		common.Logger.WithError(err).Error("could not get config")
		return &v1.PodList{}
	}
	if cfg == nil {
		common.Logger.Error("failed to determine kubernetes config")
		return &v1.PodList{}
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		common.Logger.WithError(err).Error("could not create client from config")
		return &v1.PodList{}
	}

	listOptions := metav1.ListOptions{
		LabelSelector: "app=ktrouble",
	}
	podList, err := client.CoreV1().Pods("").List(context.TODO(), listOptions)

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

func deletePod(pod PodDetail) {

	cfg, err := restConfig()
	if err != nil {
		common.Logger.WithError(err).Fatal("could not get config")
	}
	if cfg == nil {
		common.Logger.Fatal("failed to determine kubernetes config")
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		common.Logger.WithError(err).Fatal("could not create client from config")
	}

	podClient := client.CoreV1().Pods(pod.Namespace)
	derr := podClient.Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
	if derr != nil {
		common.Logger.WithError(derr).Fatal("Failed to delete pod")
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
