package cmd

import (
	"fmt"
	"io/ioutil"
	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/objects"
	"os"
	"strings"

	billy "github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull utility definitions from git",
	Long: `EXAMPLE:
  > ktrouble pull
`,
	Run: func(cmd *cobra.Command, args []string) {
		pullUtilityDefinitions()
	},
}

func pullUtilityDefinitions() {
	var storer *memory.Storage
	var fs billy.Filesystem

	gitUser := viper.GetString("gitUser")
	if len(gitUser) == 0 {
		common.Logger.Fatal("gitUser is not set")
	}
	gitTokenVar := ""
	gitToken := viper.GetString("gitToken")
	if len(gitToken) == 0 {
		gitTokenVar = viper.GetString("gitTokenVar")
		if len(gitTokenVar) == 0 {
			common.Logger.Fatal("gitToken or gitTokenVar config option is not set")
		}
		gitToken = os.Getenv(gitTokenVar)
	}

	if len(gitToken) == 0 {
		common.Logger.Fatalf("no git token set, gitToken or %s ENV VAR is not set", gitTokenVar)
	}

	storer = memory.NewStorage()
	fs = memfs.New()

	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://git.alteryx.com/futurama/farnsworth/tools/ktrouble-utils.git",
		Auth: &gitHttp.BasicAuth{
			Username: gitUser,
			Password: gitToken,
		},
		Depth:    1,
		Progress: nil,
	})
	if err != nil {
		common.Logger.WithError(err).Fatal("Error cloning to memory")
	}

	dirInfo, derr := fs.ReadDir("/")
	if derr != nil {
		common.Logger.WithError(derr).Fatal("Failed to read dir")
	}

	remoteDefs := objects.UtilityPodList{}
	remoteDefsMap := make(map[string]objects.UtilityPod, 0)

	for _, di := range dirInfo {
		if strings.HasSuffix(di.Name(), ".yaml") {
			file, err := fs.Open(di.Name())
			if err != nil {
				common.Logger.Errorf("failed to open file: %s", di.Name())
			}
			data, err := ioutil.ReadAll(file)
			if err != nil {
				common.Logger.Errorf("failed to read file data from %s", di.Name())
			}
			utilDef := objects.UtilityPod{}
			merr := yaml.Unmarshal(data, &utilDef)
			if merr != nil {
				common.Logger.Error("failed to unmarshall utility definition")
			}
			if missingLocally(utilDef.Name) {
				remoteDefs = append(remoteDefs, utilDef)
				remoteDefsMap[utilDef.Name] = utilDef
			}
		}
	}

	for _, v := range remoteDefs {
		common.Logger.Warnf("Only Remote: %s", v.Name)
	}

	if len(remoteDefs) > 0 {
		addUtils := ask.PromptForPulledUtility(remoteDefs)

		if len(addUtils) > 0 {
			for _, v := range addUtils {
				c.UtilDefs = append(c.UtilDefs, remoteDefsMap[v])
			}
			viper.Set("utilityDefinitions", c.UtilDefs)
			verr := viper.WriteConfig()
			if verr != nil {
				common.Logger.WithError(verr).Info("Failed to write config")
			}
		} else {
			fmt.Println("No definitions selected")
		}
	} else {
		fmt.Println("Up to date")
	}

}

func missingLocally(name string) bool {
	for _, v := range c.UtilDefs {
		if name == v.Name {
			return false
		}
	}
	return true
}

func init() {
	RootCmd.AddCommand(pullCmd)
}
