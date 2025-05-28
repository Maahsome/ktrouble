package cmd

import (
	"bytes"
	"fmt"
	"ktrouble/common"
	"ktrouble/internal"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

type statusParams struct {
	Environments bool
}

var statusP statusParams

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: statusHelp.Short(),
	Long:  statusHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {

		if !c.GitUpstream.VersionDirectoryExists(fmt.Sprintf("v%d", c.Semver.Major)) {
			if c.Semver.Major == 0 {
				common.Logger.Error("The repository is not initialized, please create the repository usign git before running this command")
				return
			} else {
				common.Logger.Error("The version directory in the repository does not exist.  Please run 'ktrouble migrate' to migrate data to the new version")
				common.Logger.Error("The existing data will remain in the old version directory")
				return
			}
		}
		if statusP.Environments {
			status := EnvironmentDefinitionStatus()
			c.OutputData(&status, objects.TextOptions{
				NoHeaders:     c.NoHeaders,
				Fields:        c.Fields,
				DefaultFields: c.OutputFieldsMap["status"],
			})
		} else {
			status := UtilityDefinitionStatus()
			c.OutputData(&status, objects.TextOptions{
				NoHeaders:     c.NoHeaders,
				Fields:        c.Fields,
				DefaultFields: c.OutputFieldsMap["status"],
			})
		}
	},
}

func EnvironmentDefinitionStatus() objects.StatusList {
	status := objects.StatusList{}

	remoteEnvDefs, remoteEnvDefsMap := c.GitUpstream.GetUpstreamEnvDefs()
	for _, l := range c.EnvDefs {
		if r, ok := remoteEnvDefsMap[l.Name]; !ok {
			status = append(status, objects.Status{
				Name:    l.Name,
				Status:  "only local",
				Exclude: fmt.Sprintf("%t", l.ExcludeFromShare),
			})
		} else {
			s := compareEnvDefs(l, r)
			status = append(status, objects.Status{
				Name:    l.Name,
				Status:  s,
				Exclude: fmt.Sprintf("%t", l.ExcludeFromShare),
			})
		}
	}
	for _, r := range remoteEnvDefs {
		if _, ok := c.EnvMap[r.Name]; !ok {
			status = append(status, objects.Status{
				Name:    r.Name,
				Status:  "only remote",
				Exclude: "",
			})
		}
	}
	return status
}

func UtilityDefinitionStatus() objects.StatusList {
	status := objects.StatusList{}

	remoteDefs, remoteDefsMap, remoteEnvDefs, _ := c.GitUpstream.GetUpstreamDefs()

	for _, l := range c.UtilDefs {
		if r, ok := remoteDefsMap[l.Name]; !ok {
			status = append(status, objects.Status{
				Name:    l.Name,
				Status:  "only local",
				Exclude: fmt.Sprintf("%t", l.ExcludeFromShare),
			})
		} else {
			s := compareDefs(l, r)
			status = append(status, objects.Status{
				Name:    l.Name,
				Status:  s,
				Exclude: fmt.Sprintf("%t", l.ExcludeFromShare),
			})
		}
	}
	for _, r := range remoteDefs {
		if _, ok := c.UtilMap[r.Name]; !ok {
			status = append(status, objects.Status{
				Name:    r.Name,
				Status:  "only remote",
				Exclude: "",
			})
		}
	}

	// Check to see if there are ONLY REMOTE environment definitions
	onlyRemoteEnv := false
	for _, r := range remoteEnvDefs {
		if _, ok := c.EnvMap[r.Name]; !ok {
			onlyRemoteEnv = true
		}
	}
	if onlyRemoteEnv {
		common.Logger.Warn("There are remote environment definitions that are not in the local config.  Please run 'ktrouble status --env' and 'ktrouble pull --env' to add them to the local config")
	}
	return status
}

func compareEnvDefs(local objects.Environment, remote objects.Environment) string {
	localYaml := "EnvironmentDefinition: \n  - " + unmarshallEnvDefinition(local)
	remoteYaml := "EnvironmentDefinition: \n  - " + unmarshallEnvDefinition(remote)

	origReader := bytes.NewReader([]byte(localYaml))
	origBuffer := new(bytes.Buffer)
	editReader := bytes.NewReader([]byte(remoteYaml))
	editBuffer := new(bytes.Buffer)

	serr := internal.SortYAML(origReader, origBuffer, 2)
	if serr != nil {
		common.Logger.WithError(serr).Error("Error sorting original yaml")
	}
	serr = internal.SortYAML(editReader, editBuffer, 2)
	if serr != nil {
		common.Logger.WithError(serr).Error("Error sorting edit yaml")
	}

	sortedOriginal := origBuffer.String()
	sortedEdit := editBuffer.String()

	if sortedOriginal == sortedEdit {
		return "same"
	}
	return "different"
}

func compareDefs(local objects.UtilityPod, remote objects.UtilityPod) string {

	localYaml := "UtilityDefinition: \n  - " + unmarshallUtilityDefinition(local)
	remoteYaml := "UtilityDefinition: \n  - " + unmarshallUtilityDefinition(remote)

	origReader := bytes.NewReader([]byte(localYaml))
	origBuffer := new(bytes.Buffer)
	editReader := bytes.NewReader([]byte(remoteYaml))
	editBuffer := new(bytes.Buffer)

	serr := internal.SortYAML(origReader, origBuffer, 2)
	if serr != nil {
		common.Logger.WithError(serr).Error("Error sorting original yaml")
	}
	serr = internal.SortYAML(editReader, editBuffer, 2)
	if serr != nil {
		common.Logger.WithError(serr).Error("Error sorting edit yaml")
	}

	sortedOriginal := origBuffer.String()
	sortedEdit := editBuffer.String()

	if sortedOriginal == sortedEdit {
		return "same"
	}
	return "different"
}
func init() {
	RootCmd.AddCommand(statusCmd)
	statusCmd.Flags().BoolVar(&statusP.Environments, "env", false, "Use this switch to operate on the environment definitions")
}
