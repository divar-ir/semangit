package versionAnalyzers

import (
	"go/importer"
	"reflect"
	"semangit/src/utils"
)

func GetVersionAnalyzer(name string) VersionAnalyzer {
	for _, analyzer := range GetAllAnalyzers() {
		if analyzer.GetName() == name {
			return analyzer
		}
	}
	panic("unknown version analyzer: " + name)
}

func GetAllAnalyzers() []VersionAnalyzer {
	// TODO: Use native plugin system and import .so shared libraries.
	currentPackageName := reflect.TypeOf(GetAllAnalyzers).PkgPath()
	pkg := utils.GetResultOrPanicError(importer.Default().Import(currentPackageName))
	scope := pkg.Scope()
	var analyzers []VersionAnalyzer
	for _, symbolName := range scope.Names() {
		symbol := scope.Lookup(symbolName)
		if analyzer, ok := symbol.(VersionAnalyzer); ok {
			analyzers = append(analyzers, analyzer)
		}
	}
	return analyzers
}
