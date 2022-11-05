package gitrepo

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"semangit/src/utils"
)

const RevisionNone = ""

type gitRepoManager struct {
	repo *git.Repository
}

func NewGitRepoManger(repoDir string) gitRepoManager {
	repo := utils.GetResultOrPanicError(git.PlainOpen(repoDir))
	return gitRepoManager{repo}
}

func (m *gitRepoManager) Checkout(refName string) {
	worktree := utils.GetResultOrPanicError(m.repo.Worktree())
	utils.PanicError(worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/heads/" + refName),
		Force:  true,
	}))
}

// ListChangedFiles Returns the list of filenames that are changed between the two given git revisions.
func (m *gitRepoManager) ListChangedFiles(fromRevision string, toRevision string) []string {
	fromHash := utils.GetResultOrPanicError(m.repo.ResolveRevision(plumbing.Revision(fromRevision)))
	fromCommit := utils.GetResultOrPanicError(m.repo.CommitObject(*fromHash))
	toHash := utils.GetResultOrPanicError(m.repo.ResolveRevision(plumbing.Revision(toRevision)))
	toCommit := utils.GetResultOrPanicError(m.repo.CommitObject(*toHash))
	patch := utils.GetResultOrPanicError(fromCommit.Patch(toCommit))
	filePathsSet := make(map[string]bool)
	for _, filePatch := range patch.FilePatches() {
		fromFile, toFile := filePatch.Files()
		if fromFile != nil {
			filePathsSet[fromFile.Path()] = true
		}
		if toFile != nil {
			filePathsSet[toFile.Path()] = true
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
