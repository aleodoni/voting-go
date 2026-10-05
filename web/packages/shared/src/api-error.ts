import { isAxiosError } from 'axios';

// Status em que o backend devolve uma mensagem escrita para o usuário (ver
// internal/handler/httperr): sem permissão, não encontrado e regra de negócio
// violada. Em 400 a mensagem é técnica (validação do corpo) e em 5xx é genérica.
const STATUS_COM_MENSAGEM_DO_SERVIDOR = [403, 404, 422];

/**
 * Retorna o texto a mostrar ao usuário quando uma chamada à API falha.
 *
 * Usa a mensagem do servidor ({ "error": "..." }) quando o status é 403, 404 ou
 * 422, por exemplo "já existe uma votação aberta"; caso contrário (erro de rede,
 * 400, 5xx ou erro que não veio do axios) devolve `fallback`.
 */
export function getApiErrorMessage(error: unknown, fallback: string): string {
	if (!isAxiosError(error)) {
		return fallback;
	}

	const status = error.response?.status;
	const data: unknown = error.response?.data;

	if (
		status === undefined ||
		!STATUS_COM_MENSAGEM_DO_SERVIDOR.includes(status)
	) {
		return fallback;
	}

	if (
		typeof data === 'object' &&
		data !== null &&
		'error' in data &&
		typeof data.error === 'string' &&
		data.error !== ''
	) {
		return data.error.charAt(0).toUpperCase() + data.error.slice(1);
	}

	return fallback;
}
