package base

import (
	"golang.org/x/mod/semver"
	"semangit/internal/versionanalyzers"
)

type BaseAnalyzer struct {
}

// GetName Returns the name of the analyzer. This is the name that can be used in commandline to choose this analyzer.
func (a *BaseAnalyzer) GetName() string {
	panic("not implemented")
}

func (a *BaseAnalyzer) ChangeNeedsVersionUpdate(changedFilePaths []string, extraArgs *versionanalyzers.ArgumentValues) bool {
	panic("not implemented")
}

func (a *BaseAnalyzer) ReadVersion(projectRootDir string, extraArgs *versionanalyzers.ArgumentValues) (string, error) {
	panic("not implemented")
}

func (a *BaseAnalyzer) GetExtraArgumentDefinitions() []versionanalyzers.ArgumentDefinition {
	return []versionanalyzers.ArgumentDefinition{}
}

// CompareVersions Returns 0 if the two version are equal, negative if left < right, and positive if left > right.
func (a *BaseAnalyzer) CompareVersions(left string, right string) int {
	return semver.Compare(left, right)
}
