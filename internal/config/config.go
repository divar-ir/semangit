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
const EnvsPrefix = "SEMANGIT"

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
	viper.SetDefault("LogLevel", "info")

	// Read Config from ENV
	viper.SetEnvPrefix(EnvsPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read Config from Flags
	if err := viper.BindPFlag("LogLevel", cmd.Flags().Lookup("log-level")); err != nil {
		logrus.WithError(err).Error()
		return nil, errors.WithStack(err)
	}

	logLevel := viper.GetString("LogLevel")
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
	default:
		logrus.Panic("Log level flag's value is invalid")
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

// extractVersionAnalyzersArguments function will try to read version analyzers' arguments values from flags
// and environment variables. The priority in which the values are read is as follows:
// environment variables < flags
// Environment variables are assumed to be in the following format:
// env: SEMANGIT_{VERSION_ANALYZER_NAME + ARGUMENT_NAME_WITHOUT_UNDERSCORE} (e.q. SEMANGIT_HELMROOTDIR)
// Flags are also assumed to be in the following format:
// flag: {VERSION_ANALYZER_NAME}-{ARGUMENT_NAME_WITH_UNDERSCORES_IN_LOWERCASE} (e.q. --helm-root-dir)
func extractVersionAnalyzersArguments(cmd *cobra.Command) {
	versionAnalyzersConfigs := make(map[string]*models.ArgumentValues)
	for _, versionAnalyzer := range repo.GetAllAnalyzers() {
		argValues := make(models.ArgumentValues)

		argNamePrefix := versionAnalyzer.GetName() + "-"
		for _, argDefinition := range versionAnalyzer.GetExtraArgumentDefinitions() {
			flagName := argNamePrefix + argDefinition.Name // helm-root-dir
			viperKey := strings.ToUpper(flagName)
			viperKey = strings.ReplaceAll(viperKey, "-", "")
			viperKey = strings.ReplaceAll(viperKey, ".", "") // HELMROOTDIR
			envKey := EnvsPrefix + "_" + viperKey            // SEMANGIT_HELMROOTDIR

			err := viper.BindEnv(envKey, viperKey)
			if err != nil {
				logrus.WithError(err).Error()
			}

			if err := viper.BindPFlag(viperKey, cmd.Flags().Lookup(flagName)); err != nil {
				logrus.WithError(err).Error()
			}

			argValue := viper.GetString(viperKey)
			argValues[argDefinition.Name] = &argValue
		}
		versionAnalyzersConfigs[versionAnalyzer.GetName()] = &argValues
	}
	viper.Set("VersionAnalyzersArgumentValues", versionAnalyzersConfigs)
}

func (c *Config) GetCurrentVersionAnalyzerArgumentValues() *models.ArgumentValues {
	return c.VersionAnalyzersArgumentValues[c.CurrentVersionAnalyzerName]
}
