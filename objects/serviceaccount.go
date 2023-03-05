package objects

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/maahsome/gron"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type ServiceAccountList struct {
	ServiceAccount []string `json:"serviceAccount"`
}

// ToJSON - Write the output as JSON
func (sa *ServiceAccountList) ToJSON() string {
	saJSON, err := json.MarshalIndent(sa, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(saJSON[:])
}

func (sa *ServiceAccountList) ToGRON() string {
	saJSON, err := json.MarshalIndent(sa, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(saJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		logrus.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return subValues.String()
}

func (sa *ServiceAccountList) ToYAML() string {
	saYAML, err := yaml.Marshal(sa)
	if err != nil {
		logrus.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(saYAML[:])
}

func (sa *ServiceAccountList) ToTEXT(noHeaders bool) string {
	buf := new(bytes.Buffer)
	var row []string

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		table.SetHeader([]string{"SERVICEACCOUNT"})
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

	for _, v := range sa.ServiceAccount {
		row = []string{
			v,
		}
		table.Append(row)
	}
	table.Render()

	return buf.String()
}
