package template

import (
	"bytes"
	"fmt"
	"ktrouble/common"
	"os"
	"strings"
	"text/template"
)

type TemplateProcessor interface {
	RenderTemplate(tc *TemplateConfig) string
}

type templateProcessor struct {
	Tpl string
}

type TemplateConfig struct {
	Parameters map[string]string
	Secrets    []string
	ConfigMaps []string
}

// takes in some text and pads it with spaces even if it is multiline
func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)

	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

var applicationsTemplateFuncs = template.FuncMap{
	"indent":    indent,
	"hasPrefix": strings.HasPrefix,
}

func New(templateFile string) TemplateProcessor {
	home, herr := os.UserHomeDir()
	if herr != nil {
		common.Logger.Error("failed to fetch home directory")
	}
	tmplDir := fmt.Sprintf("%s/.config/ktrouble/templates", home)

	templatePath := fmt.Sprintf("%s/%s", tmplDir, templateFile)
	fileData, err := os.ReadFile(templatePath)
	if err != nil {
		common.Logger.WithError(err).Fatal("failed to read the specified template file: %s")
	}
	tpl := string(fileData)

	return &templateProcessor{
		Tpl: tpl,
	}
}

func (t *templateProcessor) RenderTemplate(tc *TemplateConfig) string {

	applicationsTemplate := template.Must(
		template.New("pod.yaml.tpl").Funcs(applicationsTemplateFuncs).Parse(t.Tpl))
	var tpl bytes.Buffer
	if err := applicationsTemplate.Execute(&tpl, tc); err != nil {
		common.Logger.WithError(err).Error("unable to generate the template data")
	}

	return tpl.String()

}
