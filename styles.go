package main

import (
	lipgloss "github.com/charmbracelet/lipgloss"
)

var (
	styles = lipgloss.NewStyle().
		Bold(true).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0000")).
		Align(lipgloss.Left)
)
