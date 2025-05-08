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

type EphemeralSleepList []EphemeralSleep

type EphemeralSleep struct {
	Name    string `json:"name"`
	Seconds string `json:"seconds"`
}

// ToJSON - Write the output as JSON
func (es *EphemeralSleepList) ToJSON() string {
	rsJSON, err := json.MarshalIndent(es, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(rsJSON[:])
}

func (es *EphemeralSleepList) ToGRON() string {
	esJSON, err := json.MarshalIndent(es, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(esJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (es *EphemeralSleepList) ToYAML() string {
	rsYAML, err := yaml.Marshal(es)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(rsYAML[:])
}

func (es *EphemeralSleepList) ToTEXT(to TextOptions) string {

	buf, row := new(bytes.Buffer), make([]string, 0)
	table := tablewriter.NewWriter(buf)
	fields := []string{}
	if !to.NoHeaders {
		if len(to.Fields) > 0 {
			fields = append(fields, to.Fields...)
		} else {
			fields = []string{"NAME", "SECONDS"}
		}
		if len(to.AdditionalFields) > 0 {
			fields = append(fields, to.AdditionalFields...)
		}
		table.SetHeader(fields)
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

	mapList := make(map[string]EphemeralSleep, len(*es))
	nameList := []string{}

	for _, v := range *es {
		mapList[v.Name] = v
		nameList = append(nameList, v.Name)
	}

	for _, v := range nameList {
		row = []string{}
		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "NAME":
				row = append(row, mapList[v].Name)
			case "SECONDS":
				row = append(row, mapList[v].Seconds)
			}
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
