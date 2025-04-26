package cmd

import (
	"ktrouble/common"
	"ktrouble/objects"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

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
		utilityDefinitionDiffs()
	},
}

func unmarshallUtilityDefinition(util objects.UtilityPod) string {
	utilYAML, rerr := yaml.Marshal(util)
	if rerr != nil {
		return ""
	}

	return strings.Replace(string(utilYAML), "\n", "\n    ", -1)

}

func utilityDefinitionDiffs() {

	utilsDifferent := map[string]objects.Status{}

	status := UtilityDefinitionStatus()
	for _, s := range status {
		if s.Status != "different" {
			utilsDifferent[s.Name] = s
		}
	}

	_, remoteDefsMap := c.GitUpstream.GetUpstreamDefs()

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
}
