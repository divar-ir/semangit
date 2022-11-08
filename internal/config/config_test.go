package config

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
	"os"
	"semangit/internal/models/versionanalyzers"
	"semangit/internal/utils"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
	cmd *cobra.Command
}

func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) SetupSuite() {
	s.cmd = &cobra.Command{}
}

func (s *ConfigTestSuite) TearDownTest() {
	s.cmd.ResetFlags()
	if _, err := os.Stat("./config.env"); err == nil {
		utils.PanicError(os.Remove("./config.env"))
	}
}

func (s *ConfigTestSuite) AddRequiredFlags() {
	s.cmd = &cobra.Command{}
	s.cmd.Flags().String("repo-dir", "", "")
	s.cmd.Flags().String("old-rev", "", "")
	s.cmd.Flags().String("new-rev", "", "")
	s.cmd.Flags().String("version-analyzer-name", "", "")
	s.cmd.Flags().String("helm-root-dir", "", "")
}

func (s *ConfigTestSuite) TestExtraArguments() {
	s.AddRequiredFlags()
	s.NoError(s.cmd.Flags().Set("version-analyzer-name", "helm"))
	s.NoError(s.cmd.Flags().Set("helm-"+versionanalyzers.HelmArgumentKeyRootDir, "test-value"))
	conf, err := LoadConfig(s.cmd)
	s.NoError(err)
	s.Equal(*(*conf.GetCurrentVersionAnalyzerArgumentValues())[versionanalyzers.HelmArgumentKeyRootDir], "test-value")
}

func (s *ConfigTestSuite) TestNilFlags() {
	_, err := LoadConfig(s.cmd)
	s.Error(err)
	s.cmd.Flags().String("repo-dir", "", "")
	_, err = LoadConfig(s.cmd)
	s.Error(err)
	s.cmd.Flags().String("old-rev", "", "")
	_, err = LoadConfig(s.cmd)
	s.Error(err)
	s.cmd.Flags().String("new-rev", "", "")
	_, err = LoadConfig(s.cmd)
	s.Error(err)
	s.cmd.Flags().String("version-analyzer-name", "", "")
	_, err = LoadConfig(s.cmd)
	s.NoError(err)
	s.cmd.Flags().String("helm-root-dir", "", "")
	_, err = LoadConfig(s.cmd)
	s.NoError(err)
}

func (s *ConfigTestSuite) TestConfigFile() {
	s.AddRequiredFlags()
	configFile := `
	RepoDir=TEST_REPO_DIR
`
	utils.PanicError(os.WriteFile("./config.env", []byte(configFile), 0644))
	s.cmd.Flags().String("config-file", "./config.env", "")
	conf, err := LoadConfig(s.cmd)
	s.NoError(err)
	s.Equal(conf.RepoDir, "TEST_REPO_DIR")
}
