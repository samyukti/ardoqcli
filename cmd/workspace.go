package cmd

import (
	"fmt"

	"com.samyukti.ardoqcli/internal/output"
	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces",
}

var workspaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get("/api/v2/workspaces", parseQuery())
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var workspaceGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a workspace by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/api/v2/workspaces/%s", args[0]), nil)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var workspaceContextCmd = &cobra.Command{
	Use:   "context <id>",
	Short: "Get workspace context (components, references, etc.)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/api/v2/workspaces/%s/context", args[0]), nil)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceListCmd, workspaceGetCmd, workspaceContextCmd)
	rootCmd.AddCommand(workspaceCmd)
}
