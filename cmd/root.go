package cmd

import (
	"os"

	"com.samyukti.ardoqcli/internal/config"
	"com.samyukti.ardoqcli/internal/output"
	"github.com/spf13/cobra"
)

var (
	outFile    string
	queryFlag  string
	configFlag string
)

var rootCmd = &cobra.Command{
	Use:           "ardoqcli",
	Short:         "CLI for Ardoq REST API v2",
	Long:          "A command-line tool to operate Ardoq from the terminal.",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init(configFlag)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		output.Error("%v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outFile, "output", "o", "", "Write JSON output to file")
	rootCmd.PersistentFlags().StringVarP(&queryFlag, "query", "q", "", "Query parameters as key=val,key=val")
	rootCmd.PersistentFlags().StringVar(&configFlag, "config", "", "Path to config file")
}
