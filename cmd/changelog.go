package cmd

import (
	"embed"
	"fmt"
	"ktrouble/defaults"
	"strings"

	"github.com/coreos/go-semver/semver"

	// markdown "github.com/Maahsome/go-term-markdown"
	// "github.com/nathan-fiscaletti/consolesize-go"
	"github.com/spf13/cobra"
)

type changeLogParams struct {
	Version string
	Select  bool
	ShowAll bool
}

var (
	//go:embed changelog
	fs  embed.FS
	clp = changeLogParams{}
)

// changelogCmd represents the changelog command
var changelogCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: defaults.ChangelogAliases,
	Short:   "Express the 'version' of ktrouble.",
	Run: func(cmd *cobra.Command, args []string) {

		if clp.ShowAll {
			allLogsOutput()
		} else {
			version := semVer
			if clp.Version != "" {
				version = clp.Version
			}
			outputLog(version)
		}
	},
}

func outputLog(version string) {
	data, _ := fs.ReadFile(fmt.Sprintf("changelog/%s.md", version))
	doc := fmt.Sprintf("# Changelog\n\n%s", data)
	fmt.Println(doc)
}

func allLogsOutput() {
	doc := "# Changelog\n\n"
	files, _ := fs.ReadDir("changelog")

	versions := []*semver.Version{}
	for _, f := range files {
		sver := strings.Replace(f.Name()[1:], ".md", "", 1)
		versions = append(versions, semver.New(sver))
	}

	semver.Sort(versions)

	for _, v := range versions {
		data, _ := fs.ReadFile(fmt.Sprintf("changelog/v%s.md", v.String()))
		doc = fmt.Sprintf("%s%s", doc, data)

	}
	fmt.Println(doc)
}

func init() {
	RootCmd.AddCommand(changelogCmd)
	changelogCmd.Flags().StringVar(&clp.Version, "version", "", "Sepecify the version to display the changelog for")
	// changelogCmd.Flags().BoolVar(&clp.Select, "select", false, "Specify this switch to prompt for a list of versions")
	changelogCmd.Flags().BoolVar(&clp.ShowAll, "all", false, "Specify this switch to show ALL of the changelog entries")
}
