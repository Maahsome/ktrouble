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

type ServiceList []Service

// NAME                                      TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)
type Service struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	ServiceType string `json:"type"`
	ClusterIP   string `json:"clusterip"`
	ExternalIP  string `json:"externalip"`
	Ports       string `json:"ports"`
	LaunchedBy  string `json:"launchedby"`
}

// ToJSON - Write the output as JSON
func (s *ServiceList) ToJSON() string {
	podJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(podJSON[:])
}

func (s *ServiceList) ToGRON() string {
	podJSON, err := json.MarshalIndent(s, "", "  ")
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

func (s *ServiceList) ToYAML() string {
	podYAML, err := yaml.Marshal(s)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(podYAML[:])
}

func (s *ServiceList) ToTEXT(to TextOptions) string {

	buf, row := new(bytes.Buffer), make([]string, 0)
	table := tablewriter.NewWriter(buf)
	fields := []string{}

	// ************************** TableWriter ******************************
	if !to.NoHeaders {
		if len(to.Fields) > 0 {
			upperFields := fieldsToUpper(to.Fields)
			fields = append(fields, upperFields...)
		} else {
			fields = []string{"NAME", "NAMESPACE", "TYPE", "CLUSTER-IP", "EXTERNAL-IP", "PORT(S)", "LAUNCHED_BY"}
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
			case "NAMESPACE":
				row = append(row, v.Namespace)
			case "TYPE":
				row = append(row, v.ServiceType)
			case "CLUSTER-IP":
				row = append(row, v.ClusterIP)
			case "EXTERNAL-IP":
				row = append(row, v.ExternalIP)
			case "PORT(S)":
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
