import { useMutation } from '@tanstack/react-query';
import { getApi } from '@voting/shared';
import { MeetingDTO } from '@/hooks/useTodayMeetings';

async function fetchMeetingReport(meeting: MeetingDTO) {
	const response = await getApi().get(`/reunioes/${meeting.id}/relatorio`, {
		responseType: 'blob', // PDF
	});
	return response.data;
}

export function useMeetingReport() {
	return useMutation({
		mutationFn: fetchMeetingReport,
		onSuccess: (blob) => {
			// Abrir PDF em nova aba
			const url = window.URL.createObjectURL(blob);
			window.open(url, '_blank');

			// Se quiser salvar o PDF em vez de abrir, descomente abaixo:
			/*
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', 'relatorio.pdf');
      document.body.appendChild(link);
      link.click();
      link.remove();
      */

			window.URL.revokeObjectURL(url);
		},
	});
}
