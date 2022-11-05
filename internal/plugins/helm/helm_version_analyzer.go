package helm

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"semangit/internal/plugins/base"
	"semangit/internal/utils"
)
import "strings"

type HelmVersionAnalyzer struct {
	base.BaseAnalyzer
}

const ArgumentKeyRootDir = "root-dir"
const VersionAnalyzerNameHelm = "helm"

func New() *HelmVersionAnalyzer {
	return &HelmVersionAnalyzer{}
}

func (a *HelmVersionAnalyzer) GetExtraArgumentDefinitions() []base.ArgumentDefinition {
	return []base.ArgumentDefinition{
		{
			Name:         ArgumentKeyRootDir,
			DefaultValue: ".",
			Description:  "The directory where the Chart.yaml exists.",
		},
	}
}

func (a *HelmVersionAnalyzer) ChangeNeedsVersionUpdate(changedFilesPaths []string, extraArgs *base.ArgumentValues) bool {
	helmRootDir := *(*extraArgs)[ArgumentKeyRootDir]
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

func (a *HelmVersionAnalyzer) ReadVersion(projectRootDir string, extraArgs *base.ArgumentValues) (string, error) {
	rootDir := *(*extraArgs)[ArgumentKeyRootDir]
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
