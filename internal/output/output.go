package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	infoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
)

// JSON writes pretty-printed JSON to stdout or to a file if outFile is set.
func JSON(data []byte, outFile string) error {
	var buf bytes.Buffer
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		// Not valid JSON; write raw
		buf.Reset()
		buf.Write(data)
	}
	buf.WriteByte('\n')

	if outFile != "" {
		if err := os.WriteFile(outFile, buf.Bytes(), 0o644); err != nil {
			return fmt.Errorf("write output file: %w", err)
		}
		Info("Written to %s", outFile)
		return nil
	}

	fmt.Print(buf.String())
	return nil
}

// Error prints an error message to stderr.
func Error(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stderr, errorStyle.Render("Error: "+msg))
}

// Success prints a success message to stderr.
func Success(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stderr, successStyle.Render(msg))
}

// Info prints an info message to stderr.
func Info(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stderr, infoStyle.Render(msg))
}
