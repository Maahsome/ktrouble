package objects

import (
	"bytes"
	"encoding/json"
	"strings"

	"ktrouble/common"

	"github.com/maahsome/gron"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

type ResourceSizeList []ResourceSize

type ResourceSize struct {
	Name       string `json:"name"`
	LimitsCPU  string `json:"limitsCpu"`
	LimitsMEM  string `json:"limitsMem"`
	RequestCPU string `json:"requestCpu"`
	RequestMEM string `json:"requestMem"`
}

// ToJSON - Write the output as JSON
func (rs *ResourceSizeList) ToJSON() string {
	rsJSON, err := json.MarshalIndent(rs, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(rsJSON[:])
}

func (rs *ResourceSizeList) ToGRON() string {
	rsJSON, err := json.MarshalIndent(rs, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(rsJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (rs *ResourceSizeList) ToYAML() string {
	rsYAML, err := yaml.Marshal(rs)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(rsYAML[:])
}

func (rs *ResourceSizeList) ToTEXT(noHeaders bool) string {
	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		headerText := []string{"NAME", "CPU_LIMIT", "MEM_LIMIT", "CPU_REQUEST", "MEM_REQUEST"}
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

	mapList := make(map[string]ResourceSize, len(*rs))
	nameList := []string{}

	for _, v := range *rs {
		mapList[v.Name] = v
		nameList = append(nameList, v.Name)
	}

	for _, v := range nameList {
		row = []string{
			mapList[v].Name,
			mapList[v].LimitsCPU,
			mapList[v].LimitsMEM,
			mapList[v].RequestCPU,
			mapList[v].RequestMEM,
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
