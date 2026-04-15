import { createFileRoute, useMatch, useNavigate } from '@tanstack/react-router';
import { ContainerPage, H2 } from '@voting/shared';
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
				listarInativos: z.boolean().optional(),
				page: z.number().optional(),
			})
			.parse(search),
});

function ManageUsers() {
	const navigate = useNavigate({ from: Route.id });
	const match = useMatch({ from: Route.id });
	const { email, nome, page, listarInativos } = match.search ?? {};

	function handlePageChange(newPage: number) {
		navigate({
			search: (prev: Record<string, unknown>) => ({ ...prev, page: newPage }),
		});
	}

	return (
		<ContainerPage>
			<H2>Manutenção de usuários</H2>
			<div className="flex w-full py-8">
				<SearchUsers />
			</div>
			<TableUsers
				email={email}
				nome={nome}
				page={page}
				listarInativos={listarInativos}
				onPageChange={handlePageChange}
			/>
		</ContainerPage>
	);
}
