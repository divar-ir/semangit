package base

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BaseAnalyzerTestSuite struct {
	suite.Suite
	baseAnalyzer *BaseAnalyzer
}

func TestBaseAnalyzerTestSuite(t *testing.T) {
	suite.Run(t, new(BaseAnalyzerTestSuite))
}

func (s *BaseAnalyzerTestSuite) SetupSuite() {
	s.baseAnalyzer = newBaseAnalyzer()
}

func (s *BaseAnalyzerTestSuite) TestZeroIsReturnedForEqualVersions() {
	s.Equal(s.baseAnalyzer.CompareVersions("2.1.1", "2.1.1"), 0)
	s.Equal(s.baseAnalyzer.CompareVersions("2.1.13", "2.1.13"), 0)
	s.Equal(s.baseAnalyzer.CompareVersions("2.14.1", "2.14.1"), 0)
	s.Equal(s.baseAnalyzer.CompareVersions("21.1.1", "21.1.1"), 0)
}

func (s *BaseAnalyzerTestSuite) TestCanCompareVersionsWithPrefixV() {
	s.Equal(s.baseAnalyzer.CompareVersions("v2.1.1", "2.1.1"), 0)
	s.Equal(s.baseAnalyzer.CompareVersions("v2.1.1", "v2.11.1"), -1)
	s.Equal(s.baseAnalyzer.CompareVersions("21.1.1", "v2.1.1"), 1)
	s.Equal(s.baseAnalyzer.CompareVersions("v2.1.1", "v2.11.1"), -1)
}

func (s *BaseAnalyzerTestSuite) TestCanCompareMajorVersion() {
	s.Equal(s.baseAnalyzer.CompareVersions("2.1.13", "21.1.13"), -1)
	s.Equal(s.baseAnalyzer.CompareVersions("21.1.13", "2.1.13"), 1)
}

func (s *BaseAnalyzerTestSuite) TestCanCompareMinorVersion() {
	s.Equal(s.baseAnalyzer.CompareVersions("2.1.13", "2.13.13"), -1)
	s.Equal(s.baseAnalyzer.CompareVersions("2.13.13", "2.1.13"), 1)
}

func (s *BaseAnalyzerTestSuite) TestCanComparePatchVersion() {
	s.Equal(s.baseAnalyzer.CompareVersions("2.1.1", "2.1.13"), -1)
	s.Equal(s.baseAnalyzer.CompareVersions("2.1.13", "2.1.1"), 1)
}
