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
	const abortRef = useRef<AbortController | null>(null);
	const isConnectingRef = useRef(false); // ← novo: evita conexões paralelas
	const onConnectRef = useRef(onConnect);
	const onEventRef = useRef(onEvent);
	const onErrorRef = useRef(onError);

	useEffect(() => {
		onConnectRef.current = onConnect;
		onEventRef.current = onEvent;
		onErrorRef.current = onError;
	}, [onConnect, onEvent, onError]);

	useEffect(() => {
		const keycloak = getKeycloak();

		const connect = async () => {
			// ← Aborta conexão anterior antes de criar nova
			if (abortRef.current) {
				abortRef.current.abort();
				abortRef.current = null;
			}

			if (isConnectingRef.current) return; // ← evita chamadas paralelas
			isConnectingRef.current = true;

			try {
				await keycloak?.updateToken(30);
			} catch {
				isConnectingRef.current = false;
				keycloak?.login();
				return;
			}

			const token = keycloak?.token;
			if (!token) {
				isConnectingRef.current = false;
				console.warn('SSE: No token available, retrying in 3s...');
				setTimeout(connect, 3000);
				return;
			}

			const url = `${import.meta.env.VITE_API_URL}/eventos?token=${token}`;
			const controller = new AbortController();
			abortRef.current = controller;
			isConnectingRef.current = false; // ← libera após criar o controller

			try {
				const response = await fetch(url, {
					signal: controller.signal,
					headers: { Accept: 'text/event-stream' },
				});

				if (!response.ok || !response.body) {
					throw new Error(`SSE response error: ${response.status}`);
				}

				onConnectRef.current?.();

				const reader = response.body.getReader();
				const decoder = new TextDecoder();
				let buffer = '';

				while (true) {
					const { done, value } = await reader.read();
					if (done) break;

					buffer += decoder.decode(value, { stream: true });
					const lines = buffer.split('\n');
					buffer = lines.pop() ?? '';

					let eventType = '';
					let eventData = '';

					for (const line of lines) {
						if (line.startsWith('event:')) {
							eventType = line.slice(6).trim();
						} else if (line.startsWith('data:')) {
							eventData = line.slice(5).trim();
						} else if (line === '' && eventType) {
							const type = eventType as SSEEventType;
							if (SSE_EVENTS.includes(type)) {
								try {
									const payload = eventData ? JSON.parse(eventData) : undefined;
									onEventRef.current({ type, payload });
								} catch {
									onEventRef.current({ type });
								}
							}
							eventType = '';
							eventData = '';
						}
					}
				}
			} catch (err) {
				if ((err as Error).name === 'AbortError') return;
				console.error('SSE error:', err);
				onErrorRef.current?.(err as Event);
				setTimeout(connect, 3000);
			}
		};

		connect();

		return () => {
			abortRef.current?.abort();
			abortRef.current = null;
			isConnectingRef.current = false; // ← limpa flag no unmount
		};
	}, []);
}
