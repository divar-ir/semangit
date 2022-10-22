package versionAnalyzers

type versionAnalyzer struct {
}

type VersionAnalyzer interface {
	ChangeNeedsVersionUpdate(changedFilePaths []string, extraArgs *ArgumentValues) bool
	ReadVersion(projectRootPath string, extraArgs *ArgumentValues) (string, error)
	GetExtraArgumentDefinitions() []ArgumentDefinition
	GetName() string
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

func (a *versionAnalyzer) GetName() string {
	panic("not implemented")
}
