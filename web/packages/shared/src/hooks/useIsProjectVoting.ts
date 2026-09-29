import { useQuery } from '@tanstack/react-query';
import { getApi } from '../api-client';
import type { ProjetoDTO } from '../types';

async function fetchOpenVoting(): Promise<ProjetoDTO> {
	const { data } = await getApi().get<ProjetoDTO>('/votacao/aberta');
	return data;
}

export function useIsProjectVoting() {
	return useQuery({
		queryKey: ['open-voting'],
		queryFn: () => fetchOpenVoting(),
		retry: false,
	});
}
