package cli_styles

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var UlldBlue = "#0ba5e9"

var DocStyle = lipgloss.NewStyle().Margin(0, 2)

// HelpStyle styling for help context menu
var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

// ErrStyle provides styling for error messages
var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render

// AlertStyle provides styling for alert messages
var AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render

var UlldBlueLipgloss = lipgloss.Color(UlldBlue)

var TitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFDF5")).
	Background(lipgloss.Color(UlldBlue)).
	Padding(0, 1)

var UlldBlueForeground = lipgloss.NewStyle().Foreground(UlldBlueLipgloss)

func GetHuhTheme() *huh.Theme {
	t := huh.ThemeCharm()
	t.Blurred.Title = UlldBlueForeground
	t.Focused.Title = UlldBlueForeground
	t.Focused.FocusedButton = t.Focused.FocusedButton.Background(UlldBlueLipgloss)
	t.Blurred.FocusedButton = t.Blurred.FocusedButton.Background(UlldBlueLipgloss)
	return t
}
