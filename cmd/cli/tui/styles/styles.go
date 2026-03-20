// Package styles defines the visual style tokens for the voting CLI.
package styles

import "github.com/charmbracelet/lipgloss"

var (
	ColorPrimary = lipgloss.Color("#5A8CF7")
	ColorSuccess = lipgloss.Color("#3DD68C")
	ColorWarning = lipgloss.Color("#F7C948")
	ColorDanger  = lipgloss.Color("#F76E6E")
	ColorMuted   = lipgloss.Color("#6B7280")
	ColorBorder  = lipgloss.Color("#5A8CF7")
	ColorText    = lipgloss.Color("#E2E8F0")

	// Selected — fundo azul escuro com texto branco, estilo Clipper clássico
	Selected = lipgloss.NewStyle().
			Background(lipgloss.Color("#0000AA")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true)

	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorPrimary).
		MarginBottom(1)

	Subtitle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			MarginBottom(1)

	Normal = lipgloss.NewStyle().
		Foreground(ColorText)

	Muted = lipgloss.NewStyle().
		Foreground(ColorMuted)

	Success = lipgloss.NewStyle().
		Foreground(ColorSuccess)

	Warning = lipgloss.NewStyle().
		Foreground(ColorWarning)

	Danger = lipgloss.NewStyle().
		Foreground(ColorDanger)

	Badge = func(color lipgloss.Color) lipgloss.Style {
		return lipgloss.NewStyle().
			Foreground(color).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(color).
			Padding(0, 1)
	}

	Card = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorBorder).
		Padding(1, 2).
		MarginBottom(1)

	Help = lipgloss.NewStyle().
		Foreground(ColorMuted).
		MarginTop(1)
)
