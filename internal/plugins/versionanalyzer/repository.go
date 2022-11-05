package versionanalyzer

import (
	"fmt"
)

func GetVersionAnalyzer(name string) VersionAnalyzer {
	for _, analyzer := range GetAllAnalyzers() {
		if analyzer.GetName() == name {
			return analyzer
		}
	}
	panic("unknown version analyzer: " + name) // FIXME: return error instead of panic
}

var versionAnalyzers []VersionAnalyzer

func GetAllAnalyzers() []VersionAnalyzer {
	return versionAnalyzers
}

func RegisterVersionAnalyzer(a VersionAnalyzer) error {
	for _, analyzer := range versionAnalyzers {
		if analyzer.GetName() == a.GetName() {
			return fmt.Errorf("version analyzer with the name '%s' is already registered", a.GetName())
		}
	}
	versionAnalyzers = append(versionAnalyzers, a)
	return nil
}
