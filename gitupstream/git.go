package gitupstream

import (
	"fmt"
	"io/ioutil"
	"ktrouble/common"
	"ktrouble/objects"
	"os"
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
	PushLocals(localDefs objects.UtilityPodList) bool
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
	allDefsMap := make(map[string]objects.UtilityPod, 0)

	for _, def := range remoteDefs {
		if missingLocally(def.Name, localDefs) {
			missingDefs = append(missingDefs, def)
		}
			allDefsMap[def.Name] = def
	}
	return missingDefs, allDefsMap
}

func (gu *gitUpstream) PushLocals(localDefs objects.UtilityPodList) bool {

	var storer *memory.Storage
	var fs billy.Filesystem

	storer = memory.NewStorage()
	fs = memfs.New()

	repo, err := git.Clone(storer, fs, &git.CloneOptions{
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

	worktree, wterr := repo.Worktree()
	if wterr != nil {
		common.Logger.Fatal("Failed to create worktree")
	}

	adding := []string{}
	for _, v := range localDefs {

		v.Source = "ktrouble-utils"
		defData, merr := yaml.Marshal(v)
		if merr != nil {
			common.Logger.Fatal("Error Marshaling YAML data for write")
		}
		file, err := fs.OpenFile(fmt.Sprintf("%s.yaml", v.Name), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
		if err != nil {
			common.Logger.WithError(err).Error("failed to open file")
			return false
		}
		_, werr := file.Write(defData)
		if werr != nil {
			common.Logger.WithError(werr).Error("failed to write to file")
			return false
		}
		_, aerr := worktree.Add(fmt.Sprintf("%s.yaml", v.Name))
		if aerr != nil {
			common.Logger.WithError(aerr).Fatal("Failed to add file")
		}
		adding = append(adding, v.Name)
	}

	status, err := worktree.Status()
	if err != nil {
		common.Logger.Error("failed to get worktree status")
	}
	fmt.Println(status.String())

	_, cerr := worktree.Commit(fmt.Sprintf("%s is submitting utility definition(s): %s", gu.User, strings.Join(adding, ",")), &git.CommitOptions{})
	if cerr != nil {
		common.Logger.WithError(cerr).Fatal("Failed to commit changes")
	}

	perr := repo.Push(&git.PushOptions{
		Auth: &gitHttp.BasicAuth{
			Username: gu.User,
			Password: gu.Token,
		},
		RemoteName: "origin",
	})
	if perr != nil {
		common.Logger.WithError(perr).Fatal("Error pushing branch to origin")
	}

	return true
}

func missingLocally(name string, localDefs objects.UtilityPodList) bool {
	for _, v := range localDefs {
		if name == v.Name {
			return false
		}
	}
	return true
}
