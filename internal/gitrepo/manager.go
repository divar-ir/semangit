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
func (m *gitRepoManager) ListChangedFiles(srcRevision string, destRevision string) []string {
	srcHash := utils.GetResultOrPanic(m.repo.ResolveRevision(plumbing.Revision(srcRevision)))
	srcCommit := utils.GetResultOrPanic(m.repo.CommitObject(*srcHash))
	destHash := utils.GetResultOrPanic(m.repo.ResolveRevision(plumbing.Revision(destRevision)))
	destCommit := utils.GetResultOrPanic(m.repo.CommitObject(*destHash))
	patch := utils.GetResultOrPanic(srcCommit.Patch(destCommit))
	filePathsSet := make(map[string]bool)
	for _, filePatch := range patch.FilePatches() {
		srcFile, destFile := filePatch.Files()
		if srcFile != nil {
			filePathsSet[srcFile.Path()] = true
		}
		if destFile != nil {
			filePathsSet[destFile.Path()] = true
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
