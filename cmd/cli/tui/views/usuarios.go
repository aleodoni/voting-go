package views

import (
	"fmt"

	"github.com/aleodoni/voting-go/cmd/cli/tui/styles"
	"github.com/aleodoni/voting-go/internal/domain/usuario"
	tea "github.com/charmbracelet/bubbletea"
)

// UsuariosLoadedMsg é enviado ao barramento quando os usuários são carregados.
type UsuariosLoadedMsg struct {
	Usuarios []*usuario.Usuario
	Err      error
}

// UsuariosModel é o modelo da tela de gerenciamento de usuários.
type UsuariosModel struct {
	usuarios []*usuario.Usuario
	cursor   int
	loading  bool
	err      error
	loader   func() ([]*usuario.Usuario, error) // ← adicionar
}

// NewUsuariosModel cria uma nova instância de [UsuariosModel] em estado de carregamento.
func NewUsuariosModel(loader ...func() ([]*usuario.Usuario, error)) UsuariosModel {
	fn := defaultUsuarioLoader
	if len(loader) > 0 && loader[0] != nil {
		fn = loader[0]
	}
	return UsuariosModel{
		loading: true,
		loader:  fn,
	}
}

func (m UsuariosModel) Init() tea.Cmd {
	return func() tea.Msg {
		usuarios, err := m.loader()
		return UsuariosLoadedMsg{Usuarios: usuarios, Err: err}
	}
}

func (m UsuariosModel) Update(msg tea.Msg) (UsuariosModel, tea.Cmd) {
	switch msg := msg.(type) {
	case UsuariosLoadedMsg:
		m.loading = false
		m.err = msg.Err
		m.usuarios = msg.Usuarios
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.usuarios)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m UsuariosModel) View() string {
	header := styles.Title.Render("👥 Usuários") + "\n"

	if m.loading {
		return header + styles.Muted.Render("Carregando...") + "\n"
	}
	if m.err != nil {
		return header + styles.Danger.Render("Erro: "+m.err.Error()) + "\n"
	}
	if len(m.usuarios) == 0 {
		return header + styles.Muted.Render("Nenhum usuário encontrado.") + "\n"
	}

	list := ""
	for i, u := range m.usuarios {
		ativo := styles.Success.Render("● ativo  ")
		vota := styles.Muted.Render("não vota")
		admin := styles.Muted.Render("        ")

		if u.Credencial != nil {
			if !u.Credencial.IsActive() {
				ativo = styles.Danger.Render("● inativo")
			}
			if u.Credencial.CanVote() {
				vota = styles.Success.Render("vota    ")
			}
			if u.Credencial.IsAdmin() {
				admin = styles.Warning.Render("admin")
			}
		}

		row := fmt.Sprintf("%-30s  %s  %s  %s", u.Nome, ativo, vota, admin)
		if i == m.cursor {
			list += styles.Selected.Render("▶ "+row) + "\n"
		} else {
			list += styles.Normal.Render("  "+row) + "\n"
		}
	}

	help := styles.Help.Render("↑/↓  navegar   enter  editar permissões   esc  voltar")
	return header + "\n" + list + "\n" + help
}

// defaultUsuarioLoader é o loader padrão usado quando nenhum é fornecido.
// Retorna dados de exemplo para facilitar o desenvolvimento.
func defaultUsuarioLoader() ([]*usuario.Usuario, error) {
	return []*usuario.Usuario{
		{
			ID:         "user-1",
			Nome:       "Admin Teste",
			KeycloakID: "keycloak-admin",
			Credencial: &usuario.Credencial{
				Ativo:           true,
				PodeVotar:       false,
				PodeAdministrar: true,
			},
		},
		{
			ID:         "user-2",
			Nome:       "Vereador Teste",
			KeycloakID: "keycloak-vereador",
			Credencial: &usuario.Credencial{
				Ativo:           true,
				PodeVotar:       true,
				PodeAdministrar: false,
			},
		},
		{
			ID:         "user-3",
			Nome:       "Usuário Inativo",
			KeycloakID: "keycloak-inativo",
			Credencial: &usuario.Credencial{
				Ativo:           false,
				PodeVotar:       false,
				PodeAdministrar: false,
			},
		},
	}, nil
}
