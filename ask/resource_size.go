package ask

import (
	"os"

	"ktrouble/common"
	"ktrouble/objects"

	"github.com/AlecAivazis/survey/v2"
)

type (
	ResourceSizeAnswer struct {
		ResourceSize string `survey:"resourcesize"`
	}
)

func PromptForResourceSize(sizeDefs []objects.ResourceSize) string {

	rsArray := []string{}

	for _, v := range sizeDefs {
		rsArray = append(rsArray, v.Name)
	}

	var rsSurvey = []*survey.Question{
		{
			Name: "resourcesize",
			Prompt: &survey.Select{
				Message: "Choose a resource sizing:",
				Options: rsArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	rsAnswer := &ResourceSizeAnswer{}
	if err := survey.Ask(rsSurvey, rsAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No resource size selected")
	}

	return rsAnswer.ResourceSize
}
