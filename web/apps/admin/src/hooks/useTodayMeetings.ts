import { useQuery } from '@tanstack/react-query';
import { getApi } from '@voting/shared';

export type MeetingDTO = {
	id: string;
	con_id: number;
	con_desc: string;
	rec_id: number;
	con_sigla: string;
	rec_tipo_reuniao: string;
	rec_numero: string;
	pac_id: number;
	rec_data: string;
};

async function fetchMeetings(): Promise<MeetingDTO[]> {
	const { data } = await getApi().get<MeetingDTO[]>('/reunioes-dia');
	return data;
}

export function useTodayMeetings() {
	return useQuery({
		queryKey: ['today-meetings'],
		queryFn: fetchMeetings,
		refetchOnReconnect: true,
		staleTime: 0,
	});
}
