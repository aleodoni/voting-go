import { useQuery } from '@tanstack/react-query';
import { getApi } from '@voting/shared';
import { ProjectDTO } from '@/types/meeting';

async function fetchProjectsByMeeting(
	reuniaoId?: string,
): Promise<ProjectDTO[]> {
	const { data } = await getApi().get<ProjectDTO[]>(
		`/reunioes/${reuniaoId}/projetos`,
	);
	return data;
}

export function useProjectsMeeting(reuniaoId?: string) {
	return useQuery({
		queryKey: ['projects-meeting', reuniaoId],
		queryFn: () => fetchProjectsByMeeting(reuniaoId),
		enabled: !!reuniaoId,
		refetchOnReconnect: true,
		staleTime: 0,
	});
}
