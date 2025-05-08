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

type IngressList []Ingress

type Ingress struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Class      string `json:"class"`
	Hosts      string `json:"hosts"`
	Address    string `json:"address"`
	Ports      string `json:"ports"`
	LaunchedBy string `json:"launchedby"`
}

// ToJSON - Write the output as JSON
func (i *IngressList) ToJSON() string {
	podJSON, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(podJSON[:])
}

func (i *IngressList) ToGRON() string {
	podJSON, err := json.MarshalIndent(i, "", "  ")
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

func (i *IngressList) ToYAML() string {
	podYAML, err := yaml.Marshal(i)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(podYAML[:])
}

func (i *IngressList) ToTEXT(to TextOptions) string {

	buf, row := new(bytes.Buffer), make([]string, 0)
	table := tablewriter.NewWriter(buf)
	fields := []string{}
	if !to.NoHeaders {
		if len(to.Fields) > 0 {
			upperFields := fieldsToUpper(to.Fields)
			fields = append(fields, upperFields...)
		} else {
			fields = []string{"NAME", "NAMESPACE", "CLASS", "URL", "ADDRESS", "PORTS", "LAUNCHED_BY"}
		}
		if len(to.AdditionalFields) > 0 {
			fields = append(fields, to.AdditionalFields...)
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

	for _, v := range *i {
		row = []string{}
		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "NAME":
				row = append(row, v.Name)
			case "NAMESPACE":
				row = append(row, v.Namespace)
			case "CLASS":
				row = append(row, v.Class)
			case "URL":
				row = append(row, v.Hosts)
			case "ADDRESS":
				row = append(row, v.Address)
			case "PORTS":
				row = append(row, v.Ports)
			case "LAUNCHED_BY":
				row = append(row, v.LaunchedBy)
			}
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
