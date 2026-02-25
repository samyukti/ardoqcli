package cmd

import (
	"fmt"
	"net/url"

	"com.samyukti.ardoqcli/internal/output"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Manage reports",
}

var reportListCmd = &cobra.Command{
	Use:   "list",
	Short: "List reports",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get("/api/v2/reports", parseQuery())
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var reportGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a report by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/api/v2/reports/%s", args[0]), nil)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var reportRunCmd = &cobra.Command{
	Use:   "run <id>",
	Short: "Run a report",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		reportType, _ := cmd.Flags().GetString("type")
		params := url.Values{}
		if reportType != "" {
			params.Set("type", reportType)
		}
		data, err := client.Get(fmt.Sprintf("/api/v2/reports/%s/run", args[0]), params)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

func init() {
	reportRunCmd.Flags().StringP("type", "t", "", "Report type: objects or tabular")
	reportCmd.AddCommand(reportListCmd, reportGetCmd, reportRunCmd)
	rootCmd.AddCommand(reportCmd)
}
