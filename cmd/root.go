package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"semangit/internal/models/repo"
	_ "semangit/internal/models/versionanalyzers"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Run semangit",
	Long:  `A simple tool to force version update in CI.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runSemangit(cmd, args); err != nil {
			logrus.WithError(err).Fatal("Failed to run.")
		}
	},
}

func InitializeFlags() {
	rootCmd.Flags().StringP(
		"repo-dir",
		"r",
		".",
		"The repository directory",
	)
	rootCmd.Flags().StringP(
		"src-rev",
		"s",
		"",
		"The source git revision",
	)
	rootCmd.Flags().StringP(
		"dest-rev",
		"d",
		"",
		"The destination git revision",
	)
	rootCmd.Flags().StringP(
		"version-analyzer-name",
		"n",
		"helm",
		"The name of the version analyzer you want to use",
	)
	registerVersionAnalyzersArgumentsFlags(rootCmd)
}

func Execute() {
	InitializeFlags()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func registerVersionAnalyzersArgumentsFlags(cmd *cobra.Command) {
	for _, versionAnalyzer := range repo.GetAllAnalyzers() {
		argNamePrefix := versionAnalyzer.GetName() + "-"
		for _, argDefinition := range versionAnalyzer.GetExtraArgumentDefinitions() {
			cmd.Flags().String(argNamePrefix+argDefinition.Name, argDefinition.DefaultValue, argDefinition.Description)
		}
	}
}
