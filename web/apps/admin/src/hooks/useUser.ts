import { useQuery } from '@tanstack/react-query';
import { getApi, User } from '@voting/shared';

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
