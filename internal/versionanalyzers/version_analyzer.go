package versionanalyzers

type VersionAnalyzer interface {
	GetName() string
	ChangeNeedsVersionUpdate(changedFilePaths []string, extraArgs *ArgumentValues) bool
	ReadVersion(projectRootPath string, extraArgs *ArgumentValues) (string, error)
	GetExtraArgumentDefinitions() []ArgumentDefinition
	CompareVersions(left string, right string) int
}
