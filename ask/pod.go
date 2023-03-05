package ask

import (
	"fmt"
	"os"

	"ktrouble/common"

	"github.com/AlecAivazis/survey/v2"
	v1 "k8s.io/api/core/v1"
)

type (
	PodDetail struct {
		Name      string `survey:"podname"`
		Namespace string `survey:"namespace"`
		Deleted   string `survey:"deleted"`
	}
	PodAnswer struct {
		Pod string `survey:"podname"`
	}
)

func PromptForPod(podList *v1.PodList) PodDetail {

	podArray := make(map[string]PodDetail, len(podList.Items))

	for _, v := range podList.Items {
		deleting := "false"
		if v.DeletionTimestamp != nil {
			deleting = "true"
		}
		podArray[fmt.Sprintf("%s/%s", v.Namespace, v.Name)] = PodDetail{
			Name:      v.Name,
			Namespace: v.Namespace,
			Deleted:   deleting,
		}
	}

	podNames := []string{}
	for _, m := range podArray {
		if m.Deleted == "false" {
			podNames = append(podNames, fmt.Sprintf("%s/%s", m.Namespace, m.Name))
		}
	}

	var podSurvey = []*survey.Question{
		{
			Name: "podname",
			Prompt: &survey.Select{
				Message: "Choose a pod to delete:",
				Options: podNames,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	podAnswer := &PodAnswer{}
	if err := survey.Ask(podSurvey, podAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No pod selected")
	}
	return podArray[podAnswer.Pod]
}
