package cmd

import (
	"com.samyukti.ardoqcli/internal/output"
	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Test connection to Ardoq (GET /api/v2/me)",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get("/api/v2/me", nil)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

func init() {
	rootCmd.AddCommand(meCmd)
}
