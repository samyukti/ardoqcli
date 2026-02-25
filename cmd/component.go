package cmd

import (
	"fmt"

	"com.samyukti.ardoqcli/internal/output"
	"github.com/spf13/cobra"
)

var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Manage components",
}

var componentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List components",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get("/api/v2/components", parseQuery())
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var componentGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a component by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/api/v2/components/%s", args[0]), nil)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var componentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a component",
	RunE: func(cmd *cobra.Command, args []string) error {
		body, err := resolveInput(cmd)
		if err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Post("/api/v2/components", body)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var componentUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a component",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body, err := resolveInput(cmd)
		if err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Patch(fmt.Sprintf("/api/v2/components/%s", args[0]), body)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var componentDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a component",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(fmt.Sprintf("/api/v2/components/%s", args[0]))
		if err != nil {
			return err
		}
		if len(data) > 0 {
			return output.JSON(data, outFile)
		}
		output.Success("Component %s deleted", args[0])
		return nil
	},
}

func init() {
	componentCreateCmd.Flags().StringP("data", "d", "", "Inline JSON data")
	componentCreateCmd.Flags().StringP("file", "f", "", "Path to input file")
	componentCreateCmd.Flags().StringP("type", "t", "json", "Input format: json or csv")

	componentUpdateCmd.Flags().StringP("data", "d", "", "Inline JSON data")
	componentUpdateCmd.Flags().StringP("file", "f", "", "Path to input file")

	componentCmd.AddCommand(componentListCmd, componentGetCmd, componentCreateCmd, componentUpdateCmd, componentDeleteCmd)
	rootCmd.AddCommand(componentCmd)
}
