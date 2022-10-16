package gitRepoManager

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"semangit/src/utils"
)

const RevisionNone = ""

type gitRepoManager struct {
	worktree *git.Worktree
}

func NewGitRepoManger(repoDir string) gitRepoManager {
	repo := utils.GetResultOrPanicError(git.PlainOpen(repoDir))
	worktree := utils.GetResultOrPanicError(repo.Worktree())
	return gitRepoManager{worktree}
}

func (m *gitRepoManager) Checkout(refName string) {
	utils.PanicError(m.worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/heads/" + refName),
		Force:  true,
	}))
}
