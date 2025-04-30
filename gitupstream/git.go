package gitupstream

import (
	"fmt"
	"io"
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
	GetUpstreamDefs() (objects.UtilityPodList, map[string]objects.UtilityPod, objects.EnvironmentList, map[string]objects.Environment)
	GetNewUpstreamDefs(localDefs objects.UtilityPodList) (objects.UtilityPodList, map[string]objects.UtilityPod)
	PushLocals(localDefs objects.UtilityPodList) bool
	GetUpstreamEnvDefs() (objects.EnvironmentList, map[string]objects.Environment)
	GetNewUpstreamEnvDefs(localDefs objects.EnvironmentList) (objects.EnvironmentList, map[string]objects.Environment)
	PushEnvLocals(localDefs objects.EnvironmentConfig) bool
	VersionDirectoryExists(version string) bool
	Migrate(fromVer string, toVer string) bool
}

type gitUpstream struct {
	User    string
	Token   string
	GitURL  string
	Version string
}

func New(user string, token string, url string, version string) GitUpstream {

	return &gitUpstream{
		User:    user,
		Token:   token,
		GitURL:  url,
		Version: version,
	}
}

func (gu *gitUpstream) GetUpstreamDefs() (objects.UtilityPodList, map[string]objects.UtilityPod,
	objects.EnvironmentList, map[string]objects.Environment) {

	var storer *memory.Storage
	var fs billy.Filesystem

	storer = memory.NewStorage()
	fs = memfs.New()

	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: gu.GitURL,
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

	common.Logger.Tracef("Version: %s", gu.Version)
	rootDir := "/"
	if gu.Version != "v0" {
		rootDir = fmt.Sprintf("/%s/", gu.Version)
	}
	dirInfo, derr := fs.ReadDir(rootDir)
	if derr != nil {
		common.Logger.WithError(derr).Fatal("Failed to read dir")
	}

	remoteDefs := objects.UtilityPodList{}
	remoteDefsMap := make(map[string]objects.UtilityPod, 0)
	remoteEnvDefs := objects.EnvironmentList{}
	remoteEnvDefsMap := make(map[string]objects.Environment, 0)

	for _, di := range dirInfo {
		if strings.HasSuffix(di.Name(), ".yaml") {
			file, err := fs.Open(fmt.Sprintf("%s%s", rootDir, di.Name()))
			if err != nil {
				common.Logger.Errorf("failed to open file: %s", di.Name())
			}
			data, err := io.ReadAll(file)
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

		if di.Name() == "ktrouble-environments.yml" {
			common.Logger.Tracef("Found environment definition file: %s", di.Name())
			file, err := fs.Open(fmt.Sprintf("%s%s", rootDir, di.Name()))
			if err != nil {
				common.Logger.Errorf("failed to open file: %s", di.Name())
			}
			data, err := io.ReadAll(file)
			if err != nil {
				common.Logger.Errorf("failed to read file data from %s", di.Name())
			}
			envDefs := objects.EnvironmentConfig{}
			merr := yaml.Unmarshal(data, &envDefs)
			if merr != nil {
				common.Logger.Error("failed to unmarshall utility definition")
			}
			remoteEnvDefs = append(remoteEnvDefs, envDefs.Environments...)
			for _, envDef := range envDefs.Environments {
				remoteEnvDefsMap[envDef.Name] = envDef
			}
		}

	}
	return remoteDefs, remoteDefsMap, remoteEnvDefs, remoteEnvDefsMap
}

func (gu *gitUpstream) GetUpstreamEnvDefs() (objects.EnvironmentList, map[string]objects.Environment) {

	var storer *memory.Storage
	var fs billy.Filesystem

	storer = memory.NewStorage()
	fs = memfs.New()

	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: gu.GitURL,
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

	common.Logger.Tracef("Version: %s", gu.Version)
	rootDir := "/"
	if gu.Version != "v0" {
		rootDir = fmt.Sprintf("/%s/", gu.Version)
	}
	dirInfo, derr := fs.ReadDir(rootDir)
	if derr != nil {
		common.Logger.WithError(derr).Fatal("Failed to read dir")
	}

	remoteDefs := objects.EnvironmentList{}
	remoteDefsMap := make(map[string]objects.Environment, 0)

	for _, di := range dirInfo {
		common.Logger.Tracef("checking file: %s", di.Name())
		if di.Name() == "ktrouble-environments.yml" {
			common.Logger.Tracef("Found environment definition file: %s", di.Name())
			file, err := fs.Open(fmt.Sprintf("%s%s", rootDir, di.Name()))
			if err != nil {
				common.Logger.Errorf("failed to open file: %s", di.Name())
			}
			data, err := io.ReadAll(file)
			if err != nil {
				common.Logger.Errorf("failed to read file data from %s", di.Name())
			}
			envDefs := objects.EnvironmentConfig{}
			merr := yaml.Unmarshal(data, &envDefs)
			if merr != nil {
				common.Logger.Error("failed to unmarshall utility definition")
			}
			remoteDefs = append(remoteDefs, envDefs.Environments...)
			for _, envDef := range envDefs.Environments {
				remoteDefsMap[envDef.Name] = envDef
			}
		}
	}
	return remoteDefs, remoteDefsMap
}

func (gu *gitUpstream) GetNewUpstreamEnvDefs(localDefs objects.EnvironmentList) (objects.EnvironmentList, map[string]objects.Environment) {

	remoteDefs, _ := gu.GetUpstreamEnvDefs()

	missingDefs := objects.EnvironmentList{}
	allDefsMap := make(map[string]objects.Environment, 0)

	for _, def := range remoteDefs {
		if missingEnvLocally(def.Name, localDefs) {
			missingDefs = append(missingDefs, def)
		}
		allDefsMap[def.Name] = def
	}
	return missingDefs, allDefsMap
}

func (gu *gitUpstream) GetNewUpstreamDefs(localDefs objects.UtilityPodList) (objects.UtilityPodList, map[string]objects.UtilityPod) {

	remoteDefs, _, _, _ := gu.GetUpstreamDefs()

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

func (gu *gitUpstream) PushEnvLocals(envConfig objects.EnvironmentConfig) bool {
	var storer *memory.Storage
	var fs billy.Filesystem

	storer = memory.NewStorage()
	fs = memfs.New()

	repo, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: gu.GitURL,
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

	common.Logger.Tracef("Version: %s", gu.Version)
	rootDir := ""
	if gu.Version != "v0" {
		rootDir = fmt.Sprintf("%s/", gu.Version)
	}

	adding := []string{}
	filename := fmt.Sprintf("%s%s", rootDir, "ktrouble-environments.yml")

	// Read existing file data
	existingEnvConfig := objects.EnvironmentConfig{}
	file, err := fs.Open(filename)
	if err == nil {
		defer file.Close()
		content, rerr := io.ReadAll(file)
		if rerr == nil && len(content) > 0 {
			yerr := yaml.Unmarshal(content, &existingEnvConfig)
			if yerr != nil {
				common.Logger.WithError(yerr).Error("Failed to unmarshal existing YAML")
			}
		}
	}

	addEnvs := []objects.Environment{}
	for _, v := range envConfig.Environments {
		common.Logger.Tracef("Adding env %s", v.Name)
		notFound := true
		for i, e := range existingEnvConfig.Environments {
			if v.Name == e.Name {
				common.Logger.Tracef("updating existing env %s", v.Name)
				existingEnvConfig.Environments[i].Source = "ktrouble-utils"
				existingEnvConfig.Environments[i].ExcludeFromShare = v.ExcludeFromShare
				existingEnvConfig.Environments[i].Repository = v.Repository
				if v.RemoveUpstream {
					existingEnvConfig.Environments = objects.RemoveEnvIndex(existingEnvConfig.Environments, i)
				}
				notFound = false
			}
		}
		if notFound {
			addEnvs = append(addEnvs, v)
		}
	}
	common.Logger.Tracef("Adding new environments\n%#v", addEnvs)
	existingEnvConfig.Environments = append(existingEnvConfig.Environments, addEnvs...)

	// Marshal the new envConfig into map[string]interface{}
	newDataBytes, merr := yaml.Marshal(existingEnvConfig)
	if merr != nil {
		common.Logger.Fatal("Error Marshaling YAML data for write")
	}

	outFile, oerr := fs.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if oerr != nil {
		common.Logger.WithError(oerr).Error("failed to open file for writing")
		return false
	}
	defer outFile.Close()

	_, werr := outFile.Write(newDataBytes)
	if werr != nil {
		common.Logger.WithError(werr).Error("failed to write to file")
		return false
	}

	_, aerr := worktree.Add(filename)
	if aerr != nil {
		common.Logger.WithError(aerr).Fatal("Failed to add file")
	}
	adding = append(adding, filename)

	// END Push to Git

	status, err := worktree.Status()
	if err != nil {
		common.Logger.Error("failed to get worktree status")
	}
	fmt.Println(status.String())

	_, cerr := worktree.Commit(fmt.Sprintf("%s is submitting environment definition(s): %s", gu.User, strings.Join(adding, ",")), &git.CommitOptions{})
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

func (gu *gitUpstream) PushLocals(localDefs objects.UtilityPodList) bool {

	var storer *memory.Storage
	var fs billy.Filesystem

	storer = memory.NewStorage()
	fs = memfs.New()

	repo, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: gu.GitURL,
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

	common.Logger.Tracef("Version: %s", gu.Version)
	rootDir := ""
	if gu.Version != "v0" {
		rootDir = fmt.Sprintf("%s/", gu.Version)
	}

	adding := []string{}
	for _, v := range localDefs {

		v.Source = "ktrouble-utils"
		defData, merr := yaml.Marshal(v)
		if merr != nil {
			common.Logger.Fatal("Error Marshaling YAML data for write")
		}
		file, err := fs.OpenFile(fmt.Sprintf("%s%s.yaml", rootDir, v.Name), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
		if err != nil {
			common.Logger.WithError(err).Error("failed to open file")
			return false
		}
		_, werr := file.Write(defData)
		if werr != nil {
			common.Logger.WithError(werr).Error("failed to write to file")
			return false
		}
		if v.RemoveUpstream {
			fs.Remove(fmt.Sprintf("%s%s.yaml", rootDir, v.Name))
		}
		_, aerr := worktree.Add(fmt.Sprintf("%s%s.yaml", rootDir, v.Name))
		if aerr != nil {
			common.Logger.WithError(aerr).Fatal("Failed to add file")
		}
		adding = append(adding, fmt.Sprintf("%s%s", rootDir, v.Name))
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

func missingEnvLocally(name string, localDefs objects.EnvironmentList) bool {
	for _, v := range localDefs {
		if name == v.Name {
			return false
		}
	}
	return true
}
