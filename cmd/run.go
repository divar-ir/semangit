package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"semangit/internal/config"
	"semangit/internal/gitrepo"
	"semangit/internal/models/repo"
	"semangit/internal/utils"
)

func runSemangit(cmd *cobra.Command, args []string) error {
	conf := utils.GetResultOrPanic(config.LoadConfig(cmd))
	versionAnalyzer := repo.GetVersionAnalyzer(conf.CurrentVersionAnalyzerName)
	repoManager := gitrepo.NewGitRepoManger(conf.RepoDir)

	if conf.SrcRevision != gitrepo.RevisionNone {
		repoManager.Checkout(conf.SrcRevision)
	}
	srcVersion := utils.GetResultOrPanic(versionAnalyzer.ReadVersion(conf.RepoDir, conf.GetCurrentVersionAnalyzerArgumentValues()))
	repoManager.Checkout(conf.DestRevision)
	destVersion := utils.GetResultOrPanic(versionAnalyzer.ReadVersion(conf.RepoDir, conf.GetCurrentVersionAnalyzerArgumentValues()))
	changedFiles := repoManager.ListChangedFiles(conf.SrcRevision, conf.DestRevision)
	needsUpdate := versionAnalyzer.ChangeNeedsVersionUpdate(changedFiles, conf.GetCurrentVersionAnalyzerArgumentValues())
	if needsUpdate && versionAnalyzer.CompareVersions(srcVersion, destVersion) <= 0 {
		return errors.New("Version needs to be updated! Version analyzer: " + versionAnalyzer.GetName())
	}
	return nil
}
