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

type NodeList struct {
	Node []string `json:"node"`
}

// ToJSON - Write the output as JSON
func (n *NodeList) ToJSON() string {
	nJSON, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(nJSON[:])
}

func (n *NodeList) ToGRON() string {
	nJSON, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(nJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return subValues.String()
}

func (n *NodeList) ToYAML() string {
	nYAML, err := yaml.Marshal(n)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(nYAML[:])
}

func (n *NodeList) ToTEXT(to TextOptions) string {

	buf := new(bytes.Buffer)
	var row []string
	table := tablewriter.NewWriter(buf)
	fields := []string{}

	if !to.NoHeaders {
		if len(to.Fields) > 0 {
			upperFields := fieldsToUpper(to.Fields)
			fields = append(fields, upperFields...)
		} else {
			fields = []string{"NODE"}
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

	for _, v := range n.Node {
		row = []string{}
		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "NODE":
				row = append(row, v)
			}
		}
		table.Append(row)
	}
	table.Render()

	return buf.String()
}
