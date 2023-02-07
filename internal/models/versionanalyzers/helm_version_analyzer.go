package versionanalyzers

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

const HelmArgumentKeyRootDir = "root-dir"
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
			Name:         HelmArgumentKeyRootDir,
			DefaultValue: ".",
			Description:  "The directory where the Chart.yaml exists.",
		},
	}
}

func (a *HelmVersionAnalyzer) ChangeNeedsVersionUpdate(changedFilesPaths []string, extraArgs *models.ArgumentValues) bool {
	helmRootDir := *(*extraArgs)[HelmArgumentKeyRootDir]
	helmRootDir = utils.GetResultOrPanic(filepath.Abs(helmRootDir))

	helmTemplatesRootDir := filepath.Join(helmRootDir, "templates")
	valuesPath := filepath.Join(helmRootDir, "values.yaml")
	for _, path := range changedFilesPaths {
		absPath := utils.GetResultOrPanic(filepath.Abs(path))
		if strings.HasPrefix(absPath, helmTemplatesRootDir) || absPath == valuesPath {
			return true
		}
	}
	return false
}

type helmChart struct {
	Version string `yaml:"version"`
}

func (a *HelmVersionAnalyzer) ReadVersion(projectRootDir string, extraArgs *models.ArgumentValues) (string, error) {
	rootDir := filepath.Join(projectRootDir, *(*extraArgs)[HelmArgumentKeyRootDir])
	if _, err := os.Stat(filepath.Join(rootDir, "Chart.yaml")); errors.Is(err, os.ErrNotExist) {
		return "", nil
	}
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

func (a *HelmVersionAnalyzer) CompareVersions(oldVersion, newVersion string) int {
	if oldVersion == "" {
		logrus.Info("old revision has no 'Chart.yaml' file")
		return -1
	}
	return a.BaseAnalyzer.CompareVersions(oldVersion, newVersion)
}
