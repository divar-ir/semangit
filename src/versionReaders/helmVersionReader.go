package versionReaders

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const VersionReaderNameHelm = "helm"

type HelmVersionReader struct {
	baseReader
}

const ArgumentKeyRootDir = "root-dir"

type helmChart struct {
	Version string `yaml:"version"`
}

func newHelmVersionReader() HelmVersionReader {
	return HelmVersionReader{}
}

func (r *HelmVersionReader) GetName() string {
	return VersionReaderNameHelm
}

func (r *HelmVersionReader) GetArgumentDefinitions() []argumentDefinition {
	return []argumentDefinition{
		{
			Name:         ArgumentKeyRootDir,
			DefaultValue: ".",
			Help:         "The directory where the Chart.yaml exists.",
		},
	}
}

func (r *HelmVersionReader) ReadVersion(arguments *ArgumentValues) (string, error) {
	rootDir := *(*arguments)[ArgumentKeyRootDir]
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
