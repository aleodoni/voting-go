package views

import (
	"fmt"
	"time"

	"github.com/aleodoni/voting-go/cmd/cli/tui/styles"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
	tea "github.com/charmbracelet/bubbletea"
)

// ReunioesDiaLoadedMsg é enviado ao barramento quando as reuniões são carregadas.
type ReunioesDiaLoadedMsg struct {
	Reunioes []*votacao.Reuniao
	Err      error
}

// ReunioesDiaModel é o modelo da tela de reuniões do dia.
type ReunioesDiaModel struct {
	reunioes []*votacao.Reuniao
	cursor   int
	loading  bool
	err      error
	loader   func() ([]*votacao.Reuniao, error)
}

// NewReunioesDiaModel cria uma nova instância de [ReunioesDiaModel].
//
// O parâmetro loader é a função que busca as reuniões — normalmente o Execute
// de um [RetornaReunioesDiaUseCase]. Quando nil, usa dados de exemplo.
func NewReunioesDiaModel(loader ...func() ([]*votacao.Reuniao, error)) ReunioesDiaModel {
	fn := defaultReuniaoLoader
	if len(loader) > 0 && loader[0] != nil {
		fn = loader[0]
	}
	return ReunioesDiaModel{
		loading: true,
		loader:  fn,
	}
}

// Init dispara o carregamento das reuniões assim que a tela é exibida.
func (m ReunioesDiaModel) Init() tea.Cmd {
	return m.carregarReunioes()
}

// carregarReunioes retorna um [tea.Cmd] que executa o loader em background
// e publica um [ReunioesDiaLoadedMsg] com o resultado.
func (m ReunioesDiaModel) carregarReunioes() tea.Cmd {
	return func() tea.Msg {
		reunioes, err := m.loader()
		return ReunioesDiaLoadedMsg{Reunioes: reunioes, Err: err}
	}
}

func (m ReunioesDiaModel) Update(msg tea.Msg) (ReunioesDiaModel, tea.Cmd) {
	switch msg := msg.(type) {
	case ReunioesDiaLoadedMsg:
		m.loading = false
		m.err = msg.Err
		m.reunioes = msg.Reunioes
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.reunioes)-1 {
				m.cursor++
			}
		case "r":
			// Recarrega a lista manualmente
			m.loading = true
			return m, m.carregarReunioes()
		}
	}
	return m, nil
}

func (m ReunioesDiaModel) View() string {
	header := styles.Title.Render("📋 Reuniões do dia") + "\n"

	if m.loading {
		return header + styles.Muted.Render("Carregando...") + "\n"
	}
	if m.err != nil {
		return header + styles.Danger.Render("Erro: "+m.err.Error()) + "\n"
	}
	if len(m.reunioes) == 0 {
		return header + styles.Muted.Render("Nenhuma reunião agendada para hoje.") + "\n"
	}

	list := ""
	for i, r := range m.reunioes {
		row := fmt.Sprintf("%s  %s", r.ID, r.RecTipoReuniao)
		if i == m.cursor {
			list += styles.Selected.Render("▶ "+row) + "\n"
		} else {
			list += styles.Normal.Render("  "+row) + "\n"
		}
	}

	help := styles.Help.Render("↑/↓  navegar   enter  ver projetos   r  recarregar   esc  voltar")
	return header + "\n" + list + "\n" + help
}

// defaultReuniaoLoader é o loader padrão usado quando nenhum é fornecido.
// Retorna dados de exemplo para facilitar o desenvolvimento.
func defaultReuniaoLoader() ([]*votacao.Reuniao, error) {
	hoje := time.Now()
	return []*votacao.Reuniao{
		{ID: "reuniao-1", RecTipoReuniao: "Reunião Ordinária", RecData: hoje},
		{ID: "reuniao-2", RecTipoReuniao: "Reunião Extraordinária", RecData: hoje},
	}, nil
}
