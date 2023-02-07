package base

import (
	"golang.org/x/mod/semver"
	"semangit/internal/models"
)

type BaseAnalyzer struct {
}

func newBaseAnalyzer() *BaseAnalyzer {
	return &BaseAnalyzer{}
}

// GetName Returns the name of the analyzer. This is the name that can be used in commandline to choose this analyzer.
func (a *BaseAnalyzer) GetName() string {
	panic("not implemented")
}

func (a *BaseAnalyzer) ChangeNeedsVersionUpdate(changedFilePaths []string, extraArgs *models.ArgumentValues) bool {
	panic("not implemented")
}

func (a *BaseAnalyzer) ReadVersion(projectRootDir string, extraArgs *models.ArgumentValues) (string, error) {
	panic("not implemented")
}

func (a *BaseAnalyzer) GetExtraArgumentDefinitions() []models.ArgumentDefinition {
	return []models.ArgumentDefinition{}
}

// CompareVersions Returns 0 if the two version are equal, negative if oldVersion < newVersion, and positive if oldVersion > newVersion.
func (a *BaseAnalyzer) CompareVersions(oldVersion string, newVersion string) int {
	if len(oldVersion) > 0 && oldVersion[0] != 'v' {
		oldVersion = "v" + oldVersion
	}
	if len(newVersion) > 0 && newVersion[0] != 'v' {
		newVersion = "v" + newVersion
	}
	return semver.Compare(oldVersion, newVersion)
}
