package main

import (
	"fmt"
	"semangit/src/cli"
	"semangit/src/gitRepoManager"
)

func main() {
	c := cli.NewCliAndRun()
	fmt.Println("Repo: " + c.GetRepoDir())
	repoManager := gitRepoManager.NewGitRepoManger(c.GetRepoDir())
	if c.GetFromRevision() != gitRepoManager.RevisionNone {
		repoManager.Checkout(c.GetFromRevision())
	}
	repoManager.Checkout(c.GetToRevision())
	fmt.Println("Hooray!")
}
