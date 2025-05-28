package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"ktrouble/common"

	"github.com/maahsome/gron"
	"github.com/muesli/termenv"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

type PodList []Pod

type Pod struct {
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	Status        string `json:"status"`
	LaunchedBy    string `json:"launchedby"`
	Service       string `json:"service"`
	ServicePort   string `json:"servicePort"`
	ContainerName string `json:"containerName"`
}

// ToJSON - Write the output as JSON
func (p *PodList) ToJSON() string {
	podJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(podJSON[:])
}

func (p *PodList) ToGRON() string {
	podJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(podJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (p *PodList) ToYAML() string {
	podYAML, err := yaml.Marshal(p)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(podYAML[:])
}

func (p *PodList) ToTEXT(to TextOptions) string {

	termFormatter := termenv.NewOutput(os.Stdout)
	bashLinks := to.BashLinks
	utilMap := to.UtilMap
	uniqIdLength := to.UniqIdLength

	buf, row := new(bytes.Buffer), make([]string, 0)
	table := tablewriter.NewWriter(buf)
	fields := []string{}

	if !to.NoHeaders {
		if len(to.Fields) > 0 {
			upperFields := fieldsToUpper(to.Fields)
			fields = append(fields, upperFields...)
		} else {
			fields = to.DefaultFields
		}
		table.SetHeader(fields)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	for _, v := range *p {
		baseTool := v.Name[0 : len(v.Name)-(uniqIdLength+1)]
		containerParam := ""
		if v.ContainerName != "" {
			containerParam = fmt.Sprintf("-c %s", v.ContainerName)
			baseTool = v.ContainerName[0 : len(v.ContainerName)-(uniqIdLength+1)]
		}
		shellText := utilMap[baseTool].ExecCommand
		if bashLinks {
			shellText = termFormatter.Hyperlink(fmt.Sprintf("<bash:kubectl -n %s exec -it %s %s -- %s>", v.Namespace, containerParam, v.Name, utilMap[baseTool].ExecCommand), utilMap[baseTool].ExecCommand)
			if v.Service != "" {
				shellText = shellText + " " + termFormatter.Hyperlink(fmt.Sprintf("<bash:kubectl -n %s port-forward svc/%s %s:443>", v.Namespace, v.Service, v.ServicePort), fmt.Sprintf("%s:443", v.ServicePort))
			}
		} else {
			if v.Service != "" {
				shellText = shellText + fmt.Sprintf(" svc/%s %s:443", v.Service, v.ServicePort)
			}
		}
		row = []string{}
		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "NAME":
				row = append(row, v.Name)
			case "NAMESPACE":
				row = append(row, v.Namespace)
			case "STATUS":
				row = append(row, v.Status)
			case "LAUNCHED_BY":
				row = append(row, v.LaunchedBy)
			case "UTILITY":
				row = append(row, baseTool)
			case "SHELL/SERVICE":
				row = append(row, shellText)
			}
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
