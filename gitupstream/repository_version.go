package gitupstream

import (
	"ktrouble/common"

	billy "github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
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
