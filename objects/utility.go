package objects

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"

	"ktrouble/common"

	"github.com/maahsome/gron"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

type UtilityPodList []UtilityPod

type UtilityPod struct {
	Name             string `json:"name"`
	Repository       string `json:"repository"`
	ExecCommand      string `json:"execcommand"`
	Source           string `json:"source"`
	ExcludeFromShare bool   `json:"excludefromshare"`
}

// ToJSON - Write the output as JSON
func (up *UtilityPodList) ToJSON() string {
	podJSON, err := json.MarshalIndent(up, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(podJSON[:])
}

func (up *UtilityPodList) ToGRON() string {
	podJSON, err := json.MarshalIndent(up, "", "  ")
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

func (up *UtilityPodList) ToYAML() string {
	podYAML, err := yaml.Marshal(up)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(podYAML[:])
}

func (up *UtilityPodList) ToTEXT(to TextOptions) string {

	noHeaders := to.NoHeaders

	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		headerText := []string{"NAME", "REPOSITORY", "EXEC"}
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

	mapList := make(map[string]UtilityPod, len(*up))
	nameList := []string{}

	for _, v := range *up {
		mapList[v.Name] = v
		nameList = append(nameList, v.Name)
	}

	sort.Strings(nameList)

	for _, v := range nameList {
		row = []string{
			mapList[v].Name,
			mapList[v].Repository,
			mapList[v].ExecCommand,
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
