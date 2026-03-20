package votacao

import "time"

// domainEvent é uma base embutida que implementa a interface [domain.DomainEvent].
type domainEvent struct {
	occurredAt time.Time
}

func (e domainEvent) OccurredAt() time.Time { return e.occurredAt }

func newDomainEvent() domainEvent {
	return domainEvent{occurredAt: time.Now()}
}

// VotacaoAbertaEvent é publicado quando uma votação é aberta.
type VotacaoAbertaEvent struct {
	domainEvent
	VotacaoID string
	ProjetoID string
}

func (VotacaoAbertaEvent) EventName() string { return "votacao_aberta" }

// VotacaoFechadaEvent é publicado quando uma votação é fechada.
type VotacaoFechadaEvent struct {
	domainEvent
	VotacaoID string
	ProjetoID string
}

func (VotacaoFechadaEvent) EventName() string { return "votacao_fechada" }

// VotacaoCanceladaEvent é publicado quando uma votação é cancelada.
type VotacaoCanceladaEvent struct {
	domainEvent
	VotacaoID string
	ProjetoID string
}

func (VotacaoCanceladaEvent) EventName() string { return "votacao_cancelada" }

// VotoRegistradoEvent é publicado quando um voto é registrado.
type VotoRegistradoEvent struct {
	domainEvent
	VotacaoID string
	UsuarioID string
}

func (VotoRegistradoEvent) EventName() string { return "voto_registrado" }
