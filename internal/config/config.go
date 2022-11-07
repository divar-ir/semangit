package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"semangit/internal/versionanalyzers"
	"semangit/internal/versionanalyzers/repo"
	"strings"
)

const RevisionNone = ""

type Config struct {
	RepoDir                        string                                      `json:"repo_dir,omitempty"`
	SrcRevision                    string                                      `json:"src_revision,omitempty"`
	DestRevision                   string                                      `json:"dest_revision,omitempty"`
	CurrentVersionAnalyzerName     string                                      `json:"active_version_analyzer_name,omitempty""`
	VersionAnalyzersArgumentValues map[string]*versionanalyzers.ArgumentValues `json:"version_analyzers,omitempty"`
}

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	viper.SetDefault("RepoDir", ".")
	viper.SetDefault("SrcRevision", RevisionNone)
	viper.SetDefault("DestRevision", RevisionNone)
	viper.SetDefault("CurrentVersionAnalyzerName", "helm")
	viper.SetDefault("VersionAnalyzersArgumentValues", make(map[string]*versionanalyzers.ArgumentValues))

	// Read Config from ENV
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read Config from Flags
	if err := viper.BindPFlag("RepoDir", cmd.Flags().Lookup("repo-dir")); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := viper.BindPFlag("SrcRevision", cmd.Flags().Lookup("src-rev")); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := viper.BindPFlag("DestRevision", cmd.Flags().Lookup("dest-rev")); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := viper.BindPFlag("CurrentVersionAnalyzerName", cmd.Flags().Lookup("version-analyzer-name")); err != nil {
		return nil, errors.WithStack(err)
	}

	extractVersionAnalyzersArguments(cmd)

	// Read Config from file
	if configFile, err := cmd.Flags().GetString("config-file"); err == nil && configFile != "" {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}

func extractVersionAnalyzersArguments(cmd *cobra.Command) {
	versionAnalyzersConfigs := make(map[string]*versionanalyzers.ArgumentValues)
	for _, versionAnalyzer := range repo.GetAllAnalyzers() {
		argValues := make(versionanalyzers.ArgumentValues)

		argNamePrefix := versionAnalyzer.GetName() + "-"
		for _, argDefinition := range versionAnalyzer.GetExtraArgumentDefinitions() {
			argValue, err := cmd.Flags().GetString(argNamePrefix + argDefinition.Name)
			if err != nil {
				// TODO: add logger and write some error
				_ = fmt.Errorf("couldn't add %v flag: %v", argNamePrefix+argDefinition.Name, err.Error())
				continue
			}
			argValues[argDefinition.Name] = &argValue
		}
		versionAnalyzersConfigs[versionAnalyzer.GetName()] = &argValues
	}
	viper.Set("VersionAnalyzersArgumentValues", versionAnalyzersConfigs)
}

func (c *Config) GetCurrentVersionAnalyzerArgumentValues() *versionanalyzers.ArgumentValues {
	return c.VersionAnalyzersArgumentValues[c.CurrentVersionAnalyzerName]
}
