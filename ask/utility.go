package ask

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"ktrouble/common"
	"ktrouble/objects"

	"github.com/AlecAivazis/survey/v2"
)

type (
	UtilAnswer struct {
		UtilityName string `survey:"utilityname"`
	}
)

func PromptForUtility(utils []objects.UtilityPod, envMap map[string]objects.Environment, showHidden bool) (string, objects.SelectedUtilityPod) {

	var utilArray []string
	for _, v := range utils {
		if !v.Hidden || showHidden {
			if len(v.Environments) == 0 {
				for _, tag := range v.Tags {
					utilArray = append(utilArray, fmt.Sprintf("%s:%s", v.Name, tag))
				}
			} else {
				for _, env := range v.Environments {
					if !envMap[env].Hidden || showHidden {
						for _, tag := range v.Tags {
							utilArray = append(utilArray, fmt.Sprintf("%s/%s:%s", env, v.Name, tag))
						}
					}
				}
			}
		}
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
	return utilAnswer.UtilityName, buildSelectedPod(utilAnswer.UtilityName)
}

func buildSelectedPod(utilityName string) objects.SelectedUtilityPod {
	if strings.Contains(utilityName, "/") {
		// extract environment
		envParts := strings.Split(utilityName, "/")
		utilParts := strings.Split(envParts[1], ":")
		return objects.SelectedUtilityPod{
			Name:        utilParts[0],
			Image:       fmt.Sprintf("%s:latest", utilParts[0]),
			Environment: envParts[0],
			Tag:         utilParts[1],
		}
	} else {
		// there is no environment
		utilParts := strings.Split(utilityName, ":")
		return objects.SelectedUtilityPod{
			Name:        utilParts[0],
			Image:       fmt.Sprintf("%s:latest", utilParts[0]),
			Environment: "",
			Tag:         utilParts[1],
		}
	}
}
