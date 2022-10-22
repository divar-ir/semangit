package main

import (
	"fmt"
	"semangit/src/cli"
	"semangit/src/gitRepoManager"
	"semangit/src/utils"
	"semangit/src/versionAnalyzers"
	"semangit/src/versionComparers"
)

type runner struct {
	versionComparer versionComparers.VersionComparer
	versionAnalyzer versionAnalyzers.VersionAnalyzer
}

func (r runner) run() {
	c := cli.RunNewCli()
	fmt.Println("Repo: " + c.GetRepoDir())
	repoManager := gitRepoManager.NewGitRepoManger(c.GetRepoDir())
	r.versionAnalyzer = versionAnalyzers.GetVersionAnalyzer(c.GetVersionAnalyzerName())
	if c.GetFromRevision() != gitRepoManager.RevisionNone {
		repoManager.Checkout(c.GetFromRevision())
	}
	// TODO: Check if r.versionAnalyzer.ChangeNeedsVersionUpdate
	fromVersion := utils.GetResultOrPanicError(r.versionAnalyzer.ReadVersion(c.GetRepoDir(), c.GetCurrentVersionAnalyzerArgumentValues()))
	repoManager.Checkout(c.GetToRevision())
	toVersion := utils.GetResultOrPanicError(r.versionAnalyzer.ReadVersion(c.GetRepoDir(), c.GetCurrentVersionAnalyzerArgumentValues()))

	fmt.Printf("TODO: Compare versions! From Version: %s To Version: %s\n", fromVersion, toVersion)
	fmt.Printf("Compare result: %d\n", r.versionComparer.Compare(fromVersion, toVersion))
}

func main() {
	r := runner{
		versionComparer: versionComparers.SemanticVersionComparer{},
	}
	r.run()
}
