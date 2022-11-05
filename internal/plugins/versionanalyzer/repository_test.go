package versionanalyzer

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"semangit/internal/plugins/base"
	"semangit/internal/plugins/helm"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) TestRepositoryContainsHelmVersionAnalyzer() {
	analyzer := GetVersionAnalyzer(helm.VersionAnalyzerNameHelm)
	assert.IsType(s.T(), &helm.HelmVersionAnalyzer{}, analyzer)
}

type fakeAnalyzer struct {
	base.BaseAnalyzer
}

func (a *fakeAnalyzer) GetName() string {
	return "test-fake-analyzer"
}

func (s *RepositoryTestSuite) TestCanRegisterNewVersionAnalyzer() {
	previousAnalyzersCount := len(GetAllAnalyzers())
	RegisterVersionAnalyzer(&fakeAnalyzer{})
	assert.Equal(s.T(), previousAnalyzersCount+1, len(GetAllAnalyzers()))
}

func (s *RepositoryTestSuite) TestCanNotRegisterVersionAnalyzerWithRepetitiveName() {
	s.NoError(RegisterVersionAnalyzer(helm.New()))
	s.Error(RegisterVersionAnalyzer(helm.New()))
}
