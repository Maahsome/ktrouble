package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	v1 "k8s.io/api/core/v1"
)

type (
	PodDetail struct {
		Name      string `survey:"podname"`
		Namespace string `survey:"namespace"`
		Deleted   string `survey:"deleted"`
	}

	NodeDetail struct {
		Name   string   `survey:"nodename"`
		Labels []string `survey:"labels"`
	}

	UtilAnswer struct {
		UtilityName string `survey:"utilityname"`
	}
	NamespaceAnswer struct {
		Namespace string `survey:"namespace"`
	}
	ServiceAccountAnswer struct {
		ServiceAccount string `survey:"serviceaccount"`
	}
	ResourceSizeAnswer struct {
		ResourceSize string `survey:"resourcesize"`
	}
	PodAnswer struct {
		Pod string `survey:"podname"`
	}
	LabelAnswer struct {
		LabelSelector string `survey:"labelselector"`
	}
)

func askForUtility(utils map[string]UtilityPod) string {

	var utilArray []string
	for k := range utils {
		utilArray = append(utilArray, k)
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
		logrus.WithError(err).Fatal("No utility selected")
	}
	return utilAnswer.UtilityName
}

func askForNamespace(nssList *v1.NamespaceList) string {

	var nsArray []string
	for _, v := range nssList.Items {
		nsArray = append(nsArray, v.Name)
	}
	sort.Strings(nsArray)

	var nsSurvey = []*survey.Question{
		{
			Name: "namespace",
			Prompt: &survey.Select{
				Message: "Choose a namespace to create the pod in:",
				Options: nsArray,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	nsAnswer := &NamespaceAnswer{}
	if err := survey.Ask(nsSurvey, nsAnswer, opts); err != nil {
		logrus.WithError(err).Fatal("No namespace selected")
	}
	return nsAnswer.Namespace
}

func askForServiceAccount(sasList *v1.ServiceAccountList) string {

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
		logrus.WithError(err).Fatal("No service account selected")
	}
	return saAnswer.ServiceAccount
}

func askForResourceSize() map[string]string {

	rsArray := []string{
		"Small",
		"Medium",
		"Large",
		"XLarge",
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
		logrus.WithError(err).Fatal("No resource size selected")
	}

	switch rsAnswer.ResourceSize {
	case "Small":
		return map[string]string{
			"limitsCpu":  "250m",
			"limitsMem":  "2Gi",
			"requestCpu": "100m",
			"requestMem": "512Mi",
		}
	case "Medium":
		return map[string]string{
			"limitsCpu":  "500m",
			"limitsMem":  "4Gi",
			"requestCpu": "200m",
			"requestMem": "1Gi",
		}
	case "Large":
		return map[string]string{
			"limitsCpu":  "1000m",
			"limitsMem":  "8Gi",
			"requestCpu": "500m",
			"requestMem": "2Gi",
		}
	case "XLarge":
		return map[string]string{
			"limitsCpu":  "10000m",
			"limitsMem":  "30Gi",
			"requestCpu": "6000m",
			"requestMem": "2Gi",
		}
	}

	return map[string]string{
		"limitsCpu":  "250m",
		"limitsMem":  "2Gi",
		"requestCpu": "100m",
		"requestMem": "512Mi",
	}
}

func askForPod(podList *v1.PodList) PodDetail {

	podArray := make(map[string]PodDetail, len(podList.Items))

	for _, v := range podList.Items {
		deleting := "false"
		if v.DeletionTimestamp != nil {
			deleting = "true"
		}
		podArray[fmt.Sprintf("%s/%s", v.Namespace, v.Name)] = PodDetail{
			Name:      v.Name,
			Namespace: v.Namespace,
			Deleted:   deleting,
		}
	}

	podNames := []string{}
	for _, m := range podArray {
		if m.Deleted == "false" {
			podNames = append(podNames, fmt.Sprintf("%s/%s", m.Namespace, m.Name))
		}
	}

	var podSurvey = []*survey.Question{
		{
			Name: "podname",
			Prompt: &survey.Select{
				Message: "Choose a pod to delete:",
				Options: podNames,
			},
		},
	}

	opts := survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)

	podAnswer := &PodAnswer{}
	if err := survey.Ask(podSurvey, podAnswer, opts); err != nil {
		logrus.WithError(err).Fatal("No pod selected")
	}
	return podArray[podAnswer.Pod]
}

func askForNodeLabels(nodeList *v1.NodeList) string {

	var labelList []string
	err := viper.UnmarshalKey("nodeSelectorLabels", &labelList)
	if err != nil {
		logrus.Fatal("Error unmarshalling...")
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
		logrus.WithError(err).Fatal("No node selector selected")
	}
	return labelAnswer.LabelSelector
}
