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

type UtilityPodList []UtilityPod
type UtilityPodListV0 []UtilityPodV0

type UtilityPod struct {
	Name              string   `json:"name"`
	Repository        string   `json:"repository"`
	ExecCommand       string   `json:"execcommand"`
	Source            string   `json:"source"`
	RequireSecrets    bool     `json:"requiresecrets"`
	RequireConfigmaps bool     `json:"requireconfigmaps"`
	ExcludeFromShare  bool     `json:"excludefromshare"`
	Hidden            bool     `json:"hidden"`
	Hint              string   `json:"hint"`
	RemoveUpstream    bool     `json:"removeupstream"`
	Environments      []string `json:"environments"`
}

type UtilityPodV0 struct {
	Name              string `json:"name"`
	Repository        string `json:"repository"`
	ExecCommand       string `json:"execcommand"`
	Source            string `json:"source"`
	RequireSecrets    bool   `json:"requiresecrets"`
	RequireConfigmaps bool   `json:"requireconfigmaps"`
	ExcludeFromShare  bool   `json:"excludefromshare"`
	Hidden            bool   `json:"hidden"`
	Hint              string `json:"hint"`
	Version           string `json:"version"`
}

func GetUtilityMap(utilDefs UtilityPodList, envMap map[string]Environment) map[string]UtilityPod {
	utilMap := make(map[string]UtilityPod)
	for _, v := range utilDefs {
		if v.Environments == nil {
			utilMap[v.Name] = UtilityPod{
				Name:              v.Name,
				Repository:        v.Repository,
				ExecCommand:       v.ExecCommand,
				RequireSecrets:    v.RequireSecrets,
				RequireConfigmaps: v.RequireConfigmaps,
				Hint:              v.Hint,
			}
		} else {
			for _, env := range v.Environments {
				utilMap[fmt.Sprintf("%s/%s", env, v.Name)] = UtilityPod{
					Name:              v.Name,
					Repository:        fmt.Sprintf("%s/%s", envMap[env].Repository, v.Repository),
					ExecCommand:       v.ExecCommand,
					RequireSecrets:    v.RequireSecrets,
					RequireConfigmaps: v.RequireConfigmaps,
					Hint:              v.Hint,
				}
			}
		}
	}

	return utilMap
}

func RemoveUtilIndex(s UtilityPodList, index int) UtilityPodList {
	ret := make(UtilityPodList, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func MigrateLocalUtility(u UtilityPodList, toVer string) bool {

	viper.Set("utilityDefinitions", u)
	verr := viper.WriteConfig()
	if verr != nil {
		common.Logger.WithError(verr).Info("Failed to write config")
		return false
	}
	return true
}

// ToJSON - Write the output as JSON
func (up *UtilityPodList) ToJSON() string {
	podJSON, err := json.MarshalIndent(up, "", "  ")
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(podJSON[:])
}

func (up *UtilityPodList) ToGRON() string {
	podJSON, err := json.MarshalIndent(up, "", "  ")
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

func (up *UtilityPodList) ToYAML() string {
	podYAML, err := yaml.Marshal(up)
	if err != nil {
		common.Logger.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(podYAML[:])
}

func (up *UtilityPodList) ToTEXT(to TextOptions) string {

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
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	mapList := make(map[string]UtilityPod, len(*up))
	nameList := []string{}

	for _, v := range *up {
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
			case "EXEC":
				row = append(row, mapList[v].ExecCommand)
			case "HIDDEN":
				row = append(row, fmt.Sprintf("%t", mapList[v].Hidden))
			case "EXCLUDED":
				row = append(row, fmt.Sprintf("%t", mapList[v].ExcludeFromShare))
			case "REMOVE_UPSTREAM":
				row = append(row, fmt.Sprintf("%t", mapList[v].RemoveUpstream))
			case "SOURCE":
				row = append(row, mapList[v].Source)
			case "REQUIRESECRETS":
				row = append(row, fmt.Sprintf("%t", mapList[v].RequireSecrets))
			case "REQUIRECONFIGMAPS":
				row = append(row, fmt.Sprintf("%t", mapList[v].RequireConfigmaps))
			case "ENVIRONMENTS":
				row = append(row, strings.Join(mapList[v].Environments, ","))
			case "HINT":
				row = append(row, mapList[v].Hint)
			}
		}
		if !mapList[v].Hidden || to.ShowHidden {
			table.Append(row)
		}
	}

	table.Render()

	return buf.String()

}
