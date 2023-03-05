package ask

import (
	"os"
	"sort"

	"ktrouble/common"

	"github.com/AlecAivazis/survey/v2"
	v1 "k8s.io/api/core/v1"
)

type (
	ServiceAccountAnswer struct {
		ServiceAccount string `survey:"serviceaccount"`
	}
)

func PromptForServiceAccount(sasList *v1.ServiceAccountList) string {

	var saArray []string
	for _, v := range sasList.Items {
		saArray = append(saArray, v.Name)
	}
	sort.Strings(saArray)

	var saSurvey = []*survey.Question{
		{
			Name: "serviceaccount",
			Prompt: &survey.Select{
				Message: "Choose a service account to run the pod under:",
				Options: saArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	saAnswer := &ServiceAccountAnswer{}
	if err := survey.Ask(saSurvey, saAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No service account selected")
	}
	return saAnswer.ServiceAccount
}
