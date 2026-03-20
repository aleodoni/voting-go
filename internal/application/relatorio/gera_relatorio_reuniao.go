package relatorio

import (
	"context"

	domainRelatorio "github.com/aleodoni/voting-go/internal/domain/relatorio"
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
)

// GeraRelatorioReuniaoInput contém os dados necessários para gerar o relatório de uma reunião.
type GeraRelatorioReuniaoInput struct {
	ReuniaoID string
}

// GeraRelatorioReuniaoUseCase gera o relatório PDF de uma reunião com seus projetos e votações.
type GeraRelatorioReuniaoUseCase struct {
	repoReuniao domainVotacao.ReuniaoRepository
	generator   domainRelatorio.Generator
}

// NewGeraRelatorioReuniaoUseCase cria uma nova instância de [GeraRelatorioReuniaoUseCase].
func NewGeraRelatorioReuniaoUseCase(
	repoReuniao domainVotacao.ReuniaoRepository,
	generator domainRelatorio.Generator,
) *GeraRelatorioReuniaoUseCase {
	return &GeraRelatorioReuniaoUseCase{
		repoReuniao: repoReuniao,
		generator:   generator,
	}
}

// Execute gera o relatório da reunião informada em [GeraRelatorioReuniaoInput.ReuniaoID].
func (uc *GeraRelatorioReuniaoUseCase) Execute(ctx context.Context, input GeraRelatorioReuniaoInput) ([]byte, error) {
	reuniao, err := uc.repoReuniao.FindReuniaoByID(ctx, input.ReuniaoID)
	if err != nil {
		return nil, err
	}

	projetos, err := uc.repoReuniao.GetProjetosCompleto(ctx, input.ReuniaoID)
	if err != nil {
		return nil, err
	}

	output := domainRelatorio.ReuniaoOutput{
		ConDesc:        reuniao.ConDesc,
		RecNumero:      reuniao.RecNumero,
		RecTipoReuniao: reuniao.RecTipoReuniao,
		RecData:        reuniao.RecData.Format("02/01/2006"),
	}

	for _, projeto := range projetos {
		item := domainRelatorio.ProjetoItem{
			CodigoProposicao: projeto.CodigoProposicao,
			Iniciativa:       projeto.Iniciativa,
			Relator:          projeto.Relator,
		}

		if projeto.Votacao != nil {
			item.Votacao = buildVotacaoItem(projeto.Votacao)
		}

		output.Projetos = append(output.Projetos, item)
	}

	return uc.generator.Gera(ctx, output)
}

func buildVotacaoItem(v *domainVotacao.Votacao) *domainRelatorio.VotacaoItem {
	totais := map[string]int{
		"F": 0,
		"C": 0,
		"R": 0,
		"V": 0,
		"A": 0,
	}
	var votos []domainRelatorio.VotoItem

	if v.Votos != nil {
		for _, voto := range *v.Votos {
			totais[string(voto.Voto)]++

			item := domainRelatorio.VotoItem{
				UsuarioNome: voto.Usuario.Nome,
				Opcao:       string(voto.Voto),
			}

			if voto.Restricao != nil {
				item.Restricao = voto.Restricao.Restricao
			}

			if voto.VotoContrario != nil && voto.VotoContrario.Parecer != nil {
				item.VotoContrario = voto.VotoContrario.Parecer.Vereador
			}

			votos = append(votos, item)
		}
	}

	return &domainRelatorio.VotacaoItem{
		Totais: totais,
		Votos:  votos,
	}
}
