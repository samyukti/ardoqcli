package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"com.samyukti.ardoqcli/internal/api"
	"com.samyukti.ardoqcli/internal/config"
)

// newClient creates an API client from the current config.
// Returns an error if base URL or API key are not configured.
func newClient() (*api.Client, error) {
	baseURL := config.BaseURL()
	apiKey := config.APIKey()

	if baseURL == "" || apiKey == "" {
		missing := []string{}
		if baseURL == "" {
			missing = append(missing, "base_url (ARDOQ_BASE_URL)")
		}
		if apiKey == "" {
			missing = append(missing, "api_key (ARDOQ_API_KEY)")
		}
		return nil, fmt.Errorf("missing config: %s\nRun 'ardoqcli configure' or set environment variables",
			strings.Join(missing, ", "))
	}

	return api.NewClient(baseURL, apiKey), nil
}

// parseQuery parses the -q flag value "key=val,key=val" into url.Values.
func parseQuery() url.Values {
	params := url.Values{}
	if queryFlag == "" {
		return params
	}
	for _, pair := range strings.Split(queryFlag, ",") {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			params.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
	return params
}
