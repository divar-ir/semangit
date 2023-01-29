package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	s.cmd.Flags().String("log-level", "info", "")
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
	s.cmd.Flags().String("log-level", "info", "")
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

func (s *ConfigTestSuite) TestLogLevelFlag() {
	s.AddRequiredFlags()
	s.setLevelAndTestLogLevel("trace", logrus.TraceLevel)
	s.setLevelAndTestLogLevel("debug", logrus.DebugLevel)
	s.setLevelAndTestLogLevel("info", logrus.InfoLevel)
	s.setLevelAndTestLogLevel("warn", logrus.WarnLevel)
	s.setLevelAndTestLogLevel("error", logrus.ErrorLevel)
	s.setLevelAndTestLogLevel("fatal", logrus.FatalLevel)
	s.setLevelAndTestLogLevel("panic", logrus.PanicLevel)
}

func (s *ConfigTestSuite) setLevelAndTestLogLevel(logLevel string, desiredLevel logrus.Level) {
	err := s.cmd.Flags().Set("log-level", logLevel)
	s.NoError(err)
	_, err = LoadConfig(s.cmd)
	s.NoError(err)
	s.Equal(logrus.GetLevel(), desiredLevel)
}

func (s *ConfigTestSuite) TestReadConfigFromEnv() {
	s.AddRequiredFlags()
	s.NoError(os.Setenv("SEMANGIT_REPODIR", "./src"))
	s.NoError(os.Setenv("SEMANGIT_OLDREVISION", "master"))
	s.NoError(os.Setenv("SEMANGIT_NEWREVISION", "my-branch"))
	s.NoError(os.Setenv("SEMANGIT_CURRENTVERSIONANALYZERNAME", "helm"))
	s.NoError(os.Setenv("SEMANGIT_LOGLEVEL", "debug"))
	s.NoError(os.Setenv("SEMANGIT_HELMROOTDIR", "./helm"))
	conf, err := LoadConfig(s.cmd)
	s.NoError(err)
	s.Equal(viper.GetString("RepoDir"), "./src")
	s.Equal(viper.GetString("OldRevision"), "master")
	s.Equal(viper.GetString("NewRevision"), "my-branch")
	s.Equal(viper.GetString("CurrentVersionAnalyzerName"), "helm")
	s.Equal(viper.GetString("LogLevel"), "debug")
	s.Equal(*(*conf.VersionAnalyzersArgumentValues["helm"])["root-dir"], "./helm")
}
