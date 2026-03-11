package ask

import (
	"os"

	"ktrouble/common"
	"ktrouble/objects"

	"github.com/AlecAivazis/survey/v2"
)

type (
	EnvironmentAnswer struct {
		Environment string `survey:"environment"`
	}
)

func PromptForEnvironment(envDefs objects.EnvironmentList) string {

	envArray := []string{}

	for _, v := range envDefs {
		envArray = append(envArray, v.Name)
	}

	var envSurvey = []*survey.Question{
		{
			Name: "environment",
			Prompt: &survey.Select{
				Message: "Choose an environment:",
				Options: envArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	envAnswer := &EnvironmentAnswer{}
	if err := survey.Ask(envSurvey, envAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No environment selected")
	}

	return envAnswer.Environment
}
