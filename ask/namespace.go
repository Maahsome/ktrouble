package ask

import (
	"os"
	"sort"

	"ktrouble/common"

	"github.com/AlecAivazis/survey/v2"
	v1 "k8s.io/api/core/v1"
)

type (
	NamespaceAnswer struct {
		Namespace string `survey:"namespace"`
	}
)

func PromptForNamespace(nssList *v1.NamespaceList) string {

	var nsArray []string
	for _, v := range nssList.Items {
		nsArray = append(nsArray, v.Name)
	}
	sort.Strings(nsArray)

	var nsSurvey = []*survey.Question{
		{
			Name: "namespace",
			Prompt: &survey.Select{
				Message: "Choose a namespace:",
				Options: nsArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	nsAnswer := &NamespaceAnswer{}
	if err := survey.Ask(nsSurvey, nsAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No namespace selected")
	}
	return nsAnswer.Namespace
}
