package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fieldsCmd represents the fields command
var fieldsCmd = &cobra.Command{
	Use:   "fields",
	Short: fieldsHelp.Short(),
	Long:  fieldsHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		displayFieldHelp()
	},
}

func displayFieldHelp() {

	help := `A list of valid FIELDS that can be specified by command:

  COMMAND: get|add|update|remove environment

      FIELDS: NAME, REPOSITORY, EXCLUDED, HIDDEN, REMOVE_UPSTREAM

  COMMAND: get ingress

      FIELDS: NAME, NAMESPACE, CLASS, HOSTS, ADDRESS, PORTS, LAUNCHED_BY

  COMMAND: get|add|update|remove sleep

      FIELDS: NAME, SECONDS

  COMMAND: get|add|update|remove utility

      FIELDS: NAME, REPOSITORY, EXEC, HIDDEN, EXCLUDED, SOURCE, ENVIRONMENTS,
              REQUIRECONFIGMAPS, REQUIRESECRETS, HINT, REMOVE_UPSTREAM

  COMMAND: get namespace

      FIELDS: NAMESPACE

  COMMAND: get nodes

      FIELDS: NODE

  COMMAND: get running

      FIELDS: NAME, NAMESPACE, STATUS, LAUNCHED_BY, UTILITY, SHELL/SERVICE

  COMMAND: get service

      FIELDS: NAME, NAMESPACE, TYPE, CLUSTER-IP, EXTERNAL-IP, PORT(S), LAUNCHED_BY

  COMMAND: get serviceaccount

      FIELDS: SERVICE_ACCOUNT

  COMMAND: get sizes

      FIELDS: NAME, CPU_LIMIT, MEM_LIMIT, CPU_REQUEST, MEM_REQUEST

  COMMAND: status

      FIELDS: NAME, STATUS, EXCLUDE

  COMMAND version

      FIELDS: SEMVER, BUILD_DATE, GIT_COMMIT, GIT_REF

`

	fmt.Println(help)
}

func init() {
	RootCmd.AddCommand(fieldsCmd)
}
