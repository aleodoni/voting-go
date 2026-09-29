import { useQuery } from '@tanstack/react-query';
import { getApi } from '../api-client';
import type { User } from '../types';

export function useUser(userId: string) {
	return useQuery<User>({
		queryKey: ['user', userId],
		queryFn: async () => {
			const { data } = await getApi().get<User>(`/usuarios/${userId}`);
			return data;
		},
		enabled: !!userId,
	});
}
