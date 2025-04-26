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

	noHeaders := to.NoHeaders

	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		headerText := []string{"NAME", "NAMESPACE", "CLASS", "URL", "ADDRESS", "PORTS", "LAUNCHED_BY"}
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

	for _, v := range *i {
		row = []string{
			v.Name,
			v.Namespace,
			v.Class,
			v.Hosts,
			v.Address,
			v.Ports,
			v.LaunchedBy,
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}
