package votacao

import (
	"time"

	"github.com/aleodoni/go-ddd/domain"
)

type StatusVotacao string

const (
	StatusVotacaoA StatusVotacao = "A"
	StatusVotacaoF StatusVotacao = "F"
	StatusVotacaoV StatusVotacao = "V"
	StatusVotacaoC StatusVotacao = "C"
)

type Votacao struct {
	domain.AggregateRoot[string]

	ProjetoID *string
	Status    StatusVotacao
	CreatedAt time.Time
	UpdatedAt time.Time

	Projeto *Projeto
	Votos   *[]Voto
}

// Abrir abre a votação e levanta o evento [VotacaoAbertaEvent].
func (v *Votacao) Abrir(projetoID string) {
	v.Status = StatusVotacaoA
	v.ProjetoID = &projetoID
	v.RaiseEvent(VotacaoAbertaEvent{
		VotacaoID: v.ID,
		ProjetoID: projetoID,
	})
}

// Fechar fecha a votação e levanta o evento [VotacaoFechadaEvent].
func (v *Votacao) Fechar() {
	v.Status = StatusVotacaoV

	projetoID := ""
	if v.ProjetoID != nil {
		projetoID = *v.ProjetoID
	}

	v.RaiseEvent(VotacaoFechadaEvent{
		domainEvent: newDomainEvent(),
		VotacaoID:   v.ID,
		ProjetoID:   projetoID,
	})
}

// Cancelar cancela a votação e levanta o evento [VotacaoCanceladaEvent].
func (v *Votacao) Cancelar() {
	v.Status = StatusVotacaoC

	projetoID := ""
	if v.ProjetoID != nil {
		projetoID = *v.ProjetoID
	}

	v.RaiseEvent(VotacaoCanceladaEvent{
		domainEvent: newDomainEvent(),
		VotacaoID:   v.ID,
		ProjetoID:   projetoID,
	})
}

// RegistrarVoto registra um voto e levanta o evento [VotoRegistradoEvent].
func (v *Votacao) RegistrarVoto(voto *Voto) {
	v.RaiseEvent(VotoRegistradoEvent{
		domainEvent: newDomainEvent(),
		VotacaoID:   v.ID,
		UsuarioID:   voto.UsuarioID,
	})
}
