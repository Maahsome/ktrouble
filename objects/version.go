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

type Version struct {
	SemVer    string `json:"SemVer"`
	GitCommit string `json:"GitCommit"`
	BuildDate string `json:"BuildDate"`
	GitRef    string `json:"GitRef"`
}

// ToJSON - Write the output as JSON
func (v *Version) ToJSON() string {
	versionJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(versionJSON[:])
}

func (v *Version) ToGRON() string {
	versionJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(versionJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		logrus.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return subValues.String()
}

func (v *Version) ToYAML() string {
	versionYAML, err := yaml.Marshal(v)
	if err != nil {
		logrus.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(versionYAML[:])
}

func (v *Version) ToTEXT(to TextOptions) string {

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
			fields = []string{"SEMVER", "BUILD_DATE", "GIT_COMMIT", "GIT_REF"}
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

	row = []string{}
	for _, f := range fields {
		switch strings.ToUpper(f) {
		case "SEMVER":
			row = append(row, v.SemVer)
		case "BUILD_DATE":
			row = append(row, v.BuildDate)
		case "GIT_COMMIT":
			row = append(row, v.GitCommit)
		case "GIT_REF":
			row = append(row, v.GitRef)
		}
	}
	table.Append(row)

	table.Render()

	return buf.String()
}
