package cmd

import (
	"fmt"
	"os"

	"com.samyukti.ardoqcli/internal/output"
	"github.com/spf13/cobra"
)

var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Execute batch operations (POST /api/v2/batch)",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, _ := cmd.Flags().GetString("file")
		if file == "" {
			return fmt.Errorf("batch requires -f <file>")
		}

		body, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("read batch file: %w", err)
		}

		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Post("/api/v2/batch", body)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

func init() {
	batchCmd.Flags().StringP("file", "f", "", "Path to batch JSON file")
	batchCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(batchCmd)
}
