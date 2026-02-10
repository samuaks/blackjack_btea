package main

import (
	lipgloss "github.com/charmbracelet/lipgloss"
)

var (
	styles = lipgloss.NewStyle().
		Bold(true).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#1a1313")).
		Align(lipgloss.Left)
)
