package cmd

import (
	"fmt"
	"ktrouble/common"
	"ktrouble/objects"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type diffParams struct {
	Environments bool
}

var diffP statusParams

// diffCmd represents the status command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Get a context diff on each utility definition",
	Long: `EXAMPLE:
  The 'diff' command will list the differences between your local 'config.yaml'
  file 'utilities' definitions and the remote repository.

  > ktrouble diff
`,
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
		if diffP.Environments {
			environmentDefinitionDiffs()
		} else {
			utilityDefinitionDiffs()
		}
	},
}

func unmarshallEnvDefinition(env objects.Environment) string {
	envYAML, rerr := yaml.Marshal(env)
	if rerr != nil {
		return ""
	}

	return strings.Replace(string(envYAML), "\n", "\n    ", -1)

}

func unmarshallUtilityDefinition(util objects.UtilityPod) string {

	// These fields we want to not affect the diff
	// Upstream definitions will ALWAYS be RemoveUpstream = false
	// Upstream definitions will ALWAYS have Hidden = false as well
	util.RemoveUpstream = false
	util.Hidden = false
	utilYAML, rerr := yaml.Marshal(util)
	if rerr != nil {
		return ""
	}

	return strings.Replace(string(utilYAML), "\n", "\n    ", -1)

}

func environmentDefinitionDiffs() {
	envsDifferent := map[string]objects.Status{}

	status := EnvironmentDefinitionStatus()
	for _, s := range status {
		if s.Status != "different" {
			envsDifferent[s.Name] = s
		}
	}

	_, remoteDefsMap := c.GitUpstream.GetUpstreamEnvDefs()

	envNames := []string{}
	for _, l := range c.EnvDefs {
		envNames = append(envNames, l.Name)
	}
	sort.Strings(envNames)

	localDefs := "environments:\n"

	for _, l := range envNames {
		if _, ok := envsDifferent[l]; !ok {
			eDef := c.EnvMap[l]
			yDef := unmarshallEnvDefinition(eDef)
			localDefs += "  - " + yDef + "\n"
		}
	}

	remoteNames := []string{}
	for _, r := range remoteDefsMap {
		remoteNames = append(remoteNames, r.Name)
	}
	sort.Strings(remoteNames)

	remoteDefs := "environments:\n"
	for _, l := range remoteNames {
		if _, ok := envsDifferent[l]; !ok {
			eDef := remoteDefsMap[l]
			yDef := unmarshallEnvDefinition(eDef)
			remoteDefs += "  - " + yDef + "\n"
		}
	}
	common.Logger.Tracef("local: %s", localDefs)
	common.Logger.Tracef("\n\n-------------------------------\n\nremote: %s", remoteDefs)
	common.OutputContextDiff([]byte(localDefs), []byte(remoteDefs), 5)
}

func utilityDefinitionDiffs() {

	utilsDifferent := map[string]objects.Status{}

	status := UtilityDefinitionStatus()
	for _, s := range status {
		if s.Status != "different" {
			utilsDifferent[s.Name] = s
		}
	}

	_, remoteDefsMap, _, _ := c.GitUpstream.GetUpstreamDefs()

	utilNames := []string{}
	for _, l := range c.UtilDefs {
		utilNames = append(utilNames, l.Name)
	}
	sort.Strings(utilNames)

	localDefs := "UtilityDefinitions:\n"

	for _, l := range utilNames {
		if _, ok := utilsDifferent[l]; !ok {
			uDef := c.UtilMap[l]
			yDef := unmarshallUtilityDefinition(uDef)
			localDefs += "  - " + yDef + "\n"
		}
	}

	remoteNames := []string{}
	for _, r := range remoteDefsMap {
		remoteNames = append(remoteNames, r.Name)
	}
	sort.Strings(remoteNames)

	remoteDefs := "UtilityDefinitions:\n"
	for _, l := range remoteNames {
		if _, ok := utilsDifferent[l]; !ok {
			uDef := remoteDefsMap[l]
			yDef := unmarshallUtilityDefinition(uDef)
			remoteDefs += "  - " + yDef + "\n"
		}
	}
	common.Logger.Tracef("local: %s", localDefs)
	common.Logger.Tracef("\n\n-------------------------------\n\nremote: %s", remoteDefs)
	common.OutputContextDiff([]byte(localDefs), []byte(remoteDefs), 5)
}

func init() {
	RootCmd.AddCommand(diffCmd)
	diffCmd.Flags().BoolVar(&diffP.Environments, "env", false, "Use this switch to operate on the environment definitions")
}
