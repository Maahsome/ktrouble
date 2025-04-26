package ask

import (
	"os"
	"sort"

	"ktrouble/common"

	"github.com/AlecAivazis/survey/v2"
	v1 "k8s.io/api/core/v1"
)

type (
	EphemeralMountAnswer struct {
		Mount []string `survey:"mounts"`
	}
)

func PromptForEphemeralMounts(mountList []v1.VolumeMount) []v1.VolumeMount {

	var mountArray []string
	mountMap := make(map[string]v1.VolumeMount)
	for _, v := range mountList {
		mountArray = append(mountArray, v.Name)
		mountMap[v.Name] = v
	}
	sort.Strings(mountArray)

	var mountSurvey = []*survey.Question{
		{
			Name: "mount",
			Prompt: &survey.MultiSelect{
				Message: "Choose the mounts to add to the container:",
				Options: mountArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	mountAnswer := &EphemeralMountAnswer{}
	if err := survey.Ask(mountSurvey, mountAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No mounts selected")
	}

	mounts := []v1.VolumeMount{}
	for _, v := range mountAnswer.Mount {
		mounts = append(mounts, mountMap[v])
	}
	return mounts
}
