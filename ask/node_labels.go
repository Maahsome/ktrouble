package ask

import (
	"fmt"
	"os"
	"sort"

	"ktrouble/common"
	"ktrouble/defaults"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	v1 "k8s.io/api/core/v1"
)

type (
	NodeDetail struct {
		Name   string   `survey:"nodename"`
		Labels []string `survey:"labels"`
	}

	LabelAnswer struct {
		LabelSelector string `survey:"labelselector"`
	}
)

func PromptForNodeLabels(nodeList *v1.NodeList) string {

	var labelList []string
	err := viper.UnmarshalKey("nodeSelectorLabels", &labelList)
	if err != nil {
		common.Logger.Fatal("Error unmarshalling...")
	}
	if len(labelList) == 0 {
		labelList = defaults.Labels()
	}
	labelMap := make(map[string]string, len(labelList))
	for _, v := range labelList {
		labelMap[v] = v
	}

	nodeArray := make(map[string]string, len(nodeList.Items))

	for _, v := range nodeList.Items {
		for k, l := range v.Labels {
			if _, mok := labelMap[k]; mok {
				if _, ok := nodeArray[fmt.Sprintf("\"%s\": \"%s\"", k, l)]; !ok {
					nodeArray[fmt.Sprintf("\"%s\": \"%s\"", k, l)] = fmt.Sprintf("\"%s\": \"%s\"", k, l)
				}
			}
		}
	}

	labelSelections := []string{}
	labelSelections = append(labelSelections, "\"-none-\"")
	for _, m := range nodeArray {
		labelSelections = append(labelSelections, m)
	}

	sort.Strings(labelSelections)

	var labelSurvey = []*survey.Question{
		{
			Name: "labelselector",
			Prompt: &survey.Select{
				Message: "Choose a node selector:",
				Options: labelSelections,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	labelAnswer := &LabelAnswer{}
	if err := survey.Ask(labelSurvey, labelAnswer, opts); err != nil {
		common.Logger.WithError(err).Fatal("No node selector selected")
	}
	return labelAnswer.LabelSelector
}
