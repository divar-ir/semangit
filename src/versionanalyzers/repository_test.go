package versionanalyzers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) TestRepositoryContainsHelmVersionAnalyzer() {
	analyzer := GetVersionAnalyzer(VersionAnalyzerNameHelm)
	assert.IsType(s.T(), &HelmVersionAnalyzer{}, analyzer)
}

type fakeAnalyzer struct {
	versionAnalyzer
}

func (a *fakeAnalyzer) GetName() string {
	return "test-fake-analyzer"
}

func (s *RepositoryTestSuite) TestCanRegisterNewVersionAnalyzer() {
	previousAnalyzersCount := len(GetAllAnalyzers())
	registerVersionAnalyzer(&fakeAnalyzer{})
	assert.Equal(s.T(), previousAnalyzersCount+1, len(GetAllAnalyzers()))
}

func (s *RepositoryTestSuite) TestCanNotRegisterVersionAnalyzerWithRepetitiveName() {
	assert.Error(s.T(), registerVersionAnalyzer(&HelmVersionAnalyzer{}))
}
