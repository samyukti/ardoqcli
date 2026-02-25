package api

import "fmt"

// APIError represents an HTTP error response from the Ardoq API.
type APIError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *APIError) Error() string {
	if e.Body != "" {
		return fmt.Sprintf("API error %d %s: %s", e.StatusCode, e.Status, e.Body)
	}
	return fmt.Sprintf("API error %d %s", e.StatusCode, e.Status)
}
