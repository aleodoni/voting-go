// Package tui contains the root model and screen navigation for the voting CLI.
package tui

import (
	"fmt"
	"strings"

	"github.com/aleodoni/voting-go/cmd/cli/tui/styles"
	"github.com/aleodoni/voting-go/cmd/cli/tui/views"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Screen representa as telas disponíveis na aplicação.
type Screen int

const (
	ScreenMenu Screen = iota
	ScreenReunioes
	ScreenUsuarios
)

// RootModel é o modelo raiz que gerencia a navegação entre telas.
type RootModel struct {
	screen   Screen
	menu     views.MenuModel
	reunioes views.ReunioesDiaModel
	usuarios views.UsuariosModel
	width    int
	height   int
	appName  string
	version  string
	env      string
	username string
}

// NewRootModel cria uma nova instância de [RootModel] iniciando na tela de menu.
func NewRootModel(appName, version, env, username string) RootModel {
	return RootModel{
		screen:   ScreenMenu,
		menu:     views.NewMenuModel(),
		appName:  appName,
		version:  version,
		env:      env,
		username: username,
	}
}

func (m RootModel) Init() tea.Cmd { return nil }

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if m.screen == ScreenMenu {
				return m, tea.Quit
			}
		case "esc":
			m.screen = ScreenMenu
			return m, nil
		case "enter":
			if m.screen == ScreenMenu {
				return m.handleMenuSelect()
			}
		}
	}

	var cmd tea.Cmd
	switch m.screen {
	case ScreenMenu:
		m.menu, cmd = m.menu.Update(msg)
	case ScreenReunioes:
		m.reunioes, cmd = m.reunioes.Update(msg)
	case ScreenUsuarios:
		m.usuarios, cmd = m.usuarios.Update(msg)
	}
	return m, cmd
}

// screenTitle retorna o título da tela ativa.
func (m RootModel) screenTitle() string {
	switch m.screen {
	case ScreenReunioes:
		return "📋 Reuniões do dia"
	case ScreenUsuarios:
		return "👥 Usuários"
	default:
		return fmt.Sprintf("⚡ %s", m.appName)
	}
}

// helpKeys retorna os atalhos da tela ativa.
func (m RootModel) helpKeys() string {
	switch m.screen {
	case ScreenReunioes:
		return "↑/↓ navegar   enter ver projetos   r recarregar   esc voltar   ctrl+c sair"
	case ScreenUsuarios:
		return "↑/↓ navegar   enter editar   esc voltar   ctrl+c sair"
	default:
		return "↑/↓ navegar   enter selecionar   q sair"
	}
}

// statusBar monta o rodapé com atalhos à esquerda e status à direita.
func (m RootModel) statusBar(width int) string {
	envColor := styles.ColorSuccess
	envIcon := "🟢"
	if m.env == "staging" {
		envColor = styles.ColorWarning
		envIcon = "🟡"
	} else if m.env == "development" {
		envColor = styles.ColorMuted
		envIcon = "⚪"
	}

	help := styles.Muted.Render(m.helpKeys())
	user := styles.Muted.Render("👤 " + m.username)
	env := lipgloss.NewStyle().Foreground(envColor).Render(envIcon + " " + m.env)
	version := styles.Muted.Render("v" + m.version)

	right := user + "  " + env + "  " + version
	gap := width - lipgloss.Width(help) - lipgloss.Width(right)
	if gap < 1 {
		gap = 1
	}

	return help + strings.Repeat(" ", gap) + right
}

func (m RootModel) View() string {
	if m.width == 0 {
		return ""
	}

	frameWidth := m.width - 2
	innerWidth := frameWidth - 6 // borda(2) + padding(4)
	innerHeight := m.height - 6

	var content string
	switch m.screen {
	case ScreenReunioes:
		content = m.reunioes.View()
	case ScreenUsuarios:
		content = m.usuarios.View()
	default:
		content = m.menu.View()
	}

	separator := lipgloss.NewStyle().
		Foreground(styles.ColorBorder).
		Render(strings.Repeat("─", innerWidth))

	contentHeight := innerHeight - 2
	if contentHeight < 1 {
		contentHeight = 1
	}

	contentBox := lipgloss.NewStyle().
		Height(contentHeight).
		Width(innerWidth).
		Render(content)

	body := lipgloss.JoinVertical(lipgloss.Left,
		contentBox,
		separator,
		m.statusBar(innerWidth),
	)

	// Moldura sem título na borda — título fica em linha separada acima
	frame := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.ColorPrimary).
		Padding(1, 2).
		Width(frameWidth)

	// Título centralizado acima da moldura
	title := lipgloss.NewStyle().
		Foreground(styles.ColorPrimary).
		Bold(true).
		Width(frameWidth).
		Align(lipgloss.Center).
		Render("─── " + m.screenTitle() + " ───")

	return lipgloss.JoinVertical(lipgloss.Left,
		title,
		frame.Render(body),
	)
}

// handleMenuSelect navega para a tela correspondente e dispara o Init da view.
func (m RootModel) handleMenuSelect() (tea.Model, tea.Cmd) {
	switch m.menu.SelectedItem().Label {
	case "Reuniões do dia":
		m.screen = ScreenReunioes
		m.reunioes = views.NewReunioesDiaModel()
		return m, m.reunioes.Init()
	case "Gerenciar usuários":
		m.screen = ScreenUsuarios
		m.usuarios = views.NewUsuariosModel()
		return m, m.usuarios.Init()
	case "Sair":
		return m, tea.Quit
	}
	return m, nil
}
