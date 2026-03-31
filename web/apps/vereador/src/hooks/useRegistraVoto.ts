import { useMutation, useQueryClient } from '@tanstack/react-query';
import { getApi } from '@voting/shared';
import { RegistraVotoRequest } from '@/types/voto';

type RegistraVotoParams = {
	votacaoId: string;
	body: RegistraVotoRequest;
};

async function registraVoto({
	votacaoId,
	body,
}: RegistraVotoParams): Promise<void> {
	await getApi().post(`/votacao/${votacaoId}/voto`, body);
}

export function useRegistraVoto(reuniaoId?: string) {
	const queryClient = useQueryClient();

	return useMutation<void, Error, RegistraVotoParams>({
		mutationFn: registraVoto,
		onSuccess: () => {
			queryClient.invalidateQueries({
				queryKey: ['projects-meeting', reuniaoId],
			});
		},
	});
}
