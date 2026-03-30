import { useQuery } from '@tanstack/react-query';
import { getApi } from '@voting/shared';
import { ProjectDTO } from '@/types/meeting';

async function fetchOpenVoting(): Promise<ProjectDTO[]> {
	const { data } = await getApi().get<ProjectDTO[]>('/votacao/aberta');
	return data;
}
export function useIsProjectVoting() {
	return useQuery({
		queryKey: ['open-voting'],
		queryFn: () => fetchOpenVoting(),
		retry: false,
	});
}
