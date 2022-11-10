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

	if conf.NewRevision != gitrepo.RevisionNone {
		repoManager.Checkout(conf.NewRevision)
	}
	newVersion := utils.GetResultOrPanic(versionAnalyzer.ReadVersion(conf.RepoDir, conf.GetCurrentVersionAnalyzerArgumentValues()))
	repoManager.Checkout(conf.OldRevision)
	oldVersion := utils.GetResultOrPanic(versionAnalyzer.ReadVersion(conf.RepoDir, conf.GetCurrentVersionAnalyzerArgumentValues()))
	changedFiles := repoManager.ListChangedFiles(conf.OldRevision, conf.NewRevision)
	needsUpdate := versionAnalyzer.ChangeNeedsVersionUpdate(changedFiles, conf.GetCurrentVersionAnalyzerArgumentValues())
	if needsUpdate && versionAnalyzer.CompareVersions(oldVersion, newVersion) >= 0 {
		return errors.New("Version needs to be updated! Version analyzer: " + versionAnalyzer.GetName())
	}
	return nil
}
