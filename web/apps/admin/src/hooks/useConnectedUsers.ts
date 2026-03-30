import { useQuery } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

type ConnectedUser = {
	user_id: string;
	username: string;
	is_admin: boolean;
};

async function fetchConnectedUsers(): Promise<ConnectedUser[]> {
	const { data } = await getApi().get<ConnectedUser[]>('/usuarios-conectados');
	return data;
}

export function useConnectedUsers() {
	return useQuery({
		queryKey: ['connected-users'],
		queryFn: fetchConnectedUsers,
		refetchInterval: 10000,
		staleTime: 0,
	});
}
