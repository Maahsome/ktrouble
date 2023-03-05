package ask

import (
	"os"
	"sort"

	"ktrouble/common"
	"ktrouble/objects"

	"github.com/AlecAivazis/survey/v2"
)

type (
	UtilAnswer struct {
		UtilityName string `survey:"utilityname"`
	}
)

func PromptForUtility(utils []objects.UtilityPod) string {

	var utilArray []string
	for _, v := range utils {
		utilArray = append(utilArray, v.Name)
	}
	sort.Strings(utilArray)

	var utilSurvey = []*survey.Question{
		{
			Name: "utilityname",
			Prompt: &survey.Select{
				Message: "Choose a utility to create a pod with:",
				Options: utilArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	utilAnswer := &UtilAnswer{}
	if err := survey.Ask(utilSurvey, utilAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No utility selected")
	}
	return utilAnswer.UtilityName
}
