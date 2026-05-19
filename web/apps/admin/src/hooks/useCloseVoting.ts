import { useMutation, useQueryClient } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

async function fecharVotacao(projetoId: string): Promise<void> {
	await getApi().post(`/projetos/${projetoId}/votacao/fechar`);
}

export function useCloseVoting(reuniaoId?: string) {
	const queryClient = useQueryClient();

	return useMutation<void, Error, string>({
		mutationFn: fecharVotacao,
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
