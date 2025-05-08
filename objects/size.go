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

func (rs *ResourceSizeList) ToTEXT(to TextOptions) string {

	buf, row := new(bytes.Buffer), make([]string, 0)
	table := tablewriter.NewWriter(buf)
	fields := []string{}

	// ************************** TableWriter ******************************
	if !to.NoHeaders {
		if len(to.Fields) > 0 {
			upperFields := fieldsToUpper(to.Fields)
			fields = append(fields, upperFields...)
		} else {
			fields = []string{"NAME", "CPU_LIMIT", "MEM_LIMIT", "CPU_REQUEST", "MEM_REQUEST"}
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

	mapList := make(map[string]ResourceSize, len(*rs))
	nameList := []string{}

	for _, v := range *rs {
		mapList[v.Name] = v
		nameList = append(nameList, v.Name)
	}

	for _, v := range nameList {
		row = []string{}
		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "NAME":
				row = append(row, mapList[v].Name)
			case "CPU_LIMIT":
				row = append(row, mapList[v].LimitsCPU)
			case "MEM_LIMIT":
				row = append(row, mapList[v].LimitsMEM)
			case "CPU_REQUEST":
				row = append(row, mapList[v].RequestCPU)
			case "MEM_REQUEST":
				row = append(row, mapList[v].RequestMEM)
			}
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
