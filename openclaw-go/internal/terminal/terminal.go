package terminal

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Theme colors for terminal output.
var (
	colorSuccess = lipgloss.NewStyle().Foreground(lipgloss.Color("#00C853"))
	colorError   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF1744"))
	colorWarning = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFAB00"))
	colorInfo    = lipgloss.NewStyle().Foreground(lipgloss.Color("#2979FF"))
	colorMuted   = lipgloss.NewStyle().Foreground(lipgloss.Color("#757575"))
	colorBold    = lipgloss.NewStyle().Bold(true)
	colorHeader  = lipgloss.NewStyle().Bold(true).Underline(true)
)

// Success formats text in success color.
func Success(s string) string { return colorSuccess.Render(s) }

// Error formats text in error color.
func Error(s string) string { return colorError.Render(s) }

// Warning formats text in warning color.
func Warning(s string) string { return colorWarning.Render(s) }

// Info formats text in info color.
func Info(s string) string { return colorInfo.Render(s) }

// Muted formats text in muted color.
func Muted(s string) string { return colorMuted.Render(s) }

// Bold formats text in bold.
func Bold(s string) string { return colorBold.Render(s) }

// Header formats text as a header.
func Header(s string) string { return colorHeader.Render(s) }

// TableColumn defines a table column.
type TableColumn struct {
	Header string
	Width  int
}

// RenderTable renders a simple ASCII table.
func RenderTable(columns []TableColumn, rows [][]string) string {
	var sb strings.Builder

	// Header
	for i, col := range columns {
		if i > 0 {
			sb.WriteString("  ")
		}
		h := col.Header
		if col.Width > 0 && len(h) < col.Width {
			h = h + strings.Repeat(" ", col.Width-len(h))
		}
		sb.WriteString(colorHeader.Render(h))
	}
	sb.WriteString("\n")

	// Separator
	for i, col := range columns {
		if i > 0 {
			sb.WriteString("  ")
		}
		w := col.Width
		if w <= 0 {
			w = len(col.Header)
		}
		sb.WriteString(strings.Repeat("-", w))
	}
	sb.WriteString("\n")

	// Rows
	for _, row := range rows {
		for i, cell := range row {
			if i >= len(columns) {
				break
			}
			if i > 0 {
				sb.WriteString("  ")
			}
			col := columns[i]
			if col.Width > 0 && len(cell) < col.Width {
				cell = cell + strings.Repeat(" ", col.Width-len(cell))
			} else if col.Width > 0 && len(cell) > col.Width {
				cell = cell[:col.Width-1] + "…"
			}
			sb.WriteString(cell)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// PrintTable prints a table to stdout.
func PrintTable(columns []TableColumn, rows [][]string) {
	fmt.Print(RenderTable(columns, rows))
}

// IsTerminal returns true if stdout is a terminal.
func IsTerminal() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

// NoColor returns true if color output is disabled.
func NoColor() bool {
	return os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb"
}
