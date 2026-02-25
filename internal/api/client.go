package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client wraps HTTP communication with the Ardoq API.
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewClient creates an API client from base URL and API key.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL:    strings.TrimRight(baseURL, "/"),
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

// Get performs a GET request. params are added as query parameters.
func (c *Client) Get(path string, params url.Values) ([]byte, error) {
	u := c.BaseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return c.do("GET", u, nil)
}

// Post performs a POST request with a JSON body.
func (c *Client) Post(path string, body []byte) ([]byte, error) {
	return c.do("POST", c.BaseURL+path, body)
}

// Patch performs a PATCH request with a JSON body.
// It sets ifVersionMatch=latest by default.
func (c *Client) Patch(path string, body []byte) ([]byte, error) {
	u := c.BaseURL + path
	if !strings.Contains(u, "ifVersionMatch") {
		sep := "?"
		if strings.Contains(u, "?") {
			sep = "&"
		}
		u += sep + "ifVersionMatch=latest"
	}
	return c.do("PATCH", u, body)
}

// Delete performs a DELETE request.
func (c *Client) Delete(path string) ([]byte, error) {
	return c.do("DELETE", c.BaseURL+path, nil)
}

func (c *Client) do(method, url string, body []byte) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("User-Agent", "ardoqcli/0.1")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(data),
		}
	}

	return data, nil
}
