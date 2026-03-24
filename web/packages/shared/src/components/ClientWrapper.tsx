import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { type ReactNode, useState } from 'react';

export function ClientWrapper({ children }: { children: ReactNode }) {
	// Mantemos o useState para garantir que o QueryClient seja instanciado apenas uma vez no ciclo de vida do App
	const [queryClient] = useState(
		() =>
			new QueryClient({
				defaultOptions: {
					queries: {
						// Opcional: Configurações globais para sua API Go
						retry: 1,
						refetchOnWindowFocus: false,
					},
				},
			}),
	);

	return (
		<QueryClientProvider client={queryClient}>
			{children}
			{/* No Vite, usamos import.meta.env.DEV para detectar ambiente de desenvolvimento */}
			{import.meta.env.DEV && <ReactQueryDevtools initialIsOpen={false} />}
		</QueryClientProvider>
	);
}
