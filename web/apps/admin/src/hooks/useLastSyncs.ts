import { useQuery } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

export const LAST_SYNCHRONIZATIONS_QUERY_KEY = 'ultimas-sincronias';

export type SynchronizationDTO = {
	id: string;
	iniciado_em: string;
	finalizado_em: string | null;
	sucesso: boolean | null;
	reunioes_sincronizadas: number;
	projetos_sincronizados: number;
	pareceres_sincronizados: number;
};

type SynchronizationsResponse = {
	sincronias: SynchronizationDTO[];
};

async function fetchLastSynchs(): Promise<SynchronizationDTO[]> {
	const { data } = await getApi().get<SynchronizationsResponse>('/sincronia');

	return data.sincronias;
}

export function useLastSynchs(refetchInterval: number | false = false) {
	return useQuery({
		queryKey: [LAST_SYNCHRONIZATIONS_QUERY_KEY],
		queryFn: fetchLastSynchs,
		refetchOnReconnect: true,
		staleTime: 0,
		refetchInterval,
	});
}
