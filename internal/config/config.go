package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const (
	appName    = "ardoqcli"
	configFile = "hosts.yml"
)

// ConfigDir returns the config directory path.
func ConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", appName)
}

// ConfigPath returns the full path to the config file.
func ConfigPath() string {
	return filepath.Join(ConfigDir(), configFile)
}

// Init sets up viper to read config and env vars.
func Init(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("hosts")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(ConfigDir())
	}

	viper.SetEnvPrefix("ARDOQ")
	viper.AutomaticEnv()

	// Map config keys to env vars
	viper.BindEnv("base_url", "ARDOQ_BASE_URL")
	viper.BindEnv("api_key", "ARDOQ_API_KEY")

	_ = viper.ReadInConfig()
}

// Save writes base_url and api_key to the config file.
func Save(baseURL, apiKey string) error {
	dir := ConfigDir()
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data := map[string]string{
		"base_url": baseURL,
		"api_key":  apiKey,
	}

	out, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	if err := os.WriteFile(ConfigPath(), out, 0o600); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}

// BaseURL returns the configured base URL.
func BaseURL() string {
	return viper.GetString("base_url")
}

// APIKey returns the configured API key.
func APIKey() string {
	return viper.GetString("api_key")
}
