package cmd

import (
	"fmt"
	"ktrouble/common"
	"ktrouble/objects"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: migrateHelp.Short(),
	Long:  migrateHelp.Long(),
	Run: func(cmd *cobra.Command, args []string) {
		if c.GitUpstream.VersionDirectoryExists(fmt.Sprintf("v%d", c.Semver.Major)) {
			common.Logger.Error("The version directory in the repository already exists, no need to re-run the migration command")
			return
		}
		fromVersion := fmt.Sprintf("v%d", c.Semver.Major-1)
		toVersion := fmt.Sprintf("v%d", c.Semver.Major)
		if !c.GitUpstream.Migrate(fromVersion, toVersion) {
			common.Logger.Error("Failed to migrate the repository")
			return
		}
		if !objects.MigrateLocalUtility(c.UtilDefs, toVersion) {
			common.Logger.Error("Failed to migrate the local utility definitions")
			return
		}
		if !objects.MigrateLocalEnvironments(c.EnvDefs, toVersion) {
			common.Logger.Error("Failed to migrate the local environment definitions")
			return
		}
		fmt.Println("Migration complete")
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}
