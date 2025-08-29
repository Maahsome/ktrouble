package gitupstream

import (
	"fmt"
	"io"
	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/migrate"
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

func (gu *gitUpstream) Migrate(fromVer string, toVer string, dryRun bool) bool {
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

	modifiedFiles := []string{}
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
				modifiedFiles = append(modifiedFiles, fmt.Sprintf("%s/%s", toVer, di.Name()))
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
				modifiedFiles = append(modifiedFiles, fmt.Sprintf("%s/%s", toVer, di.Name()))
			}
		}
	} else {
		// Copy all files from the old version directory to the new version directory
		common.Logger.Tracef("Migrating from %s to %s", fromVer, toVer)
		dirMap := make(map[string]string)
		for _, di := range dirInfo {
			if di.IsDir() {
				dirMap[di.Name()] = di.Name()
			}
		}
		// check to see if "old version dir" exists, if not, then we simply
		// create a new directory, leave it empty, and report that the user
		// simply push their local config upstream
		if _, ok := dirMap[fromVer]; !ok {
			common.Logger.Tracef("Old version dir %s does not exist, creating empty dir %s", fromVer, toVer)
			err := fs.MkdirAll(toVer, 0755)
			if err != nil {
				common.Logger.WithError(err).Fatal("Failed to create new version dir")
			}
			// create a file named "ktrouble-environments.yml" with a single line
			// of content:
			// environments:
			ke, err := fs.Create(fmt.Sprintf("%s/ktrouble-environments.yml", toVer))
			if err != nil {
				common.Logger.WithError(err).Fatal("Failed to create environments file")
			}
			_, err = ke.Write([]byte("environments:\n"))
			if err != nil {
				common.Logger.WithError(err).Fatal("Failed to write environments file")
			}
			err = ke.Close()
			if err != nil {
				common.Logger.WithError(err).Fatal("Failed to close environments file")
			}
			_, aerr := worktree.Add(fmt.Sprintf("%s/ktrouble-environments.yml", toVer))
			if aerr != nil {
				common.Logger.WithError(aerr).Fatal("Failed to add file")
			}
			common.Logger.Infof("There was no previous upstream version to migrate, a new %s directory has been created", toVer)
			common.Logger.Infof("Please use 'ktrouble push' to add the initial config to the upstream git repository")
		} else {
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

								common.Logger.Tracef("Adding %s/%s to worktree", toVer, odi.Name())
								_, aerr := worktree.Add(fmt.Sprintf("%s/%s", toVer, odi.Name()))
								if aerr != nil {
									common.Logger.WithError(aerr).Fatal("Failed to add file")
								}
								modifiedFiles = append(modifiedFiles, fmt.Sprintf("%s/%s", toVer, odi.Name()))
							}
							if odi.Name() == "ktrouble-environments.yml" {
								err = MigrateEnvironmentYAML(odi.Name(), fs, fromVer, toVer)
								if err != nil {
									common.Logger.WithError(err).Fatal("Failed to migrate utility yaml")
								}
								common.Logger.Tracef("Adding %s/%s to worktree", toVer, odi.Name())
								_, aerr := worktree.Add(fmt.Sprintf("%s/%s", toVer, odi.Name()))
								if aerr != nil {
									common.Logger.WithError(aerr).Fatal("Failed to add file")
								}
								modifiedFiles = append(modifiedFiles, fmt.Sprintf("%s/%s", toVer, odi.Name()))
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

	if dryRun {
		fileList := ask.PromptForGenericList(modifiedFiles, "Select the file(s) to output to stdout")
		for _, file := range fileList {
			data, err := fs.Open(file)
			if err != nil {
				common.Logger.WithError(err).Error("Failed to read file")
				continue
			}
			defer data.Close()
			content, err := io.ReadAll(data)
			if err != nil {
				common.Logger.WithError(err).Error("Failed to read file content")
				continue
			}
			fmt.Printf("=== %s ===\n%s\n", file, content)
		}
		fmt.Println("Dry run enabled, not pushing changes")
		return false
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
		return false
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

	switch fromVer {
	case "v0":
		utilDef := objects.UtilityPod{}
		merr := yaml.Unmarshal(data, &utilDef)
		if merr != nil {
			common.Logger.Error("failed to unmarshall utility definition")
		}

		defData, merr := yaml.Marshal(utilDef)
		if merr != nil {
			common.Logger.Fatal("Error Marshaling YAML data for write")
		}
		common.Logger.Tracef("Migrating utility definition %s from %s to %s, filename %s", utilDef.Name, fromVer, toVer, filename)
		WriteNewUtilityFile(toVer, filename, fs, defData)
	case "v1":
		utilDef := objects.UtilityPodV1{}
		merr := yaml.Unmarshal(data, &utilDef)
		if merr != nil {
			common.Logger.Error("failed to unmarshall utility definition")
		}

		utilDefV2 := migrate.UpdateUtilityV2(utilDef)
		defData, merr := yaml.Marshal(utilDefV2)
		if merr != nil {
			common.Logger.Fatal("Error Marshaling YAML data for write")
		}
		common.Logger.Tracef("Migrating utility definition %s from %s to %s, filename %s", utilDef.Name, fromVer, toVer, filename)
		WriteNewUtilityFile(toVer, filename, fs, defData)

	}

	return nil
}

func WriteNewUtilityFile(toVer string, filename string, fs billy.Filesystem, data []byte) error {
	newFile, err := fs.OpenFile(fmt.Sprintf("%s/%s", toVer, filename), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		common.Logger.WithError(err).Error("failed to open file")
		return err
	}
	defer newFile.Close()

	common.Logger.Tracef("Writing new definition file %s", fmt.Sprintf("%s/%s", toVer, filename))
	_, werr := newFile.Write(data)
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
