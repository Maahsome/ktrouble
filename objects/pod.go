package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"ktrouble/common"

	"github.com/maahsome/gron"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

type PodList []Pod

type Pod struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Status    string `json:"status"`
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

func (p *PodList) ToTEXT(noHeaders bool, showExec bool, utilMap map[string]UtilityPod, uniqIdLength int) string {
	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		headerText := []string{"NAME", "NAMESPACE", "STATUS"}
		if showExec {
			headerText = append(headerText, "EXEC")
		}
		table.SetHeader(headerText)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	displayOptions := 0
	if showExec {
		displayOptions += 1
	}

	for _, v := range *p {
		switch displayOptions {
		case 0:
			row = []string{
				v.Name,
				v.Namespace,
				v.Status,
			}
		case 1:
			baseTool := v.Name[0 : len(v.Name)-(uniqIdLength+1)]
			row = []string{
				v.Name,
				v.Namespace,
				v.Status,
				fmt.Sprintf("<bash:kubectl -n %s exec -it %s -- %s>", v.Namespace, v.Name, utilMap[baseTool].ExecCommand),
			}
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
