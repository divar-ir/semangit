package versionanalyzers

import "golang.org/x/mod/semver"

type versionAnalyzer struct {
}

type VersionAnalyzer interface {
	GetName() string
	ChangeNeedsVersionUpdate(changedFilePaths []string, extraArgs *ArgumentValues) bool
	ReadVersion(projectRootPath string, extraArgs *ArgumentValues) (string, error)
	GetExtraArgumentDefinitions() []ArgumentDefinition
	CompareVersions(left string, right string) int
}

// GetName Returns the name of the analyzer. This is the name that can be used in commandline to choose this analyzer.
func (a *versionAnalyzer) GetName() string {
	panic("not implemented")
}

func (a *versionAnalyzer) ChangeNeedsVersionUpdate(changedFilePaths []string, extraArgs *ArgumentValues) bool {
	panic("not implemented")
}

func (a *versionAnalyzer) ReadVersion(projectRootDir string, extraArgs *ArgumentValues) (string, error) {
	panic("not implemented")
}

func (a *versionAnalyzer) GetExtraArgumentDefinitions() []ArgumentDefinition {
	return []ArgumentDefinition{}
}

// CompareVersions Returns 0 if the two version are equal, negative if left < right, and positive if left > right.
func (a *versionAnalyzer) CompareVersions(left string, right string) int {
	return semver.Compare(left, right)
}
