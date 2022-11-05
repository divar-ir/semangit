package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"semangit/internal/config"
	"semangit/internal/gitrepo"
	"semangit/internal/plugins/versionanalyzer"
	"semangit/internal/utils"
)

func serveSemanGit(cmd *cobra.Command, args []string) error {
	conf := utils.GetResultOrPanic(config.LoadConfig(cmd))
	versionAnalyzer := versionanalyzer.GetVersionAnalyzer(conf.ActiveVersionAnalyzerName)
	repoManager := gitrepo.NewGitRepoManger(conf.RepoDir)

	if conf.SrcRevision != gitrepo.RevisionNone {
		repoManager.Checkout(conf.SrcRevision)
	}

	// TODO: Check if r.versionAnalyzer.ChangeNeedsVersionUpdate
	srcVersion := utils.GetResultOrPanic(versionAnalyzer.ReadVersion(conf.RepoDir, conf.GetCurrentVersionAnalyzerArgumentValues()))
	repoManager.Checkout(conf.DestRevision)
	toVersion := utils.GetResultOrPanic(versionAnalyzer.ReadVersion(conf.RepoDir, conf.GetCurrentVersionAnalyzerArgumentValues()))

	fmt.Printf("TODO: Compare versions! From Version: %s To Version: %s\n", srcVersion, toVersion)
	fmt.Printf("Compare result: %d\n", versionAnalyzer.CompareVersions(srcVersion, toVersion))
	return nil
}
