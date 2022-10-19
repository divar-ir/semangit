package main

import (
	"fmt"
	"semangit/src/cli"
	"semangit/src/gitRepoManager"
	"semangit/src/utils"
	"semangit/src/versionReaders"
)

func main() {
	c := cli.RunNewCli()
	fmt.Println("Repo: " + c.GetRepoDir())
	repoManager := gitRepoManager.NewGitRepoManger(c.GetRepoDir())
	versionReader := versionReaders.GetVersionReader(c.GetVersionReaderName())
	if c.GetFromRevision() != gitRepoManager.RevisionNone {
		repoManager.Checkout(c.GetFromRevision())
	}
	fromVersion := utils.GetResultOrPanicError(versionReader.ReadVersion(c.GetCurrentVersionReaderArguments()))
	repoManager.Checkout(c.GetToRevision())
	toVersion := utils.GetResultOrPanicError(versionReader.ReadVersion(c.GetCurrentVersionReaderArguments()))
	fmt.Printf("TODO: Compare versions! From Version: %s To Version: %s", fromVersion, toVersion)
}
