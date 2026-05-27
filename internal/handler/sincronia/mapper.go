package sincronia

import domainSincronia "github.com/aleodoni/voting-go/internal/domain/sincronia"

func ToSincroniaResponse(s *domainSincronia.Sincronia) *SincroniaResponse {
	resp := &SincroniaResponse{
		ID:                     s.ID,
		IniciadoEm:             s.IniciadoEm,
		FinalizadoEm:           s.FinalizadoEm,
		Sucesso:                s.Sucesso,
		MensagemErro:           s.MensagemErro,
		ReunioesSincronizadas:  s.ReunioesSincronizadas,
		ProjetosSincronizados:  s.ProjetosSincronizados,
		PareceresSincronizados: s.PareceresSincronizados,
	}
	return resp
}

func ToListUltimasSincroniasResponse(output *domainSincronia.ListSincronia) ListUltimasSincroniasResponse {
	sincronias := make([]*SincroniaResponse, len(output.Sincronias))

	for i, s := range output.Sincronias {
		sincronias[i] = ToSincroniaResponse(s)
	}

	return ListUltimasSincroniasResponse{
		Sincronias: sincronias,
	}
}
