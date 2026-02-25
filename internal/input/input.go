package input

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Resolve returns JSON bytes from either inline data (-d) or a file (-f).
// The format parameter controls how the file is interpreted: "json" or "csv".
func Resolve(data, file, format string) ([]byte, error) {
	if data != "" && file != "" {
		return nil, fmt.Errorf("specify either -d or -f, not both")
	}

	if data != "" {
		// Validate it's valid JSON
		if !json.Valid([]byte(data)) {
			return nil, fmt.Errorf("invalid JSON in -d flag")
		}
		return []byte(data), nil
	}

	if file != "" {
		raw, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("read file: %w", err)
		}

		format = strings.ToLower(format)
		if format == "csv" {
			return CSVToJSON(raw)
		}

		// Default: treat as JSON
		if !json.Valid(raw) {
			return nil, fmt.Errorf("file %s does not contain valid JSON", file)
		}
		return raw, nil
	}

	return nil, fmt.Errorf("provide input with -d (inline JSON) or -f (file path)")
}
