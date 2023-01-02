package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"semangit/internal/models"
	"semangit/internal/models/repo"
	"strings"
)

const RevisionNone = ""

type Config struct {
	RepoDir                        string                            `json:"repo_dir,omitempty"`
	OldRevision                    string                            `json:"old_revision,omitempty"`
	NewRevision                    string                            `json:"new_revision,omitempty"`
	CurrentVersionAnalyzerName     string                            `json:"active_version_analyzer_name,omitempty"`
	VersionAnalyzersArgumentValues map[string]*models.ArgumentValues `json:"version_analyzers,omitempty"`
}

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	viper.SetDefault("RepoDir", ".")
	viper.SetDefault("OldRevision", RevisionNone)
	viper.SetDefault("NewRevision", RevisionNone)
	viper.SetDefault("CurrentVersionAnalyzerName", "helm")
	viper.SetDefault("VersionAnalyzersArgumentValues", make(map[string]*models.ArgumentValues))

	// Read Config from ENV
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read Config from Flags
	logLevel := strings.ToLower(cmd.Flags().Lookup("log-level").Value.String())
	switch logLevel {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	}

	if err := viper.BindPFlag("RepoDir", cmd.Flags().Lookup("repo-dir")); err != nil {
		logrus.WithError(err).Error()
		return nil, errors.WithStack(err)
	}

	if err := viper.BindPFlag("OldRevision", cmd.Flags().Lookup("old-rev")); err != nil {
		logrus.WithError(err).Error()
		return nil, errors.WithStack(err)
	}

	if err := viper.BindPFlag("NewRevision", cmd.Flags().Lookup("new-rev")); err != nil {
		logrus.WithError(err).Error()
		return nil, errors.WithStack(err)
	}

	if err := viper.BindPFlag("CurrentVersionAnalyzerName", cmd.Flags().Lookup("version-analyzer-name")); err != nil {
		logrus.WithError(err).Error()
		return nil, errors.WithStack(err)
	}

	extractVersionAnalyzersArguments(cmd)

	// Read Config from file
	if configFile, err := cmd.Flags().GetString("config-file"); err == nil && configFile != "" {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			logrus.WithError(err).Error()
			return nil, errors.WithStack(err)
		}
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		logrus.WithError(err).Error()
		return nil, errors.WithStack(err)
	}

	return &config, nil
}

func extractVersionAnalyzersArguments(cmd *cobra.Command) {
	versionAnalyzersConfigs := make(map[string]*models.ArgumentValues)
	for _, versionAnalyzer := range repo.GetAllAnalyzers() {
		argValues := make(models.ArgumentValues)

		argNamePrefix := versionAnalyzer.GetName() + "-"
		for _, argDefinition := range versionAnalyzer.GetExtraArgumentDefinitions() {
			argValue, err := cmd.Flags().GetString(argNamePrefix + argDefinition.Name)
			if err != nil {
				logrus.Errorf("couldn't add %v flag: %v", argNamePrefix+argDefinition.Name, err.Error())
				continue
			}
			argValues[argDefinition.Name] = &argValue
		}
		versionAnalyzersConfigs[versionAnalyzer.GetName()] = &argValues
	}
	viper.Set("VersionAnalyzersArgumentValues", versionAnalyzersConfigs)
}

func (c *Config) GetCurrentVersionAnalyzerArgumentValues() *models.ArgumentValues {
	return c.VersionAnalyzersArgumentValues[c.CurrentVersionAnalyzerName]
}
