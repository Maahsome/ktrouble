package get

import (
	"fmt"
	"text/tabwriter"
	"os"

	"ktrouble/defaults"

	"github.com/spf13/cobra"
)

type envVarEntry struct {
	Name        string
	Description string
}

var ktroublEnvVars = []envVarEntry{
	{"GIT_TOKEN", "Git authentication token (name overridable via config GitTokenVar)"},
	{"KUBECONFIG", "Path to kubernetes config file"},
	{"KTROUBLE_CONFIG", "Config file path override"},
	{"KTROUBLE_DEFAULT_INGRESS_TEMPLATE", "Custom ingress template file name"},
	{"KTROUBLE_DEFAULT_SERVICE_TEMPLATE", "Custom service template file name"},
	{"KTROUBLE_DEFAULT_TEMPLATE", "Custom pod template file name"},
	{"KTROUBLE_PAGER", "Custom pager for diffs"},
	{"NAMESPACE", "Default namespace (overridden by --namespace flag)"},
	{"PAGER", "Fallback pager for diffs"},
	{"USER", "Username label for pod filtering and creation"},
	{"XDG_CONFIG_HOME", "XDG Base Directory for config file location"},
}

// envVarsCmd represents the envvars command
var envVarsCmd = &cobra.Command{
	Use:     "envvars",
	Aliases: defaults.GetEnvVarsAliases,
	Short:   getEnvVarsHelp.Short(),
	Long:    getEnvVarsHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		if !c.NoHeaders {
			fmt.Fprintln(w, "VARIABLE\tDESCRIPTION")
		}
		for _, ev := range ktroublEnvVars {
			fmt.Fprintf(w, "%s\t%s\n", ev.Name, ev.Description)
		}
		w.Flush()
	},
}

func init() {
	getCmd.AddCommand(envVarsCmd)
}
