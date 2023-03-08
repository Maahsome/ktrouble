package ask

import (
	"os"
	"sort"

	"ktrouble/common"

	"github.com/AlecAivazis/survey/v2"
	v1 "k8s.io/api/core/v1"
)

type (
	MountConfigMapsAnswer struct {
		ConfigMap []string `survey:"configmap"`
	}
)

func PromptForConfigMaps(configmapList *v1.ConfigMapList) []string {

	var cmArray []string
	for _, v := range configmapList.Items {
		cmArray = append(cmArray, v.Name)
	}
	sort.Strings(cmArray)

	var cmSurvey = []*survey.Question{
		{
			Name: "configmap",
			Prompt: &survey.MultiSelect{
				Message: "Choose the configmaps to mount in the POD:",
				Options: cmArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	cmAnswer := &MountConfigMapsAnswer{}
	if err := survey.Ask(cmSurvey, cmAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No configmaps selected")
	}
	return cmAnswer.ConfigMap
}
