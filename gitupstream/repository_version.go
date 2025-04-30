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

func (gu *gitUpstream) VersionDirectoryExists(version string) bool {

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

	dirInfo, derr := fs.ReadDir("/")
	if derr != nil {
		common.Logger.WithError(derr).Fatal("Failed to read dir")
	}

	// v0 files are simply in the root of the repo, so if we can read the dir, it exists
	if version == "v0" {
		return true
	}

	for _, di := range dirInfo {
		if di.Name() == version {
			if di.IsDir() {
				return true
			}
		}
	}
	return false
}

func (gu *gitUpstream) Migrate(fromVer string, toVer string) bool {
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

	dirInfo, derr := fs.ReadDir("/")
	if derr != nil {
		common.Logger.WithError(derr).Fatal("Failed to read dir")
	}

	err = fs.MkdirAll(toVer, 0755)
	if err != nil {
		common.Logger.WithError(err).Fatal("Failed to create version directory")
	}

	if fromVer == "v0" {
		// Copy all files from the root to the new version directory
		for _, di := range dirInfo {
			if strings.HasSuffix(di.Name(), ".yaml") {
				err = MigrateUtilityYAML(di.Name(), fs, fromVer, toVer)
				if err != nil {
					common.Logger.WithError(err).Fatal("Failed to migrate utility yaml")
				}
				_, aerr := worktree.Add(fmt.Sprintf("%s/%s", toVer, di.Name()))
				if aerr != nil {
					common.Logger.WithError(aerr).Fatal("Failed to add file")
				}
			}
			if di.Name() == "ktrouble-environments.yml" {
				err = MigrateEnvironmentYAML(di.Name(), fs, fromVer, toVer)
				if err != nil {
					common.Logger.WithError(err).Fatal("Failed to migrate utility yaml")
				}
				_, aerr := worktree.Add(fmt.Sprintf("%s/%s", toVer, di.Name()))
				if aerr != nil {
					common.Logger.WithError(aerr).Fatal("Failed to add file")
				}
			}
		}
	} else {
		// Copy all files from the old version directory to the new version directory
		for _, di := range dirInfo {
			if di.Name() == fromVer {
				if di.IsDir() {
					oldDirInfo, err := fs.ReadDir(di.Name())
					if err != nil {
						common.Logger.WithError(err).Fatal("Failed to read old version dir")
					}
					for _, odi := range oldDirInfo {
						if strings.HasSuffix(odi.Name(), ".yaml") {
							common.Logger.Tracef("Migrating %s/%s to %s/%s", fromVer, odi.Name(), toVer, odi.Name())
							err = MigrateUtilityYAML(odi.Name(), fs, fromVer, toVer)
							if err != nil {
								common.Logger.WithError(err).Fatal("Failed to migrate utility yaml")
							}
							_, aerr := worktree.Add(fmt.Sprintf("%s/%s", toVer, di.Name()))
							if aerr != nil {
								common.Logger.WithError(aerr).Fatal("Failed to add file")
							}
						}
						if di.Name() == "ktrouble-environments.yml" {
							err = MigrateEnvironmentYAML(di.Name(), fs, fromVer, toVer)
							if err != nil {
								common.Logger.WithError(err).Fatal("Failed to migrate utility yaml")
							}
							_, aerr := worktree.Add(fmt.Sprintf("%s/%s", toVer, di.Name()))
							if aerr != nil {
								common.Logger.WithError(aerr).Fatal("Failed to add file")
							}
						}
					}
				}
			}
		}
	}
	status, err := worktree.Status()
	if err != nil {
		common.Logger.Error("failed to get worktree status")
	}
	fmt.Println(status.String())

	_, cerr := worktree.Commit(fmt.Sprintf("%s is migrating utility definitions: %s -> %s", gu.User, fromVer, toVer), &git.CommitOptions{})
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

func MigrateUtilityYAML(filename string, fs billy.Filesystem, fromVer string, toVer string) error {

	fromFile := filename
	if fromVer != "v0" {
		fromFile = fmt.Sprintf("%s/%s", fromVer, filename)
	}
	file, err := fs.Open(fromFile)
	if err != nil {
		common.Logger.Errorf("failed to open file: %s", filename)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		common.Logger.Errorf("failed to read file data from %s", filename)
	}
	utilDef := objects.UtilityPod{}
	merr := yaml.Unmarshal(data, &utilDef)
	if merr != nil {
		common.Logger.Error("failed to unmarshall utility definition")
	}

	defData, merr := yaml.Marshal(utilDef)
	if merr != nil {
		common.Logger.Fatal("Error Marshaling YAML data for write")
	}
	newFile, err := fs.OpenFile(fmt.Sprintf("%s/%s", toVer, filename), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		common.Logger.WithError(err).Error("failed to open file")
		return err
	}
	defer newFile.Close()

	_, werr := newFile.Write(defData)
	if werr != nil {
		common.Logger.WithError(werr).Error("failed to write to file")
		return werr
	}

	return nil
}

func MigrateEnvironmentYAML(filename string, fs billy.Filesystem, fromVer string, toVer string) error {

	fromFile := filename
	if fromVer != "v0" {
		fromFile = fmt.Sprintf("%s/%s", fromVer, filename)
	}
	file, err := fs.Open(fromFile)
	if err != nil {
		common.Logger.Errorf("failed to open file: %s", filename)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		common.Logger.Errorf("failed to read file data from %s", filename)
	}
	envConfigDef := objects.EnvironmentConfig{}
	merr := yaml.Unmarshal(data, &envConfigDef)
	if merr != nil {
		common.Logger.Error("failed to unmarshall environment definitions")
	}

	defData, merr := yaml.Marshal(envConfigDef)
	if merr != nil {
		common.Logger.Fatal("Error Marshaling YAML data for write")
	}
	newFile, err := fs.OpenFile(fmt.Sprintf("%s/%s", toVer, filename), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		common.Logger.WithError(err).Error("failed to open file")
		return err
	}
	defer newFile.Close()

	_, werr := newFile.Write(defData)
	if werr != nil {
		common.Logger.WithError(werr).Error("failed to write to file")
		return werr
	}

	return nil
}
