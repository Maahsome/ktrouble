package ask

import (
	"os"
	"sort"

	"ktrouble/common"

	"github.com/AlecAivazis/survey/v2"
)

type (
	GenericAnswer struct {
		Name []string `survey:"name"`
	}
)

func PromptForGenericList(list []string, prompt string) []string {

	sort.Strings(list)

	var utilSurvey = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.MultiSelect{
				Message: prompt,
				Options: list,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	genericAnswer := &GenericAnswer{}
	if err := survey.Ask(utilSurvey, genericAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("Nothing selected")
	}
	return genericAnswer.Name
}
