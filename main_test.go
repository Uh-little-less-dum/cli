package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// Needed to permit charm tests in CI/CD.
func init() {
	lipgloss.SetColorProfile(termenv.Ascii)
}
