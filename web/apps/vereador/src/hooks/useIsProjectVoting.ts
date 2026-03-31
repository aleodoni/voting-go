import { useQuery } from '@tanstack/react-query';
import { getApi, ProjetoDTO } from '@voting/shared';

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
