package gitupstream

import (
	"io/ioutil"
	"ktrouble/common"
	"ktrouble/objects"
	"strings"

	billy "github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
	"gopkg.in/yaml.v2"
)

type GitUpstream interface {
	GetUpstreamDefs() (objects.UtilityPodList, map[string]objects.UtilityPod)
	GetNewUpstreamDefs(localDefs objects.UtilityPodList) (objects.UtilityPodList, map[string]objects.UtilityPod)
}

type gitUpstream struct {
	User  string
	Token string
}

func New(user string, token string) GitUpstream {

	return &gitUpstream{
		User:  user,
		Token: token,
	}
}

func (gu *gitUpstream) GetUpstreamDefs() (objects.UtilityPodList, map[string]objects.UtilityPod) {

	var storer *memory.Storage
	var fs billy.Filesystem

	storer = memory.NewStorage()
	fs = memfs.New()

	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://git.alteryx.com/futurama/farnsworth/tools/ktrouble-utils.git",
		Auth: &gitHttp.BasicAuth{
			Username: gu.User,
			Password: gu.Token,
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
			remoteDefs = append(remoteDefs, utilDef)
			remoteDefsMap[utilDef.Name] = utilDef
		}
	}
	return remoteDefs, remoteDefsMap
}
func (gu *gitUpstream) GetNewUpstreamDefs(localDefs objects.UtilityPodList) (objects.UtilityPodList, map[string]objects.UtilityPod) {

	remoteDefs, _ := gu.GetUpstreamDefs()

	missingDefs := objects.UtilityPodList{}
	missingDefsMap := make(map[string]objects.UtilityPod, 0)

	for _, def := range remoteDefs {
		if missingLocally(def.Name, localDefs) {
			missingDefs = append(missingDefs, def)
			missingDefsMap[def.Name] = def
		}
	}
	return missingDefs, missingDefsMap
}

func missingLocally(name string, localDefs objects.UtilityPodList) bool {
	for _, v := range localDefs {
		if name == v.Name {
			return false
		}
	}
	return true
}
