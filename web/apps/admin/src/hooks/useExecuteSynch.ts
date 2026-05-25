import { useMutation, useQueryClient } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

type ExecuteSynchronizationResponse = {
	message?: string;
};

async function executeSynch(): Promise<ExecuteSynchronizationResponse> {
	const { data } =
		await getApi().post<ExecuteSynchronizationResponse>('/sincronia');

	return data;
}

export function useExecuteSynch() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationFn: executeSynch,

		onSuccess: () => {
			// Atualiza a lista de sincronias após executar
			queryClient.invalidateQueries({
				queryKey: ['ultimas-sincronias'],
			});
		},
	});
}
