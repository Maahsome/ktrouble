package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"ktrouble/common"

	"github.com/maahsome/gron"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type EnvironmentConfig struct {
	Environments EnvironmentList `json:"environments"`
}

type EnvironmentList []Environment

type Environment struct {
	Name             string `json:"name"`
	Repository       string `json:"repository"`
	ExcludeFromShare bool   `json:"excludefromshare"`
	Hidden           bool   `json:"hidden"`
	RemoveUpstream   bool   `json:"removeupstream"`
}

func GetEnvironmentMap(envDefs EnvironmentList) map[string]Environment {
	envMap := make(map[string]Environment)

	for _, e := range envDefs {
		envMap[e.Name] = Environment{
			Name:             e.Name,
			Repository:       e.Repository,
			ExcludeFromShare: e.ExcludeFromShare,
		}
	}
	return envMap
}

func RemoveEnvIndex(s EnvironmentList, index int) EnvironmentList {
	ret := make(EnvironmentList, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func MigrateLocalEnvironments(u EnvironmentList, toVer string) bool {

	viper.Set("environments", u)
	verr := viper.WriteConfig()
	if verr != nil {
		common.Logger.WithError(verr).Info("Failed to write config")
		return false
	}

	return true
}

func EnvironmentsExist(envMap map[string]Environment, envNames []string) bool {
	for _, envName := range envNames {
		common.Logger.Tracef("Checking environment: %s", envName)
		if _, ok := envMap[envName]; !ok {
			common.Logger.Tracef("Environment %s not found: %t", envName, ok)
			return false
		}
	}
	return true
}

// ToJSON - Write the output as JSON
func (en *EnvironmentList) ToJSON() string {
	envJSON, err := json.MarshalIndent(en, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(envJSON[:])
}

func (en *EnvironmentList) ToGRON() string {
	envJSON, err := json.MarshalIndent(en, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(envJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		common.Logger.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (en *EnvironmentList) ToYAML() string {
	envYAML, err := yaml.Marshal(en)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(envYAML[:])
}

func (en *EnvironmentList) ToTEXT(to TextOptions) string {

	buf, row := new(bytes.Buffer), make([]string, 0)
	table := tablewriter.NewWriter(buf)
	fields := []string{}
	if !to.NoHeaders {
		if len(to.Fields) > 0 {
			upperFields := fieldsToUpper(to.Fields)
			fields = append(fields, upperFields...)
		} else {
			fields = to.DefaultFields
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

	mapList := make(map[string]Environment, len(*en))
	nameList := []string{}

	for _, v := range *en {
		mapList[v.Name] = v
		nameList = append(nameList, v.Name)
	}

	sort.Strings(nameList)

	for _, v := range nameList {
		row = []string{}

		for _, f := range fields {
			switch strings.ToUpper(f) {
			case "NAME":
				row = append(row, mapList[v].Name)
			case "REPOSITORY":
				row = append(row, mapList[v].Repository)
			case "EXCLUDED":
				row = append(row, fmt.Sprintf("%t", mapList[v].ExcludeFromShare))
			case "HIDDEN":
				row = append(row, fmt.Sprintf("%t", mapList[v].Hidden))
			case "REMOVE_UPSTREAM":
				row = append(row, fmt.Sprintf("%t", mapList[v].RemoveUpstream))
			}
		}
		if !mapList[v].Hidden || to.ShowHidden {
			table.Append(row)
		}
	}

	table.Render()

	return buf.String()

}
