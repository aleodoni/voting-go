// Command voting-cli provides a terminal UI for managing voting sessions.
package main

import (
	"fmt"
	"os"

	"github.com/aleodoni/voting-go/cmd/cli/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		tui.NewRootModel("Voting CLI", "1.0.0", "development", "admin"),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "erro ao iniciar o CLI: %v\n", err)
		os.Exit(1)
	}
}
