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
	listarInativos?: boolean;
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
			listarInativos: params.listarInativos ?? false,
			page: params.page ?? 1,
			limit: params.limit ?? 10,
		},
	});

	return data;
}

export function useUsers(
	nome = '',
	email = '',
	listarInativos = false,
	page = 1,
	limit = 10,
) {
	return useQuery({
		queryKey: ['users', nome, email, listarInativos, page, limit],
		queryFn: () =>
			fetchUsers({
				nome,
				email,
				listarInativos,
				page,
				limit,
			}),
		staleTime: 0,
	});
}
