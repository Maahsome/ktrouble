package config

import (
	"fmt"
	"ktrouble/gitupstream"
	"ktrouble/kubernetes"
	"ktrouble/objects"
	"strings"
)

type (
	Config struct {
		VersionDetail       objects.Version
		VersionJSON         string
		OutputFormat        string
		FormatOverridden    bool
		NoHeaders           bool
		CACert              string
		CABundle            string
		LogLevel            string
		LogFile             string
		Namespace           string
		EnableBashLinks     bool
		UniqIdLength        int
		ShowHidden          bool
		UtilMap             map[string]objects.UtilityPod
		UtilDefs            objects.UtilityPodList
		SizeMap             map[string]objects.ResourceSize
		SizeDefs            objects.ResourceSizeList
		NodeSelectorLabels  []string
		Client              kubernetes.KubernetesClient
		Fields              []string
		GitUser             string
		GitToken            string
		GitUpstream         gitupstream.GitUpstream
		GitURL              string
		PromptForSecrets    bool
		PromptForConfigMaps bool
		TemplateFile        string
	}

	Outputtable interface {
		ToJSON() string
		ToYAML() string
		ToGRON() string
		ToTEXT(to objects.TextOptions) string
	}
)

func (c *Config) outputData(data Outputtable, to objects.TextOptions) string {
	switch strings.ToLower(c.OutputFormat) {
	case "raw":
		return fmt.Sprintf("%#v", data)
	case "json":
		return data.ToJSON()
	case "gron":
		return data.ToGRON()
	case "yaml":
		return data.ToYAML()
	case "text", "table":
		return data.ToTEXT(to)
	default:
		return data.ToTEXT(to)
	}
}

func (c *Config) OutputData(data Outputtable, to objects.TextOptions) {
	fmt.Println(c.outputData(data, to))
}
