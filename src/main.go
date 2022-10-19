package main

import (
	"fmt"
	"semangit/src/cli"
	"semangit/src/gitRepoManager"
	"semangit/src/utils"
	"semangit/src/versionComparers"
	"semangit/src/versionReaders"
)

type runner struct {
	versionComparer versionComparers.VersionComparer
	versionReader   versionReaders.VersionReader
}

func (r runner) run() {
	c := cli.RunNewCli()
	fmt.Println("Repo: " + c.GetRepoDir())
	repoManager := gitRepoManager.NewGitRepoManger(c.GetRepoDir())
	r.versionReader = versionReaders.GetVersionReader(c.GetVersionReaderName())
	if c.GetFromRevision() != gitRepoManager.RevisionNone {
		repoManager.Checkout(c.GetFromRevision())
	}
	fromVersion := utils.GetResultOrPanicError(r.versionReader.ReadVersion(c.GetCurrentVersionReaderArguments()))
	repoManager.Checkout(c.GetToRevision())
	toVersion := utils.GetResultOrPanicError(r.versionReader.ReadVersion(c.GetCurrentVersionReaderArguments()))

	fmt.Printf("TODO: Compare versions! From Version: %s To Version: %s\n", fromVersion, toVersion)
	fmt.Printf("Compare result: %d\n", r.versionComparer.Compare(fromVersion, toVersion))
}

func main() {
	r := runner{
		versionComparer: versionComparers.SemanticVersionComparer{},
	}
	r.run()
}
