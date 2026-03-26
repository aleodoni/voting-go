import { createFileRoute, useMatch } from '@tanstack/react-router';
import { ContainerPage, H2, Header, useAuth, useTheme } from '@voting/shared';
import { z } from 'zod';
import { SearchUsers } from '@/components/SearchUsers';
import { TableUsers } from '@/components/TableUsers';
export const Route = createFileRoute('/manage-users')({
	component: ManageUsers,
	validateSearch: (search) =>
		z
			.object({
				nome: z.string().optional(),
				email: z.string().optional(),
				page: z.number().optional(),
			})
			.parse(search),
});

function ManageUsers() {
	const { setTheme } = useTheme();
	const { logout } = useAuth();
	const match = useMatch({ from: Route.id });
	const { email, nome, page } = match.search ?? {};

	return (
		<div className="w-full min-h-screen px-6 py-4">
			<div className="max-w-7xl mx-auto flex flex-col gap-4">
				<Header
					subtitulo="Módulo administrativo"
					logout={logout}
					setTheme={setTheme}
				/>
				<ContainerPage>
					<H2>Manutenção de usuários</H2>
					<div className="flex w-full py-8">
						<SearchUsers />
					</div>
					<TableUsers email={email} nome={nome} page={page} />
				</ContainerPage>
			</div>
		</div>
	);
}
