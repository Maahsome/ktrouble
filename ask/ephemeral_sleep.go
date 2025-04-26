package ask

import (
	"os"

	"ktrouble/common"
	"ktrouble/objects"

	"github.com/AlecAivazis/survey/v2"
)

type (
	EphemeralSleepAnswer struct {
		EphemeralSleep string `survey:"ephemeralsleep"`
	}
)

func PromptForEphemeralSleep(ephemeralSleepDefs []objects.EphemeralSleep) string {

	esArray := []string{}
	esMap := make(map[string]objects.EphemeralSleep)

	for _, v := range ephemeralSleepDefs {
		esArray = append(esArray, v.Name)
		esMap[v.Name] = v
	}

	var esSurvey = []*survey.Question{
		{
			Name: "ephemeralsleep",
			Prompt: &survey.Select{
				Message: "Choose a sleep time:",
				Options: esArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	esAnswer := &EphemeralSleepAnswer{}
	if err := survey.Ask(esSurvey, esAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No sleep time selected")
	}

	return esMap[esAnswer.EphemeralSleep].Seconds
}
