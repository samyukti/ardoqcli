package input

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

// CSVToJSON converts CSV bytes to JSON.
// If the CSV has a single data row, it returns a JSON object.
// If multiple data rows, it returns a JSON array of objects.
func CSVToJSON(data []byte) ([]byte, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parse CSV: %w", err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV must have a header row and at least one data row")
	}

	headers := records[0]
	for i, h := range headers {
		headers[i] = strings.TrimSpace(h)
	}

	var rows []map[string]any
	for _, record := range records[1:] {
		row := make(map[string]any, len(headers))
		for i, h := range headers {
			if i < len(record) {
				row[h] = record[i]
			}
		}
		rows = append(rows, row)
	}

	// Single row: return object, not array
	if len(rows) == 1 {
		return json.Marshal(rows[0])
	}

	return json.Marshal(rows)
}
