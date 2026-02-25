package cmd

import (
	"com.samyukti.ardoqcli/internal/config"
	"com.samyukti.ardoqcli/internal/output"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Interactive setup for Ardoq connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		var baseURL, apiKey string

		// Set defaults from existing config
		baseURL = config.BaseURL()
		apiKey = config.APIKey()

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Ardoq Base URL").
					Description("e.g. https://myorg.ardoq.com").
					Value(&baseURL),
				huh.NewInput().
					Title("API Key").
					Description("Your Ardoq API token").
					EchoMode(huh.EchoModePassword).
					Value(&apiKey),
			),
		)

		if err := form.Run(); err != nil {
			return err
		}

		if err := config.Save(baseURL, apiKey); err != nil {
			return err
		}

		output.Success("Configuration saved to %s", config.ConfigPath())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
