package repo

import (
	"github.com/stretchr/testify/suite"
	"semangit/internal/versionanalyzers/base"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
}

func (*RepositoryTestSuite) TearDownTest() {
	RemoveVersionAnalyzerIfExists(&fakeAnalyzer{})
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

type fakeAnalyzer struct {
	base.BaseAnalyzer
}

func (a *fakeAnalyzer) GetName() string {
	return "test-fake-analyzer"
}

func (s *RepositoryTestSuite) TestCanRegisterNewVersionAnalyzer() {
	previousAnalyzersCount := len(GetAllAnalyzers())
	s.NoError(RegisterVersionAnalyzer(&fakeAnalyzer{}))
	s.Equal(previousAnalyzersCount+1, len(GetAllAnalyzers()))
}

func (s *RepositoryTestSuite) TestCanNotRegisterVersionAnalyzerWithRepetitiveName() {
	s.NoError(RegisterVersionAnalyzer(&fakeAnalyzer{}))
	s.Error(RegisterVersionAnalyzer(&fakeAnalyzer{}))
}
