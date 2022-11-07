package repo

import (
	"fmt"
	"semangit/internal/versionanalyzers"
)

func GetVersionAnalyzer(name string) versionanalyzers.VersionAnalyzer {
	for _, analyzer := range GetAllAnalyzers() {
		if analyzer.GetName() == name {
			return analyzer
		}
	}
	panic("unknown version analyzer: " + name) // FIXME: return error instead of panic
}

var versionAnalyzers []versionanalyzers.VersionAnalyzer

func GetAllAnalyzers() []versionanalyzers.VersionAnalyzer {
	return versionAnalyzers
}

func RegisterVersionAnalyzer(a versionanalyzers.VersionAnalyzer) error {
	for _, analyzer := range versionAnalyzers {
		if analyzer.GetName() == a.GetName() {
			return fmt.Errorf("version analyzer with the name '%s' is already registered", a.GetName())
		}
	}
	versionAnalyzers = append(versionAnalyzers, a)
	return nil
}

func RemoveVersionAnalyzerIfExists(a versionanalyzers.VersionAnalyzer) {
	for i, analyzer := range versionAnalyzers {
		if analyzer.GetName() == a.GetName() {
			versionAnalyzers[i] = versionAnalyzers[len(versionAnalyzers)-1]
			versionAnalyzers[len(versionAnalyzers)-1] = nil
			versionAnalyzers = versionAnalyzers[:len(versionAnalyzers)-1]
			return
		}
	}
}
