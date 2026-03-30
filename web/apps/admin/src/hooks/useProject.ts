import { useQuery } from '@tanstack/react-query';
import { getApi } from '@voting/shared';
import { ProjectDTO } from '@/types/meeting';

async function fetchProject(projetoId: string): Promise<ProjectDTO> {
	const { data } = await getApi().get<ProjectDTO>(`/projetos/${projetoId}`);
	return data;
}

export function useProject(projetoId?: string) {
	return useQuery({
		queryKey: ['project', projetoId],
		enabled: Boolean(projetoId), // só executa se projetoId existir
		queryFn: () => {
			if (!projetoId) {
				// retorna um erro controlado, evita usar "!"
				return Promise.reject(new Error('projectId é obrigatório'));
			}
			return fetchProject(projetoId);
		},
		retry: false,
	});
}
