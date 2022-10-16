package main

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing"
	"semangit/src/cli"
)
import "github.com/go-git/go-git/v5"

func main() {
	c := cli.NewCliAndRun()
	fmt.Println("Checking repo: " + c.GetRepoDir() + " ...")
	repo, err := git.PlainOpen(c.GetRepoDir())
	if err != nil {
		panic(err)
	}
	worktree, err := repo.Worktree()
	if err != nil {
		panic(err)
	}
	if c.GetFromRevision() != cli.RevisionNone {
		err = worktree.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/heads/" + c.GetFromRevision()),
			Force:  true,
		})
		if err != nil {
			panic(err)
		}
	}
	if c.GetToRevision() == cli.RevisionNone {
		panic("Provide TO revision (-t) to compare the version on it.")
	} else {
		err = worktree.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/heads/" + c.GetToRevision()),
			Force:  true,
		})
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Hooray!")
}
