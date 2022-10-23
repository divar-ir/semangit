package versionAnalyzers

func GetVersionAnalyzer(name string) VersionAnalyzer {
	for _, analyzer := range GetAllAnalyzers() {
		if analyzer.GetName() == name {
			return analyzer
		}
	}
	panic("unknown version analyzer: " + name)
}

var versionAnalyzers []VersionAnalyzer

func GetAllAnalyzers() []VersionAnalyzer {
	return versionAnalyzers
}

func registerVersionAnalyzer(a VersionAnalyzer) {
	versionAnalyzers = append(versionAnalyzers, a)
}
