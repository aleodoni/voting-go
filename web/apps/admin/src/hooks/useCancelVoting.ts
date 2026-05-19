import { useMutation, useQueryClient } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

async function cancelarVotacao(projetoId: string): Promise<void> {
	await getApi().delete(`/projetos/${projetoId}/votacao`);
}

export function useCancelVoting(reuniaoId?: string) {
	const queryClient = useQueryClient();

	return useMutation<void, Error, string>({
		mutationFn: cancelarVotacao,
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
