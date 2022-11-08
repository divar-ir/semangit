package gitrepo

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"semangit/internal/utils"
)

const RevisionNone = ""

type gitRepoManager struct {
	repo *git.Repository
}

func NewGitRepoManger(repoDir string) gitRepoManager {
	repo := utils.GetResultOrPanic(git.PlainOpen(repoDir))
	return gitRepoManager{repo}
}

func (m *gitRepoManager) Checkout(refName string) {
	worktree := utils.GetResultOrPanic(m.repo.Worktree())
	utils.PanicError(worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/heads/" + refName),
		Force:  true,
	}))
}

// ListChangedFiles Returns the list of filenames that are changed between the two given git revisions.
func (m *gitRepoManager) ListChangedFiles(oldRevision string, newRevision string) []string {
	oldHash := utils.GetResultOrPanic(m.repo.ResolveRevision(plumbing.Revision(oldRevision)))
	oldCommit := utils.GetResultOrPanic(m.repo.CommitObject(*oldHash))
	newHash := utils.GetResultOrPanic(m.repo.ResolveRevision(plumbing.Revision(newRevision)))
	newCommit := utils.GetResultOrPanic(m.repo.CommitObject(*newHash))
	patch := utils.GetResultOrPanic(oldCommit.Patch(newCommit))
	filePathsSet := make(map[string]bool)
	for _, filePatch := range patch.FilePatches() {
		oldFile, newFile := filePatch.Files()
		if oldFile != nil {
			filePathsSet[oldFile.Path()] = true
		}
		if newFile != nil {
			filePathsSet[newFile.Path()] = true
		}
	}

	changedFilePaths := make([]string, len(filePathsSet))

	i := 0
	for path := range filePathsSet {
		changedFilePaths[i] = path
		i++
	}

	return changedFilePaths
}
