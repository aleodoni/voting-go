import { useMutation, useQueryClient } from '@tanstack/react-query';
import { getApi } from '@voting/shared';
import toast from 'react-hot-toast';
import { LAST_SYNCHRONIZATIONS_QUERY_KEY } from './useLastSyncs';

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

		onSuccess: (data) => {
			toast.success(data.message || 'Sincronia executada com sucesso!');
			// Atualiza a lista de sincronias após executar
			queryClient.invalidateQueries({
				queryKey: [LAST_SYNCHRONIZATIONS_QUERY_KEY],
			});
		},
		onError: () => {
			toast.error('Erro ao executar a sincronia');
		},
	});
}
