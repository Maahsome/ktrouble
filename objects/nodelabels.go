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

type NodeLabels []string

// ToJSON - Write the output as JSON
func (nl *NodeLabels) ToJSON() string {
	podJSON, err := json.MarshalIndent(nl, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(podJSON[:])
}

func (nl *NodeLabels) ToGRON() string {
	podJSON, err := json.MarshalIndent(nl, "", "  ")
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

func (nl *NodeLabels) ToYAML() string {
	podYAML, err := yaml.Marshal(nl)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(podYAML[:])
}

func (nl *NodeLabels) ToTEXT(to TextOptions) string {

	noHeaders := to.NoHeaders

	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		headerText := []string{"LABEL"}
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

	for _, v := range *nl {
		row = []string{
			v,
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
