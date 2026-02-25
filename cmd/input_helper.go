package cmd

import (
	"com.samyukti.ardoqcli/internal/input"
	"github.com/spf13/cobra"
)

// resolveInput reads the -d, -f, and -t flags from a command and returns JSON bytes.
func resolveInput(cmd *cobra.Command) ([]byte, error) {
	data, _ := cmd.Flags().GetString("data")
	file, _ := cmd.Flags().GetString("file")
	format, _ := cmd.Flags().GetString("type")
	if format == "" {
		format = "json"
	}
	return input.Resolve(data, file, format)
}
