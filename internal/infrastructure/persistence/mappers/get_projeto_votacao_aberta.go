package mappers

import (
	"encoding/json"

	"github.com/aleodoni/go-ddd/domain"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

type projetoVotacaoAbertaJSON struct {
	ProjetoID         string        `json:"projetoId"`
	Sumula            string        `json:"sumula"`
	Relator           string        `json:"relator"`
	TemEmendas        bool          `json:"temEmendas"`
	PacID             int           `json:"pacId"`
	ParID             int           `json:"parId"`
	CodigoProposicao  string        `json:"codigoProposicao"`
	Iniciativa        string        `json:"iniciativa"`
	ConclusaoComissao string        `json:"conclusaoComissao"`
	ConclusaoRelator  string        `json:"conclusaoRelator"`
	CreatedAt         flexTime      `json:"createdAt"`
	UpdatedAt         flexTime      `json:"updatedAt"`
	ReuniaoID         string        `json:"reuniaoId"`
	VotacaoID         string        `json:"votacaoId"`
	VotacaoStatus     string        `json:"votacaoStatus"`
	VotacaoCreatedAt  flexTime      `json:"votacaoCreatedAt"`
	VotacaoUpdatedAt  flexTime      `json:"votacaoUpdatedAt"`
	Votos             []votoJSON    `json:"votos"`
	Pareceres         []parecerJSON `json:"pareceres"`
}

type votoJSON struct {
	ID        string      `json:"id"`
	Voto      string      `json:"voto"`
	VotacaoID string      `json:"votacaoId"`
	UsuarioID string      `json:"usuarioId"`
	CreatedAt flexTime    `json:"createdAt"`
	UpdatedAt flexTime    `json:"updatedAt"`
	Usuario   usuarioJSON `json:"usuario"`
}

type usuarioJSON struct {
	ID           string   `json:"id"`
	KeycloakID   string   `json:"keycloakId"`
	Email        string   `json:"email"`
	Nome         string   `json:"nome"`
	NomeFantasia *string  `json:"nomeFantasia"`
	Username     *string  `json:"username"`
	CreatedAt    flexTime `json:"createdAt"`
	UpdatedAt    flexTime `json:"updatedAt"`
}

type parecerJSON struct {
	ID               string   `json:"id"`
	CodigoProposicao string   `json:"codigoProposicao"`
	TCPNome          string   `json:"tcpNome"`
	Vereador         string   `json:"vereador"`
	IDTexto          int      `json:"idTexto"`
	ProjetoID        string   `json:"projetoId"`
	CreatedAt        flexTime `json:"createdAt"`
	UpdatedAt        flexTime `json:"updatedAt"`
}

// ToDomainProjetoVotacaoAberta converte o JSON retornado por GetProjectOpenVoting
// para a entidade de domínio [votacao.Projeto].
func ToDomainProjetoVotacaoAberta(data []byte) (*votacao.Projeto, error) {
	var row projetoVotacaoAbertaJSON
	if err := json.Unmarshal(data, &row); err != nil {
		return nil, err
	}

	votos := make([]votacao.Voto, 0, len(row.Votos))
	for _, vj := range row.Votos {
		u := domainUsuario.Usuario{
			AggregateRoot: domain.NewAggregateRoot(vj.Usuario.ID),
			KeycloakID:    vj.Usuario.KeycloakID,
			Email:         vj.Usuario.Email,
			Nome:          vj.Usuario.Nome,
			NomeFantasia:  vj.Usuario.NomeFantasia,
			CreatedAt:     vj.Usuario.CreatedAt.Time,
			UpdatedAt:     vj.Usuario.UpdatedAt.Time,
		}
		if vj.Usuario.Username != nil {
			u.Username = *vj.Usuario.Username
		}

		votos = append(votos, votacao.Voto{
			ID:        vj.ID,
			Voto:      votacao.OpcaoVoto(vj.Voto),
			VotacaoID: vj.VotacaoID,
			UsuarioID: vj.UsuarioID,
			CreatedAt: vj.CreatedAt.Time,
			UpdatedAt: vj.UpdatedAt.Time,
			Usuario:   u,
		})
	}

	pareceres := make([]votacao.Parecer, 0, len(row.Pareceres))
	for _, pj := range row.Pareceres {
		pareceres = append(pareceres, votacao.Parecer{
			ID:               pj.ID,
			CodigoProposicao: pj.CodigoProposicao,
			TCPNome:          pj.TCPNome,
			Vereador:         pj.Vereador,
			IDTexto:          pj.IDTexto,
			ProjetoID:        pj.ProjetoID,
			CreatedAt:        pj.CreatedAt.Time,
			UpdatedAt:        pj.UpdatedAt.Time,
		})
	}

	projetoID := row.ProjetoID
	v := &votacao.Votacao{
		AggregateRoot: domain.NewAggregateRoot(row.VotacaoID),
		ProjetoID:     &projetoID,
		Status:        votacao.StatusVotacao(row.VotacaoStatus),
		CreatedAt:     row.VotacaoCreatedAt.Time,
		UpdatedAt:     row.VotacaoUpdatedAt.Time,
		Votos:         &votos,
	}

	return &votacao.Projeto{
		ID:                row.ProjetoID,
		Sumula:            row.Sumula,
		Relator:           row.Relator,
		TemEmendas:        row.TemEmendas,
		PacID:             row.PacID,
		ParID:             row.ParID,
		CodigoProposicao:  row.CodigoProposicao,
		Iniciativa:        row.Iniciativa,
		ConclusaoComissao: row.ConclusaoComissao,
		ConclusaoRelator:  row.ConclusaoRelator,
		ReuniaoID:         row.ReuniaoID,
		CreatedAt:         row.CreatedAt.Time,
		UpdatedAt:         row.UpdatedAt.Time,
		Pareceres:         &pareceres,
		Votacao:           v,
	}, nil
}
