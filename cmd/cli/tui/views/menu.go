// Package views contains the individual screen models for the voting CLI.
package views

import (
	"fmt"

	"github.com/aleodoni/voting-go/cmd/cli/tui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

// MenuItem representa uma opção do menu principal.
type MenuItem struct {
	Label       string
	Description string
	Icon        string
}

// MenuModel é o modelo da tela principal de navegação.
type MenuModel struct {
	items  []MenuItem
	cursor int
	width  int
	height int
}

// NewMenuModel cria uma nova instância de [MenuModel] com as opções padrão.
func NewMenuModel() MenuModel {
	return MenuModel{
		items: []MenuItem{
			{Icon: "📋", Label: "Reuniões do dia", Description: "Listar reuniões agendadas para hoje"},
			{Icon: "📮", Label: "Abrir votação", Description: "Abrir votação para um projeto"}, // era 🗳️
			{Icon: "🔒", Label: "Fechar votação", Description: "Fechar a votação em andamento"},
			{Icon: "🚫", Label: "Cancelar votação", Description: "Cancelar uma votação fechada"}, // era ❌
			{Icon: "👥", Label: "Gerenciar usuários", Description: "Listar e editar permissões de usuários"},
			{Icon: "🚪", Label: "Sair", Description: "Encerrar o programa"},
		},
	}
}

// SelectedItem retorna o item atualmente selecionado no menu.
func (m MenuModel) SelectedItem() MenuItem {
	return m.items[m.cursor]
}

func (m MenuModel) Init() tea.Cmd { return nil }

func (m MenuModel) Update(msg tea.Msg) (MenuModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	// Sem help aqui — fica no rodapé fixo do RootModel
	header := styles.Title.Render("⚡ Voting CLI") + "\n" +
		styles.Subtitle.Render("Painel de administração de votações") + "\n"

	menu := ""
	for i, item := range m.items {
		icon := item.Icon + " "
		if i == m.cursor {
			label := styles.Selected.Render(fmt.Sprintf(" ▶ %s%s ", icon, item.Label))
			desc := styles.Muted.Render("   " + item.Description)
			menu += label + "\n" + desc + "\n\n"
		} else {
			label := styles.Normal.Render(fmt.Sprintf("   %s%s", icon, item.Label))
			menu += label + "\n\n"
		}
	}

	return header + "\n" + menu
}
