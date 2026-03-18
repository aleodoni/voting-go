package reuniao

import (
	domainVotacao "github.com/aleodoni/voting-go/internal/domain/votacao"
)

func toReuniaoResponse(r *domainVotacao.Reuniao) *ReuniaoResponse {
	resp := &ReuniaoResponse{
		ID:             r.ID,
		ConID:          r.ConID,
		RecID:          r.RecID,
		PacID:          r.PacID,
		ConDesc:        r.ConDesc,
		ConSigla:       r.ConSigla,
		RecTipoReuniao: r.RecTipoReuniao,
		RecNumero:      r.RecNumero,
		RecData:        r.RecData,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}

	if r.Projetos != nil {
		projetos := make([]ProjetoResponse, len(*r.Projetos))
		for i, p := range *r.Projetos {
			projetos[i] = toProjetoResponse(&p)
		}
		resp.Projetos = &projetos
	}

	return resp
}

func toProjetoResponse(p *domainVotacao.Projeto) ProjetoResponse {
	resp := ProjetoResponse{
		ID:                p.ID,
		Sumula:            p.Sumula,
		Relator:           p.Relator,
		TemEmendas:        p.TemEmendas,
		PacID:             p.PacID,
		ParID:             p.ParID,
		CodigoProposicao:  p.CodigoProposicao,
		Iniciativa:        p.Iniciativa,
		ConclusaoComissao: p.ConclusaoComissao,
		ConclusaoRelator:  p.ConclusaoRelator,
		ReuniaoID:         p.ReuniaoID,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}

	if p.Pareceres != nil {
		pareceres := make([]ParecerResponse, len(*p.Pareceres))
		for i, par := range *p.Pareceres {
			pareceres[i] = toParecerResponse(&par)
		}
		resp.Pareceres = &pareceres
	}

	if p.Votacao != nil {
		v := toVotacaoResponse(p.Votacao)
		resp.Votacao = &v
	}

	return resp
}

func toReunioesDiaResponse(reunioes []*domainVotacao.Reuniao) []*ReuniaoResponse {
	resp := make([]*ReuniaoResponse, len(reunioes))
	for i, r := range reunioes {
		resp[i] = toReuniaoResponse(r)
	}
	return resp
}

func toProjetosResponse(projetos []*domainVotacao.Projeto) []ProjetoResponse {
	resp := make([]ProjetoResponse, len(projetos))
	for i, p := range projetos {
		resp[i] = toProjetoResponse(p)
	}
	return resp
}

func toParecerResponse(p *domainVotacao.Parecer) ParecerResponse {
	return ParecerResponse{
		ID:               p.ID,
		CodigoProposicao: p.CodigoProposicao,
		TCPNome:          p.TCPNome,
		Vereador:         p.Vereador,
		IDTexto:          p.IDTexto,
		ProjetoID:        p.ProjetoID,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

func toVotacaoResponse(v *domainVotacao.Votacao) VotacaoResponse {
	return VotacaoResponse{
		ID:        v.ID,
		ProjetoID: v.ProjetoID,
		Status:    string(v.Status),
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
