import { useQuery } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

export const LAST_SYNCHRONIZATIONS_QUERY_KEY = 'ultimas-sincronias';

export type SynchronizationDTO = {
	id: string;
	iniciado_em: string;
	finalizado_em: string;
	sucesso: boolean;
	reunioes_sincronizadas: number;
	projetos_sincronizados: number;
	pareceres_sincronizados: number;
};

type SynchronizationsResponse = {
	sincronias: SynchronizationDTO[];
};

async function fetchSynchronizations(): Promise<SynchronizationDTO[]> {
	const { data } = await getApi().get<SynchronizationsResponse>('/sincronia');

	return data.sincronias;
}

export function useSynchronizations() {
	return useQuery({
		queryKey: [LAST_SYNCHRONIZATIONS_QUERY_KEY],
		queryFn: fetchSynchronizations,
		refetchOnReconnect: true,
		staleTime: 0,
	});
}
