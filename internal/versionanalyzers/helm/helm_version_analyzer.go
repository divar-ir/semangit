package helm

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"semangit/internal/utils"
	"semangit/internal/versionanalyzers"
	"semangit/internal/versionanalyzers/base"
)
import "strings"

type HelmVersionAnalyzer struct {
	base.BaseAnalyzer
}

const argumentKeyRootDir = "root-dir"
const VersionAnalyzerNameHelm = "helm"

func New() *HelmVersionAnalyzer {
	return &HelmVersionAnalyzer{}
}

func (a *HelmVersionAnalyzer) GetExtraArgumentDefinitions() []versionanalyzers.ArgumentDefinition {
	return []versionanalyzers.ArgumentDefinition{
		{
			Name:         argumentKeyRootDir,
			DefaultValue: ".",
			Description:  "The directory where the Chart.yaml exists.",
		},
	}
}

func (a *HelmVersionAnalyzer) ChangeNeedsVersionUpdate(changedFilesPaths []string, extraArgs *versionanalyzers.ArgumentValues) bool {
	helmRootDir := *(*extraArgs)[argumentKeyRootDir]
	helmRootDir = utils.GetResultOrPanic(filepath.Abs(helmRootDir))
	helmTemplatesRootDir := filepath.Join(helmRootDir, "templates")
	for _, path := range changedFilesPaths {
		if strings.HasPrefix(path, helmTemplatesRootDir) {
			return true
		}
	}
	return false
}

type helmChart struct {
	Version string `yaml:"version"`
}

func (a *HelmVersionAnalyzer) ReadVersion(projectRootDir string, extraArgs *versionanalyzers.ArgumentValues) (string, error) {
	rootDir := *(*extraArgs)[argumentKeyRootDir]
	chartFileContent, err := os.ReadFile(filepath.Join(rootDir, "Chart.yaml"))
	if err != nil {
		return "", err
	}
	var chart helmChart
	err = yaml.Unmarshal(chartFileContent, &chart)
	if err != nil {
		return "", err
	}
	return chart.Version, nil
}

func (a *HelmVersionAnalyzer) GetName() string {
	return VersionAnalyzerNameHelm
}
