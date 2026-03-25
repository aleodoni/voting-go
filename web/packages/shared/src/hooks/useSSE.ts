import { useEffect, useRef } from 'react';
import { getKeycloak } from '../keycloak';

export type SSEEventType =
	| 'votacao_aberta'
	| 'votacao_fechada'
	| 'votacao_cancelada'
	| 'voto_registrado';

export type SSEEvent = {
	type: SSEEventType;
	payload?: unknown;
};

type UseSSEOptions = {
	onConnect?: () => void;
	onEvent: (event: SSEEvent) => void;
	onError?: (error: Event) => void;
};

const SSE_EVENTS: SSEEventType[] = [
	'votacao_aberta',
	'votacao_fechada',
	'votacao_cancelada',
	'voto_registrado',
];

export function useSSE({ onConnect, onEvent, onError }: UseSSEOptions) {
	// 1. Pegamos o token no topo do hook para usá-lo como dependência
	const keycloak = getKeycloak();
	const token = keycloak.token;

	// 2. Refs para manter os callbacks sempre atualizados sem reiniciar o efeito
	const onConnectRef = useRef(onConnect);
	const onEventRef = useRef(onEvent);
	const onErrorRef = useRef(onError);

	useEffect(() => {
		onConnectRef.current = onConnect;
		onEventRef.current = onEvent;
		onErrorRef.current = onError;
	}, [onConnect, onEvent, onError]);

	// 3. Efeito principal que gerencia a conexão
	useEffect(() => {
		// Se não houver token, não iniciamos a conexão
		if (!token) return;

		const url = `${import.meta.env.VITE_API_URL}/eventos?token=${token}`;
		const eventSource = new EventSource(url);

		eventSource.onopen = () => {
			onConnectRef.current?.();
		};

		SSE_EVENTS.forEach((type) => {
			eventSource.addEventListener(type, (e: MessageEvent) => {
				try {
					const payload = e.data ? JSON.parse(e.data) : undefined;
					onEventRef.current({ type, payload });
				} catch (_err) {
					onEventRef.current({ type });
				}
			});
		});

		eventSource.onerror = (e) => {
			onErrorRef.current?.(e);
		};

		return () => {
			eventSource.close();
		};
		// Agora o 'token' existe neste escopo e o efeito reinicia se ele mudar
	}, [token]);
}
