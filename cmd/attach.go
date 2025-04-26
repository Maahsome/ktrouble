package cmd

import (
	"ktrouble/defaults"

	"github.com/spf13/cobra"
)

type attachParam struct {
	PromptForSecrets    bool
	PromptForConfigMaps bool
	PromptForVolumes    bool
	PromptForMysql      bool
	PromptForPostgres   bool
	CreateIngress       bool
	Port                int
	Host                string
	Path                string
}

var a attachParam

// attachCmd represents the default command
var attachCmd = &cobra.Command{
	Use:     "attach",
	Aliases: defaults.AttachAliases,
	Short:   "attach a kubernetes troubleshooting container to a running pod",
	Long: `EXAMPLE:
  Just running ktrouble attach will prompt for all the things required to run.
  Attaching a container to an existing pod is done through the Ephemeral
  Container feature of Kubernetes.  This feature is only available in
  Kubernetes 1.16 and later, and must be enabled in the cluster.  The way that
  Ephemeral Containers work is that a new container is created in the same
  namespace as the pod, and the new container is attached to the pod's network
  namespace.  This allows the new container to see the same network as the pod.
  These Ephemeral Containers are not persisted, and are removed when the primary
  command that starts the container exits.  From the command line, you launch a
  new container and after you exit the container, the Ephemeral Container is
  terminated.  In order to allow us to attach a container and also be able to
  exec and exit the container without it terminating, we simply run the "sleep"
  command, and when that sleep duration is over, the container will exit.  There
  is NO other way to remove an Ephemeral Container definition from a pod.

  > ktrouble attach

`,
	Run: func(cmd *cobra.Command, args []string) {
		utility := ""
		sa := ""
		if len(args) > 0 && len(args[0]) > 0 {
			utility = args[0]
		}
		if len(args) > 1 && len(args[1]) > 0 {
			sa = args[1]
		}

		standardAttach(utility, sa)
	},
}

func init() {
	RootCmd.AddCommand(attachCmd)
}
