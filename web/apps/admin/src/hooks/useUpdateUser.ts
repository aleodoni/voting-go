import { useMutation, useQueryClient } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

type UpdateUserData = {
	nome_fantasia: string;
	ativo: boolean;
	pode_administrar: boolean;
	pode_votar: boolean;
};

export function useUpdateUser() {
	const queryClient = useQueryClient();

	const mutation = useMutation({
		mutationFn: async ({
			userId,
			data,
		}: {
			userId: string;
			data: UpdateUserData;
		}) => {
			const { data: response } = await getApi().put(
				'/usuarios/fantasia-credenciais',
				{
					user_id: userId,
					display_name: data.nome_fantasia,
					is_active: data.ativo,
					can_admin: data.pode_administrar,
					can_vote: data.pode_votar,
				},
			);
			return response;
		},
		onSuccess: (_, { userId }) => {
			queryClient.invalidateQueries({ queryKey: ['user', userId] });
			queryClient.invalidateQueries({ queryKey: ['users'] });
		},
	});

	return {
		updateUser: (userId: string, data: UpdateUserData) =>
			mutation.mutateAsync({ userId, data }),
		isPending: mutation.isPending,
	};
}
