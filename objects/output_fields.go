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

type OutputFieldsList []OutputFields

type OutputFields struct {
	Name   string `json:"name"`
	Fields string `json:"fields"`
}

// ToJSON - Write the output as JSON
func (of *OutputFieldsList) ToJSON() string {
	ofJSON, err := json.MarshalIndent(of, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(ofJSON[:])
}

func (of *OutputFieldsList) ToGRON() string {
	ofJSON, err := json.MarshalIndent(of, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(ofJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (of *OutputFieldsList) ToYAML() string {
	ofYAML, err := yaml.Marshal(of)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(ofYAML[:])
}

func (of *OutputFieldsList) ToTEXT(to TextOptions) string {

	buf, row := new(bytes.Buffer), make([]string, 0)
	table := tablewriter.NewWriter(buf)
	fields := []string{}

	// ************************** TableWriter ******************************
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

	mapList := make(map[string]OutputFields, len(*of))
	nameList := []string{}

	for _, v := range *of {
		mapList[v.Name] = v
		nameList = append(nameList, v.Name)
	}

	for _, v := range nameList {
		row = []string{}
		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "NAME":
				row = append(row, mapList[v].Name)
			case "FIELDS":
				row = append(row, mapList[v].Fields)
			}
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
