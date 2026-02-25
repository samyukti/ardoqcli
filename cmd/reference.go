package cmd

import (
	"fmt"

	"com.samyukti.ardoqcli/internal/output"
	"github.com/spf13/cobra"
)

var referenceCmd = &cobra.Command{
	Use:   "reference",
	Short: "Manage references",
}

var referenceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List references",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get("/api/v2/references", parseQuery())
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var referenceGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a reference by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/api/v2/references/%s", args[0]), nil)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var referenceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a reference",
	RunE: func(cmd *cobra.Command, args []string) error {
		body, err := resolveInput(cmd)
		if err != nil {
			return err
		}
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Post("/api/v2/references", body)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var referenceUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a reference",
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
		data, err := client.Patch(fmt.Sprintf("/api/v2/references/%s", args[0]), body)
		if err != nil {
			return err
		}
		return output.JSON(data, outFile)
	},
}

var referenceDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a reference",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(fmt.Sprintf("/api/v2/references/%s", args[0]))
		if err != nil {
			return err
		}
		if len(data) > 0 {
			return output.JSON(data, outFile)
		}
		output.Success("Reference %s deleted", args[0])
		return nil
	},
}

func init() {
	referenceCreateCmd.Flags().StringP("data", "d", "", "Inline JSON data")
	referenceCreateCmd.Flags().StringP("file", "f", "", "Path to input file")
	referenceCreateCmd.Flags().StringP("type", "t", "json", "Input format: json or csv")

	referenceUpdateCmd.Flags().StringP("data", "d", "", "Inline JSON data")
	referenceUpdateCmd.Flags().StringP("file", "f", "", "Path to input file")

	referenceCmd.AddCommand(referenceListCmd, referenceGetCmd, referenceCreateCmd, referenceUpdateCmd, referenceDeleteCmd)
	rootCmd.AddCommand(referenceCmd)
}
