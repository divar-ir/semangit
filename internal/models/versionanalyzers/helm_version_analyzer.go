package versionanalyzers

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"semangit/internal/models"
	"semangit/internal/models/base"
	"semangit/internal/models/repo"
	"semangit/internal/utils"
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

func (a *HelmVersionAnalyzer) GetExtraArgumentDefinitions() []models.ArgumentDefinition {
	return []models.ArgumentDefinition{
		{
			Name:         ArgumentKeyRootDir,
			DefaultValue: ".",
			Description:  "The directory where the Chart.yaml exists.",
		},
	}
}

func (a *HelmVersionAnalyzer) ChangeNeedsVersionUpdate(changedFilesPaths []string, extraArgs *models.ArgumentValues) bool {
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

func (a *HelmVersionAnalyzer) ReadVersion(projectRootDir string, extraArgs *models.ArgumentValues) (string, error) {
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
