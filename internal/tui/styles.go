package tui

import "github.com/charmbracelet/lipgloss"

// Catppuccin Mocha colors
var (
	MochaRosewater = lipgloss.Color("#f5e0dc")
	MochaFlamingo  = lipgloss.Color("#f2cdcd")
	MochaPink      = lipgloss.Color("#f5c2e7")
	MochaMauve     = lipgloss.Color("#cba6f7")
	MochaRed       = lipgloss.Color("#f38ba8")
	MochaMaroon    = lipgloss.Color("#eba0ac")
	MochaPeach     = lipgloss.Color("#fab387")
	MochaYellow    = lipgloss.Color("#f9e2af")
	MochaGreen     = lipgloss.Color("#a6e3a1")
	MochaTeal      = lipgloss.Color("#94e2d5")
	MochaSky       = lipgloss.Color("#89dceb")
	MochaSapphire  = lipgloss.Color("#74c7ec")
	MochaBlue      = lipgloss.Color("#89b4fa")
	MochaLavender  = lipgloss.Color("#b4befe")
	MochaText      = lipgloss.Color("#cdd6f4")
	MochaSubtext1  = lipgloss.Color("#bac2de")
	MochaSubtext0  = lipgloss.Color("#a6adc8")
	MochaOverlay2  = lipgloss.Color("#9399b2")
	MochaOverlay1  = lipgloss.Color("#7f849c")
	MochaOverlay0  = lipgloss.Color("#6c7086")
	MochaSurface2  = lipgloss.Color("#585b70")
	MochaSurface1  = lipgloss.Color("#45475a")
	MochaSurface0  = lipgloss.Color("#313244")
	MochaBase      = lipgloss.Color("#1e1e2e")
	MochaMantle    = lipgloss.Color("#181825")
	MochaCrust     = lipgloss.Color("#11111b")
)

var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(MochaMauve).
			Padding(0).
			MarginBottom(0)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(MochaMauve).
			Bold(true)

	NormalStyle = lipgloss.NewStyle().
			Foreground(MochaText)

	HelpStyle = lipgloss.NewStyle().
			Foreground(MochaOverlay1).
			Italic(true).
			MarginBottom(1)

	FooterStyle = lipgloss.NewStyle().
			Foreground(MochaOverlay0).
			MarginTop(1)

	StatusNavStyle = lipgloss.NewStyle().
			Foreground(MochaBase).
			Background(MochaBlue).
			Padding(0, 1).
			MarginRight(1).
			Bold(true)

	StatusInputStyle = lipgloss.NewStyle().
				Foreground(MochaBase).
				Background(MochaGreen).
				Padding(0, 1).
				MarginRight(1).
				Bold(true)

	InputPromptStyle = lipgloss.NewStyle().
				Foreground(MochaPeach).
				Bold(true)

	PlaceholderStyle = lipgloss.NewStyle().
				Foreground(MochaOverlay0)
)
