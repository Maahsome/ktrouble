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

type ServiceAccountList struct {
	ServiceAccount []string `json:"serviceAccount"`
}

// ToJSON - Write the output as JSON
func (sa *ServiceAccountList) ToJSON() string {
	saJSON, err := json.MarshalIndent(sa, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(saJSON[:])
}

func (sa *ServiceAccountList) ToGRON() string {
	saJSON, err := json.MarshalIndent(sa, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(saJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return subValues.String()
}

func (sa *ServiceAccountList) ToYAML() string {
	saYAML, err := yaml.Marshal(sa)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(saYAML[:])
}

func (sa *ServiceAccountList) ToTEXT(to TextOptions) string {

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

	for _, v := range sa.ServiceAccount {
		row = []string{}
		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "SERVICE_ACCOUNT":
				row = append(row, v)
			}
		}
		table.Append(row)
	}
	table.Render()

	return buf.String()
}
