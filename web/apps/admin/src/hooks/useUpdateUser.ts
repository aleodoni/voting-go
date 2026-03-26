import { useMutation } from '@tanstack/react-query';

type UpdateUserData = {
	nome: string;
	email: string;
	nome_fantasia: string;
};

export function useUpdateUser() {
	const mutation = useMutation({
		mutationFn: async ({
			userId,
			data,
		}: {
			userId: string;
			data: UpdateUserData;
		}) => {
			// chamada API
			console.log(userId, data);
		},
	});

	return {
		updateUser: (userId: string, data: UpdateUserData) =>
			mutation.mutateAsync({ userId, data }),

		isPending: mutation.isPending,
	};
}
