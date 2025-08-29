package set

import (
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type configFieldsParam struct {
	Name   string
	Fields []string
}

var fieldsP configFieldsParam

// setOutputFieldsCmd
var setOutputFieldsCmd = &cobra.Command{
	Use:     "output-fields",
	Aliases: defaults.OutputFieldsAliases,
	Short:   setOutputFieldsHelp.Short(),
	Long:    setOutputFieldsHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {

		if validateOutputFieldsParams() {
			if validateFieldNames(fieldsP) {
				if strings.ToLower(fieldsP.Fields[0]) == "default" {
					common.Logger.Warnf("Setting output fields for '%s' to default fields", fieldsP.Name)
					setOutputFieldsToDefault(fieldsP.Name)
				} else {
					common.Logger.Warnf("Setting output fields for '%s' to: %s", fieldsP.Name, strings.Join(fieldsP.Fields, ", "))
					setOutputFields(fieldsP)
				}
			}
		}
	},
}

func setOutputFields(param configFieldsParam) {
	// We have the defs in the config
	currentFieldDefs := c.OutputFieldsDefs
	newFieldDefs := objects.OutputFieldsList{}
	for _, def := range currentFieldDefs {
		if def.Name == param.Name {
			// Set the fields
			def.Fields = strings.ReplaceAll(strings.ToUpper(strings.Join(param.Fields, ",")), " ", "")
			newFieldDefs = append(newFieldDefs, def)
		} else {
			newFieldDefs = append(newFieldDefs, def)
		}
	}
	c.OutputFieldsDefs = newFieldDefs
	c.OutputFieldsMap[param.Name] = param.Fields
	viper.Set("outputFields", c.OutputFieldsDefs)
	verr := viper.WriteConfig()
	if verr != nil {
		logrus.WithError(verr).Info("Failed to write config")
	}
}

func setOutputFieldsToDefault(name string) {
	// We have the defs in the config
	currentFieldDefs := c.OutputFieldsDefs
	newFieldDefs := objects.OutputFieldsList{}
	for _, def := range currentFieldDefs {
		if def.Name == name {
			// Set to default fields
			def.Fields = strings.ReplaceAll(defaults.OutputNamesList()[name], " ", "")
			newFieldDefs = append(newFieldDefs, def)
		} else {
			newFieldDefs = append(newFieldDefs, def)
		}
	}
	c.OutputFieldsDefs = newFieldDefs
	c.OutputFieldsMap[name] = strings.Split(defaults.ValidOutputFields()[name], ",")
	viper.Set("outputFields", c.OutputFieldsDefs)
	verr := viper.WriteConfig()
	if verr != nil {
		logrus.WithError(verr).Info("Failed to write config")
	}
}

func validateFieldNames(params configFieldsParam) bool {
	if len(params.Fields) == 1 && strings.ToLower(params.Fields[0]) == "default" {
		return true
	}

	validFields, fieldList := defaults.ValidFieldNamesForOutputName(params.Name)

	for _, field := range params.Fields {
		if _, ok := validFields[strings.TrimSpace(strings.ToUpper(field))]; !ok {
			common.Logger.Warnf("Invalid field name: %s", field)
			common.Logger.Warnf("Valid fields for '%s': %s", params.Name, fieldList)
			return false
		}
	}
	return true
}

// validateOutputFieldsParams checks if the required parameters for setting output fields are provided.
func validateOutputFieldsParams() bool {
	if fieldsP.Name == "" {
		common.Logger.Warn("--output-name is required")
		return false
	}
	if len(fieldsP.Fields) == 0 {
		common.Logger.Warn("--fields is required")
		return false
	}
	return true
}

func init() {
	setCmd.AddCommand(setOutputFieldsCmd)

	setOutputFieldsCmd.Flags().StringVar(&fieldsP.Name, "output-name", "", "The name of the output to set the fields for (e.g., pod, environment, utility, etc.)")
	setOutputFieldsCmd.Flags().StringSliceVar(&fieldsP.Fields, "fields", []string{}, "Comma-separated list of fields to set for the output (e.g., name,image,hidden); 'default' to set to default fields")
}
