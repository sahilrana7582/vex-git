package ui

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
)

var (
	Banner  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#fc036f"))
	Success = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("42"))
	Error   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("196"))
)

func Interactive() bool {
	return isatty.IsTerminal(os.Stdin.Fd()) && isatty.IsTerminal(os.Stdout.Fd())
}
