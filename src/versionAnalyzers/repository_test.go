package versionAnalyzers

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
