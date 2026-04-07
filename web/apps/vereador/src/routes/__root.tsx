import { useQueryClient } from '@tanstack/react-query';
import { createRootRoute, Outlet } from '@tanstack/react-router';
import { Header, useAuth, useSSE, useTheme } from '@voting/shared';

export const Route = createRootRoute({
	component: RootComponent,
});

function RootComponent() {
	const queryClient = useQueryClient();
	const { logout } = useAuth();
	const { setTheme } = useTheme();

	useSSE({
		onConnect: () => {
			queryClient.invalidateQueries({
				queryKey: ['connected-users'],
			});

			queryClient.invalidateQueries({
				queryKey: ['open-voting'],
			});
		},
		onEvent: (event) => {
			switch (event.type) {
				case 'votacao_aberta':
				case 'votacao_fechada':
				case 'votacao_cancelada':
				case 'voto_registrado':
					queryClient.invalidateQueries({
						queryKey: ['open-voting'],
					});
					break;
			}
		},
	});

	return (
		<div className="w-full min-h-screen bg-muted/30 px-6 py-6">
			<div className="max-w-7xl mx-auto flex flex-col gap-6 min-h-full">
				<Header
					subtitulo="Módulo vereador"
					logout={logout}
					setTheme={setTheme}
				/>

				<Outlet />
			</div>
		</div>
	);
}
