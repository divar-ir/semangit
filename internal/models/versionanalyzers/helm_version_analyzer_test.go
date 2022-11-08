package versionanalyzers

import (
	"github.com/stretchr/testify/suite"
	"os"
	"semangit/internal/models"
	"semangit/internal/models/repo"
	"semangit/internal/utils"
	"testing"
)

type HelmVersionAnalyzerTestSuite struct {
	suite.Suite
	helmVersionAnalyzer HelmVersionAnalyzer
}

func TestHelmVersionAnalyzer(t *testing.T) {
	suite.Run(t, new(HelmVersionAnalyzerTestSuite))
}

func (s *HelmVersionAnalyzerTestSuite) SetupTest() {
	chartContent := `
apiVersion: v2
name: test-chart
description: A sample chart for testing purposes

# A chart can be either an 'application' or a 'library' chart.
#
# Application charts are a collection of templates that can be packaged into versioned archives
# to be deployed.
#
# Library charts provide useful utilities or functions for the chart developer. They're included as
# a dependency of application charts to inject those utilities and functions into the rendering
# pipeline. Library charts do not define any templates and therefore cannot be deployed.
type: application

# This is the chart version. This version number should be incremented each time you make changes
# to the chart and its templates, including the app version.
# Versions are expected to follow Semantic Versioning (https://semver.org/)
version: 1.2.3

# This is the version number of the application being deployed. This version number should be
# incremented each time you make changes to the application. Versions are not expected to
# follow Semantic Versioning. They should reflect the version the application is using.
# It is recommended to use it with quotes.
appVersion: "0.1.0"
`
	utils.PanicError(os.WriteFile("./Chart.yaml", []byte(chartContent), 0644))
	s.helmVersionAnalyzer = HelmVersionAnalyzer{}
}

func (s *HelmVersionAnalyzerTestSuite) TearDownTest() {
	utils.PanicError(os.Remove("./Chart.yaml"))
}

func (s *HelmVersionAnalyzerTestSuite) TestRepositoryContainsHelmVersionAnalyzer() {
	analyzer := repo.GetVersionAnalyzer(versionAnalyzerNameHelm)
	s.IsType(&HelmVersionAnalyzer{}, analyzer)
}

func (s *HelmVersionAnalyzerTestSuite) TestCanReadChartVersion() {
	helmRootDir := "."
	version := utils.GetResultOrPanic(s.helmVersionAnalyzer.ReadVersion(".", &models.ArgumentValues{
		ArgumentKeyRootDir: &helmRootDir,
	}))
	s.Equal("1.2.3", version)
}

func (s *HelmVersionAnalyzerTestSuite) TestVersionUpdateIsNotNeededWhenNoChangeIsMade() {
	helmRootDir := "."
	needsVersionUpdate := s.helmVersionAnalyzer.ChangeNeedsVersionUpdate([]string{}, &models.ArgumentValues{
		ArgumentKeyRootDir: &helmRootDir,
	})
	s.False(needsVersionUpdate)
}

func (s *HelmVersionAnalyzerTestSuite) TestVersionUpdateIsNotNeededWhenChangesAreOutsideHelmRootDir() {
	helmRootDir := "/some/project/helm/"
	needsVersionUpdate := s.helmVersionAnalyzer.ChangeNeedsVersionUpdate([]string{
		"/some/project/non-helm.txt",
	}, &models.ArgumentValues{
		ArgumentKeyRootDir: &helmRootDir,
	})
	s.False(needsVersionUpdate)
}

func (s *HelmVersionAnalyzerTestSuite) TestVersionUpdateIsNotNeededWhenChangesAreInsideHelmRootDirButOutsideTemplatesDir() {
	helmRootDir := "/some/project/helm/"
	needsVersionUpdate := s.helmVersionAnalyzer.ChangeNeedsVersionUpdate([]string{
		"/some/project/helm/Chart.yaml",
	}, &models.ArgumentValues{
		ArgumentKeyRootDir: &helmRootDir,
	})
	s.False(needsVersionUpdate)
}

func (s *HelmVersionAnalyzerTestSuite) TestVersionUpdateIsNeededWhenSomeChangesAreInsideHelmTemplatesDir() {
	helmRootDir := "/some/project/helm/"
	needsVersionUpdate := s.helmVersionAnalyzer.ChangeNeedsVersionUpdate([]string{
		"/some/project/helm/templates/deployment.yaml",
		"/some/project/non-helm.txt",
	}, &models.ArgumentValues{
		ArgumentKeyRootDir: &helmRootDir,
	})
	s.True(needsVersionUpdate)
}
