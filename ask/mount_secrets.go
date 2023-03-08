package ask

import (
	"os"
	"sort"

	"ktrouble/common"

	"github.com/AlecAivazis/survey/v2"
	v1 "k8s.io/api/core/v1"
)

type (
	MountSecretsAnswer struct {
		Secret []string `survey:"secret"`
	}
)

func PromptForSecrets(secretList *v1.SecretList) []string {

	var secretArray []string
	for _, v := range secretList.Items {
		secretArray = append(secretArray, v.Name)
	}
	sort.Strings(secretArray)

	var secretSurvey = []*survey.Question{
		{
			Name: "secret",
			Prompt: &survey.MultiSelect{
				Message: "Choose the secrets to mount in the POD:",
				Options: secretArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	secretAnswer := &MountSecretsAnswer{}
	if err := survey.Ask(secretSurvey, secretAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No secrets selected")
	}
	return secretAnswer.Secret
}
