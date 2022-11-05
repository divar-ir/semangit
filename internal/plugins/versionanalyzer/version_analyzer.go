package versionanalyzer

import "semangit/internal/plugins/base"

type VersionAnalyzer interface {
	GetName() string
	ChangeNeedsVersionUpdate(changedFilePaths []string, extraArgs *base.ArgumentValues) bool
	ReadVersion(projectRootPath string, extraArgs *base.ArgumentValues) (string, error)
	GetExtraArgumentDefinitions() []base.ArgumentDefinition
	CompareVersions(left string, right string) int
}
