package helm

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"semangit/internal/utils"
	"semangit/internal/versionanalyzers"
	"semangit/internal/versionanalyzers/base"
	"semangit/internal/versionanalyzers/repo"
)
import "strings"

type HelmVersionAnalyzer struct {
	base.BaseAnalyzer
}

const ArgumentKeyRootDir = "root-dir"
const versionAnalyzerNameHelm = "helm"

func New() *HelmVersionAnalyzer {
	return &HelmVersionAnalyzer{}
}

func init() {
	utils.PanicError(repo.RegisterVersionAnalyzer(New()))
}

func (a *HelmVersionAnalyzer) GetExtraArgumentDefinitions() []versionanalyzers.ArgumentDefinition {
	return []versionanalyzers.ArgumentDefinition{
		{
			Name:         ArgumentKeyRootDir,
			DefaultValue: ".",
			Description:  "The directory where the Chart.yaml exists.",
		},
	}
}

func (a *HelmVersionAnalyzer) ChangeNeedsVersionUpdate(changedFilesPaths []string, extraArgs *versionanalyzers.ArgumentValues) bool {
	helmRootDir := *(*extraArgs)[ArgumentKeyRootDir]
	helmRootDir = utils.GetResultOrPanic(filepath.Abs(helmRootDir))

	helmTemplatesRootDir := filepath.Join(helmRootDir, "templates")
	for _, path := range changedFilesPaths {
		absPath := utils.GetResultOrPanic(filepath.Abs(path))
		valuesPath := filepath.Join(helmRootDir, "values.yaml")
		if strings.HasPrefix(path, helmTemplatesRootDir) || absPath == valuesPath {
			return true
		}
	}
	return false
}

type helmChart struct {
	Version string `yaml:"version"`
}

func (a *HelmVersionAnalyzer) ReadVersion(projectRootDir string, extraArgs *versionanalyzers.ArgumentValues) (string, error) {
	rootDir := filepath.Join(projectRootDir, *(*extraArgs)[ArgumentKeyRootDir])
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
	return versionAnalyzerNameHelm
}
