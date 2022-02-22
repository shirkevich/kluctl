package git

import (
	"github.com/codablock/kluctl/pkg/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// PoorMansClone poor mans clone from a local repo, which does not rely on go-git using git-upload-pack
func PoorMansClone(sourceDir string, targetDir string, ref string) error {
	err := os.MkdirAll(targetDir, 0o700)
	if err != nil {
		return err
	}
	err = os.Mkdir(filepath.Join(targetDir, ".git"), 0o700)
	if err != nil {
		return err
	}
	des, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}
	for _, de := range des {
		s := filepath.Join(sourceDir, de.Name())
		d := filepath.Join(targetDir, ".git", de.Name())
		if de.Name() == ".cache.lock" {
			continue
		}
		if de.Name() == "objects" {
			err = os.Symlink(s, d)
			if err != nil && runtime.GOOS == "windows" {
				// windows 10 does not support symlinks as unprivileged users...
				err = utils.CopyDir(s, d)
			}
		} else {
			if de.IsDir() {
				err = utils.CopyDir(s, d)
			} else {
				err = utils.CopyFile(s, d)
			}
		}
		if err != nil {
			return err
		}
	}

	gitConfigReader, err := os.Open(filepath.Join(targetDir, ".git", "config"))
	if err != nil {
		return err
	}
	defer gitConfigReader.Close()
	gitConfig, err := config.ReadConfig(gitConfigReader)
	if err != nil {
		return err
	}
	gitConfig.Core.IsBare = false

	b, err := gitConfig.Marshal()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(targetDir, ".git", "config"), b, 0o600)
	if err != nil {
		return err
	}

	r, err := git.PlainOpen(targetDir)
	if err != nil {
		return err
	}

	wt, err := r.Worktree()
	if err != nil {
		return err
	}
	var ref2 plumbing.ReferenceName
	if ref != "" {
		ref2 = plumbing.NewBranchReferenceName(ref)
	} else {
		x, err := r.Reference("HEAD", true)
		if err != nil {
			return err
		}
		ref2 = x.Name()
	}
	err = wt.Checkout(&git.CheckoutOptions{
		Branch: ref2,
	})
	if err != nil {
		return err
	}
	return nil
}
