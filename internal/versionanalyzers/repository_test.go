package versionanalyzers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"semangit/internal/versionanalyzers/base"
	"semangit/internal/versionanalyzers/helm"
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
	assert.Error(s.T(), RegisterVersionAnalyzer(&helm.HelmVersionAnalyzer{}))
}
