package votacao_test

import (
	"testing"
	"time"

	"github.com/aleodoni/go-ddd/domain"

	"github.com/aleodoni/voting-go/internal/domain/votacao"
)

func novaVotacao() *votacao.Votacao {
	return &votacao.Votacao{AggregateRoot: domain.NewAggregateRoot("votacao-1")}
}

// Todo evento levantado pela votação deve trazer o instante em que ocorreu.
func TestVotacao_EventosTrazemOccurredAt(t *testing.T) {
	casos := []struct {
		nome   string
		acao   func(v *votacao.Votacao)
		evento string
	}{
		{"Abrir", func(v *votacao.Votacao) { v.Abrir("projeto-1") }, "votacao_aberta"},
		{"Fechar", func(v *votacao.Votacao) { v.Fechar() }, "votacao_fechada"},
		{"Cancelar", func(v *votacao.Votacao) { v.Cancelar() }, "votacao_cancelada"},
		{"RegistrarVoto", func(v *votacao.Votacao) { v.RegistrarVoto(&votacao.Voto{UsuarioID: "usuario-1"}) }, "voto_registrado"},
	}

	for _, tc := range casos {
		t.Run(tc.nome, func(t *testing.T) {
			v := novaVotacao()

			antes := time.Now()
			tc.acao(v)
			depois := time.Now()

			eventos := v.PullEvents()
			if len(eventos) != 1 {
				t.Fatalf("esperava 1 evento, obteve %d", len(eventos))
			}

			if eventos[0].EventName() != tc.evento {
				t.Fatalf("esperava o evento %q, obteve %q", tc.evento, eventos[0].EventName())
			}

			ocorreuEm := eventos[0].OccurredAt()
			if ocorreuEm.IsZero() {
				t.Fatal("OccurredAt não deveria ser zero")
			}

			if ocorreuEm.Before(antes) || ocorreuEm.After(depois) {
				t.Fatalf("OccurredAt %v fora do intervalo [%v, %v]", ocorreuEm, antes, depois)
			}
		})
	}
}

func TestAbrir_AtualizaStatusEProjeto(t *testing.T) {
	v := novaVotacao()

	v.Abrir("projeto-1")

	if v.Status != votacao.StatusVotacaoA {
		t.Fatalf("esperava status %q, obteve %q", votacao.StatusVotacaoA, v.Status)
	}

	if v.ProjetoID == nil || *v.ProjetoID != "projeto-1" {
		t.Fatalf("esperava ProjetoID %q, obteve %v", "projeto-1", v.ProjetoID)
	}
}

func TestAbrir_EventoTrazOsIDs(t *testing.T) {
	v := novaVotacao()

	v.Abrir("projeto-1")

	eventos := v.PullEvents()
	if len(eventos) != 1 {
		t.Fatalf("esperava 1 evento, obteve %d", len(eventos))
	}

	evento, ok := eventos[0].(votacao.VotacaoAbertaEvent)
	if !ok {
		t.Fatalf("esperava VotacaoAbertaEvent, obteve %T", eventos[0])
	}

	if evento.VotacaoID != "votacao-1" || evento.ProjetoID != "projeto-1" {
		t.Fatalf("IDs inesperados no evento: %+v", evento)
	}
}
