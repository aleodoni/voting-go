package persistence

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

// ── Structs intermediárias (JSON bruto da função SQL) ─────────────

type ProjetoCompletoJSON struct {
	Projeto   json.RawMessage `json:"projeto"`
	Pareceres json.RawMessage `json:"pareceres"`
	Votacao   json.RawMessage `json:"votacao"`
}

type projetoJSON struct {
	ID                string `json:"id"`
	Sumula            string `json:"sumula"`
	Relator           string `json:"relator"`
	TemEmendas        bool   `json:"tem_emendas"`
	PacID             int    `json:"pac_id"`
	ParID             int    `json:"par_id"`
	CodigoProposicao  string `json:"codigo_proposicao"`
	Iniciativa        string `json:"iniciativa"`
	ConclusaoComissao string `json:"conclusao_comissao"`
	ConclusaoRelator  string `json:"conclusao_relator"`
	ReuniaoID         string `json:"reuniao_id"`
	CreatedAt         pgTime `json:"created_at"`
	UpdatedAt         pgTime `json:"updated_at"`
}

type parecerJSONRaw struct {
	ID               string `json:"id"`
	CodigoProposicao string `json:"codigo_proposicao"`
	TCPNome          string `json:"tcp_nome"`
	Vereador         string `json:"vereador"`
	IDTexto          int    `json:"id_texto"`
	ProjetoID        string `json:"projeto_id"`
	CreatedAt        pgTime `json:"created_at"`
	UpdatedAt        pgTime `json:"updated_at"`
}

type usuarioJSON struct {
	ID           string  `json:"id"`
	KeycloakID   string  `json:"keycloak_id"`
	Email        string  `json:"email"`
	Nome         string  `json:"nome"`
	NomeFantasia *string `json:"nome_fantasia"`
	Username     string  `json:"username"`
	CreatedAt    pgTime  `json:"created_at"`
	UpdatedAt    pgTime  `json:"updated_at"`
}

type restricaoJSON struct {
	ID        string `json:"id"`
	Restricao string `json:"restricao"`
	VotoID    string `json:"voto_id"`
	CreatedAt pgTime `json:"created_at"`
	UpdatedAt pgTime `json:"updated_at"`
}

type votoContrarioJSON struct {
	ID        string          `json:"id"`
	IDTexto   int             `json:"id_texto"`
	ParecerID string          `json:"parecer_id"`
	VotoID    string          `json:"voto_id"`
	Parecer   *parecerJSONRaw `json:"parecer"`
	CreatedAt pgTime          `json:"created_at"`
	UpdatedAt pgTime          `json:"updated_at"`
}

type votoJSON struct {
	ID            string              `json:"id"`
	Voto          string              `json:"voto"`
	VotacaoID     string              `json:"votacao_id"`
	UsuarioID     string              `json:"usuario_id"`
	CreatedAt     pgTime              `json:"created_at"`
	UpdatedAt     pgTime              `json:"updated_at"`
	Usuario       *usuarioJSON        `json:"usuario"`
	Restricoes    []restricaoJSON     `json:"restricoes"`
	VotoContrario []votoContrarioJSON `json:"votoContrario"`
}

type votacaoJSON struct {
	ID        string     `json:"id"`
	ProjetoID *string    `json:"projeto_id"`
	Status    string     `json:"status"`
	Votos     []votoJSON `json:"votos"`
}

// ── Mapper público ────────────────────────────────────────────────

func ToDomainProjetoFromJSON(raw ProjetoCompletoJSON) (*votacao.Projeto, error) {
	var pj projetoJSON
	if err := json.Unmarshal(raw.Projeto, &pj); err != nil {
		return nil, fmt.Errorf("unmarshal projeto: %w", err)
	}

	var pareceres []parecerJSONRaw
	if err := json.Unmarshal(raw.Pareceres, &pareceres); err != nil {
		return nil, fmt.Errorf("unmarshal pareceres: %w", err)
	}

	var votacaoPtr *votacao.Votacao
	if len(raw.Votacao) > 0 && string(raw.Votacao) != "null" {
		var votacaoRaw votacaoJSON
		if err := json.Unmarshal(raw.Votacao, &votacaoRaw); err != nil {
			return nil, fmt.Errorf("unmarshal votacao: %w", err)
		}
		votacaoPtr = mapVotacao(votacaoRaw)
	}

	return &votacao.Projeto{
		ID:                pj.ID,
		Sumula:            pj.Sumula,
		Relator:           pj.Relator,
		TemEmendas:        pj.TemEmendas,
		PacID:             pj.PacID,
		ParID:             pj.ParID,
		CodigoProposicao:  pj.CodigoProposicao,
		Iniciativa:        pj.Iniciativa,
		ConclusaoComissao: pj.ConclusaoComissao,
		ConclusaoRelator:  pj.ConclusaoRelator,
		ReuniaoID:         pj.ReuniaoID,
		CreatedAt:         pj.CreatedAt.Time,
		UpdatedAt:         pj.UpdatedAt.Time,
		Pareceres:         mapParecerSlice(pareceres),
		Votacao:           votacaoPtr,
	}, nil
}

// ── Mappers internos ──────────────────────────────────────────────

func mapParecerSlice(js []parecerJSONRaw) *[]votacao.Parecer {
	if len(js) == 0 {
		return nil
	}

	out := make([]votacao.Parecer, len(js))
	for i, j := range js {
		out[i] = votacao.Parecer{
			ID:               j.ID,
			CodigoProposicao: j.CodigoProposicao,
			TCPNome:          j.TCPNome,
			Vereador:         j.Vereador,
			IDTexto:          j.IDTexto,
			ProjetoID:        j.ProjetoID,
			CreatedAt:        j.CreatedAt.Time,
			UpdatedAt:        j.UpdatedAt.Time,
		}
	}

	return &out
}

func mapVotacao(j votacaoJSON) *votacao.Votacao {
	if j.ID == "" {
		return nil
	}

	return &votacao.Votacao{
		AggregateRoot: domain.NewAggregateRoot(j.ID),
		ProjetoID:     j.ProjetoID,
		Status:        votacao.StatusVotacao(j.Status),
		Votos:         mapVotoSlice(j.Votos),
	}
}

func mapVotoSlice(js []votoJSON) *[]votacao.Voto {
	if len(js) == 0 {
		return nil
	}

	out := make([]votacao.Voto, len(js))
	for i, j := range js {
		out[i] = votacao.Voto{
			ID:            j.ID,
			Voto:          votacao.OpcaoVoto(j.Voto),
			VotacaoID:     j.VotacaoID,
			UsuarioID:     j.UsuarioID,
			CreatedAt:     j.CreatedAt.Time,
			UpdatedAt:     j.UpdatedAt.Time,
			Restricao:     mapRestricao(j.Restricoes),
			VotoContrario: mapVotoContrario(j.VotoContrario),
		}
	}

	return &out
}

func mapRestricao(js []restricaoJSON) *votacao.Restricao {
	if len(js) == 0 {
		return nil
	}

	j := js[0]

	return &votacao.Restricao{
		ID:        j.ID,
		Restricao: j.Restricao,
		VotoID:    j.VotoID,
		CreatedAt: j.CreatedAt.Time,
		UpdatedAt: j.UpdatedAt.Time,
	}
}

func mapVotoContrario(js []votoContrarioJSON) *votacao.VotoContrario {
	if len(js) == 0 {
		return nil
	}

	j := js[0]

	var parecer *votacao.Parecer
	if j.Parecer != nil {
		parecer = &votacao.Parecer{
			ID:               j.Parecer.ID,
			CodigoProposicao: j.Parecer.CodigoProposicao,
			TCPNome:          j.Parecer.TCPNome,
			Vereador:         j.Parecer.Vereador,
			IDTexto:          j.Parecer.IDTexto,
			ProjetoID:        j.Parecer.ProjetoID,
			CreatedAt:        j.Parecer.CreatedAt.Time,
			UpdatedAt:        j.Parecer.UpdatedAt.Time,
		}
	}

	return &votacao.VotoContrario{
		ID:        j.ID,
		IDTexto:   j.IDTexto,
		ParecerID: j.ParecerID,
		VotoID:    j.VotoID,
		Parecer:   parecer,
	}
}

// ── Time parser ───────────────────────────────────────────────────

type pgTime struct {
	time.Time
}

func (t *pgTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	formats := []string{
		"2006-01-02T15:04:05.999999",
		"2006-01-02T15:04:05",
		time.RFC3339Nano,
		time.RFC3339,
	}

	for _, f := range formats {
		if parsed, err := time.Parse(f, s); err == nil {
			t.Time = parsed
			return nil
		}
	}

	return fmt.Errorf("cannot parse time: %s", s)
}
