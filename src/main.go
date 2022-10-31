package main

import (
	"fmt"
	"semangit/src/cli"
	"semangit/src/gitrepo"
	"semangit/src/utils"
	"semangit/src/versionanalyzers"
)

type runner struct {
	versionAnalyzer versionanalyzers.VersionAnalyzer
}

func (r runner) run() {
	c := cli.RunNewCli()
	fmt.Println("Repo: " + c.GetRepoDir())
	repoManager := gitrepo.NewGitRepoManger(c.GetRepoDir())
	r.versionAnalyzer = versionanalyzers.GetVersionAnalyzer(c.GetVersionAnalyzerName())
	if c.GetFromRevision() != gitrepo.RevisionNone {
		repoManager.Checkout(c.GetFromRevision())
	}
	// TODO: Check if r.versionAnalyzer.ChangeNeedsVersionUpdate
	fromVersion := utils.GetResultOrPanicError(r.versionAnalyzer.ReadVersion(c.GetRepoDir(), c.GetCurrentVersionAnalyzerArgumentValues()))
	repoManager.Checkout(c.GetToRevision())
	toVersion := utils.GetResultOrPanicError(r.versionAnalyzer.ReadVersion(c.GetRepoDir(), c.GetCurrentVersionAnalyzerArgumentValues()))

	fmt.Printf("TODO: Compare versions! From Version: %s To Version: %s\n", fromVersion, toVersion)
	fmt.Printf("Compare result: %d\n", r.versionAnalyzer.CompareVersions(fromVersion, toVersion))
}

func main() {
	r := runner{}
	r.run()
}
