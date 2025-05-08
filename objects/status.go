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

type StatusList []Status

type Status struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Exclude string `json:"exclude"`
}

// ToJSON - Write the output as JSON
func (s *StatusList) ToJSON() string {
	sJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(sJSON[:])
}

func (s *StatusList) ToGRON() string {
	sJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(sJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return subValues.String()
}

func (s *StatusList) ToYAML() string {
	sYAML, err := yaml.Marshal(s)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(sYAML[:])
}

func (s *StatusList) ToTEXT(to TextOptions) string {

	buf := new(bytes.Buffer)
	var row []string
	table := tablewriter.NewWriter(buf)
	fields := []string{}

	// ************************** TableWriter ******************************
	if !to.NoHeaders {
		if len(to.Fields) > 0 {
			upperFields := fieldsToUpper(to.Fields)
			fields = append(fields, upperFields...)
		} else {
			fields = []string{"NAME", "STATUS", "PUSH_EXCLUDE"}
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

	for _, v := range *s {
		row = []string{}
		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "NAME":
				row = append(row, v.Name)
			case "STATUS":
				row = append(row, v.Status)
			case "EXCLUDE":
				row = append(row, v.Exclude)
			}
		}
		table.Append(row)
	}
	table.Render()

	return buf.String()
}
