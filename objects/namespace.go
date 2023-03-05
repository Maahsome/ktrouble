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

type NamespaceList struct {
	Namespace []string `json:"namespace"`
}

// ToJSON - Write the output as JSON
func (n *NamespaceList) ToJSON() string {
	nsJSON, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(nsJSON[:])
}

func (n *NamespaceList) ToGRON() string {
	nsJSON, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(nsJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return subValues.String()
}

func (n *NamespaceList) ToYAML() string {
	nsYAML, err := yaml.Marshal(n)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(nsYAML[:])
}

func (n *NamespaceList) ToTEXT(to TextOptions) string {

	noHeaders := to.NoHeaders

	buf := new(bytes.Buffer)
	var row []string

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		table.SetHeader([]string{"NAMESPACE"})
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

	for _, v := range n.Namespace {
		row = []string{
			v,
		}
		table.Append(row)
	}
	table.Render()

	return buf.String()
}
