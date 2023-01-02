package cmd

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"semangit/internal/config"
	"semangit/internal/gitrepo"
	"semangit/internal/models/repo"
	"semangit/internal/utils"
)

func runSemangit(cmd *cobra.Command, args []string) error {
	conf := utils.GetResultOrPanic(config.LoadConfig(cmd))
	logrus.Debug("Loaded config:\n", utils.InterfaceToString(conf))
	versionAnalyzer := repo.GetVersionAnalyzer(conf.CurrentVersionAnalyzerName)
	repoManager := gitrepo.NewGitRepoManger(conf.RepoDir)

	if conf.NewRevision != gitrepo.RevisionNone {
		repoManager.Checkout(conf.NewRevision)
		logrus.Debug("Successfully checked out the new revision ", conf.NewRevision)
	}

	newVersion := utils.GetResultOrPanic(versionAnalyzer.ReadVersion(conf.RepoDir, conf.GetCurrentVersionAnalyzerArgumentValues()))
	logrus.Info("New Version: ", newVersion)

	repoManager.Checkout(conf.OldRevision)
	logrus.Debug("Successfully checked out the old revision ", conf.OldRevision)
	oldVersion := utils.GetResultOrPanic(versionAnalyzer.ReadVersion(conf.RepoDir, conf.GetCurrentVersionAnalyzerArgumentValues()))
	logrus.Info("Old Version: ", oldVersion)

	changedFiles := repoManager.ListChangedFiles(conf.OldRevision, conf.NewRevision)
	logrus.Debug("Changed files: ", utils.InterfaceToString(changedFiles))

	needsUpdate := versionAnalyzer.ChangeNeedsVersionUpdate(changedFiles, conf.GetCurrentVersionAnalyzerArgumentValues())
	logrus.Debug("Does the version need update: ", needsUpdate)
	if needsUpdate && versionAnalyzer.CompareVersions(oldVersion, newVersion) >= 0 {
		return errors.New("Version needs to be updated! Version analyzer: " + versionAnalyzer.GetName())
	}
	return nil
}
