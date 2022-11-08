package base

import (
	"golang.org/x/mod/semver"
	"semangit/internal/models"
)

type BaseAnalyzer struct {
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

// CompareVersions Returns 0 if the two version are equal, negative if left < right, and positive if left > right.
func (a *BaseAnalyzer) CompareVersions(left string, right string) int {
	return semver.Compare("v"+left, "v"+right)
}
