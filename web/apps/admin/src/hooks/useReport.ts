import { useMutation } from '@tanstack/react-query';
import { getApi } from '@voting/shared';
import toast from 'react-hot-toast';
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
		// O erro vem como Blob (responseType: 'blob'), então a mensagem do servidor não é legível.
		onError: () => {
			toast.error('Erro ao gerar o relatório');
		},
	});
}
