package main

import (
	"flag"
	"fmt"
	"github.com/go-git/go-git/v5/plumbing"
)
import "github.com/go-git/go-git/v5"

const RevisionNone = ""

func main() {
	repoDir := flag.String("d", ".", "Repo root path. Defaults to current directory.")
	fromRevision := flag.String("f", RevisionNone, "From revision. A git reference to get version from.")
	toRevision := flag.String("t", RevisionNone, "From revision. A git reference to get version from.")
	flag.Parse()
	fmt.Println("Checking repo: " + *repoDir + " ...")
	repo, err := git.PlainOpen(*repoDir)
	if err != nil {
		panic(err)
	}
	worktree, err := repo.Worktree()
	if err != nil {
		panic(err)
	}
	if *fromRevision != RevisionNone {
		err = worktree.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/heads/" + *fromRevision),
			Force:  true,
		})
		if err != nil {
			panic(err)
		}
	}
	if *toRevision == RevisionNone {
		panic("Provide TO revision (-t) to compare the version on it.")
	} else {
		err = worktree.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/heads/" + *toRevision),
			Force:  true,
		})
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Hooray!")
}
