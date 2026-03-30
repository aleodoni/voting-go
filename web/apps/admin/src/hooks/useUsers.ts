import { useQuery } from '@tanstack/react-query';
import type { User } from '@voting/shared';
import { getApi } from '@voting/shared';

type UsersResponse = {
	usuarios: User[];
	total: number;
	page: number;
	limit: number;
};

type FetchUsersParams = {
	nome?: string;
	email?: string;
	page?: number;
	limit?: number;
};

async function fetchUsers(
	params: FetchUsersParams = {},
): Promise<UsersResponse> {
	const { data } = await getApi().get<UsersResponse>('/usuarios', {
		params: {
			Nome: params.nome ?? '',
			Email: params.email ?? '',
			page: params.page ?? 1,
			limit: params.limit ?? 10,
		},
	});

	return data;
}

export function useUsers(nome = '', email = '', page = 1) {
	return useQuery({
		queryKey: ['users', nome, email, page],
		queryFn: () =>
			fetchUsers({
				nome,
				email,
				page,
				limit: 10,
			}),
		staleTime: 0,
	});
}
