import { useMutation, useQueryClient } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

async function abrirVotacao(projetoId: string): Promise<void> {
	await getApi().post(`/projetos/${projetoId}/votacao/abrir`);
}

export function useOpenVoting(reuniaoId?: string) {
	const queryClient = useQueryClient();

	return useMutation<void, Error, string>({
		mutationFn: abrirVotacao,
		onSuccess: () => {
			queryClient.invalidateQueries({
				queryKey: ['projects-meeting', reuniaoId],
			});

			queryClient.invalidateQueries({
				queryKey: ['open-voting'],
			});
		},
	});
}
