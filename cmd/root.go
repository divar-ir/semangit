/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"semangit/internal/plugins/helm"
	"semangit/internal/plugins/versionanalyzer"
	"semangit/internal/utils"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve semangit",
	Long:  `With semangit you can check your versions!`, // TODO: complete this description
	Run: func(cmd *cobra.Command, args []string) {
		if err := serveSemanGit(cmd, args); err != nil {
			logrus.WithError(err).Fatal("Failed to serve.")
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Register version analyzers
	utils.PanicError(versionanalyzer.RegisterVersionAnalyzer(helm.New()))

	rootCmd.Flags().StringP(
		"repo-dir",
		"r",
		".",
		"Specify repository address",
	)
	rootCmd.Flags().StringP(
		"src-rev",
		"s",
		"",
		"Specify source revision",
	)
	rootCmd.Flags().StringP(
		"dest-rev",
		"d",
		"",
		"Specify destination revision",
	)
	rootCmd.Flags().StringP(
		"version-analyzer-name",
		"n",
		"helm",
		"Specify the name of the version analyzer you want to use",
	)
	registerVersionAnalyzersArgumentsFlags(rootCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".semangit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".semangit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func registerVersionAnalyzersArgumentsFlags(cmd *cobra.Command) {
	for _, versionAnalyzer := range versionanalyzer.GetAllAnalyzers() {
		argNamePrefix := versionAnalyzer.GetName() + "-"
		for _, argDefinition := range versionAnalyzer.GetExtraArgumentDefinitions() {
			cmd.Flags().String(argNamePrefix+argDefinition.Name, argDefinition.DefaultValue, argDefinition.Description)
		}
	}
}
