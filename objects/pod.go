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
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Status      string `json:"status"`
	LaunchedBy  string `json:"launchedby"`
	Service     string `json:"service"`
	ServicePort string `json:"servicePort"`
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
	noHeaders := to.NoHeaders
	bashLinks := to.BashLinks
	utilMap := to.UtilMap
	uniqIdLength := to.UniqIdLength

	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		headerText := []string{"NAME", "NAMESPACE", "STATUS", "LAUNCHED_BY", "SHELL/SERVICE"}
		table.SetHeader(headerText)
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
		shellText := utilMap[baseTool].ExecCommand
		if bashLinks {
			shellText = termFormatter.Hyperlink(fmt.Sprintf("<bash:kubectl -n %s exec -it %s -- %s>", v.Namespace, v.Name, utilMap[baseTool].ExecCommand), utilMap[baseTool].ExecCommand)
			if v.Service != "" {
				shellText = shellText + " " + termFormatter.Hyperlink(fmt.Sprintf("<bash:kubectl -n %s port-forward svc/%s %s:443>", v.Namespace, v.Service, v.ServicePort), fmt.Sprintf("%s:443", v.ServicePort))
			}
		} else {
			if v.Service != "" {
				shellText = shellText + fmt.Sprintf(" svc/%s %s:443", v.Service, v.ServicePort)
			}
		}
		row = []string{
			v.Name,
			v.Namespace,
			v.Status,
			v.LaunchedBy,
			shellText,
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
