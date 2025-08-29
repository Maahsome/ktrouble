package cmd

import (
	"fmt"
	"ktrouble/common"

	"github.com/spf13/cobra"
)

type MigrateParam struct {
	DryRun bool
}

var migrateParam = MigrateParam{}

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
		if !c.GitUpstream.Migrate(fromVersion, toVersion, migrateParam.DryRun) {
			common.Logger.Error("Failed to migrate the repository")
			return
		}
		// TODO: what was this for?  This routine simply update upstream.  The
		// Local migration should be handled upon run, not sure what the scoop
		// is here.
		// if !objects.MigrateLocalUtility(c.UtilDefs, toVersion) {
		// 	common.Logger.Error("Failed to migrate the local utility definitions")
		// 	return
		// }
		// if !objects.MigrateLocalEnvironments(c.EnvDefs, toVersion) {
		// 	common.Logger.Error("Failed to migrate the local environment definitions")
		// 	return
		// }
		fmt.Println("Migration complete")
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().BoolVar(&migrateParam.DryRun, "dry-run", false, "Specify --dry-run to simulate the migration without making changes")
}
