import {
	Pagination,
	PaginationContent,
	PaginationEllipsis,
	PaginationItem,
	PaginationLink,
	PaginationNext,
	PaginationPrevious,
	Table,
	TableBody,
	TableCaption,
	TableHead,
	TableHeader,
	TableRow,
	User,
} from '@voting/shared';
import { useUsers } from '@/hooks/useUsers';
import { UserCard } from './UserCard';
import { UserItem } from './UserItem';

type TableUserProps = {
	nome?: string;
	email?: string;
	page?: number;
	listarInativos?: boolean;
	onPageChange?: (page: number) => void;
};

export function TableUsers({
	email = '',
	nome = '',
	page = 1,
	listarInativos = false,
	onPageChange,
}: TableUserProps) {
	const { data: users, isLoading } = useUsers(
		nome,
		email,
		listarInativos,
		page,
		10,
	);

	if (isLoading) {
		return (
			<div className="flex flex-col w-full h-full items-center justify-center py-10">
				<p className="font-bold text-xl text-muted-foreground">Carregando...</p>
			</div>
		);
	}

	if (!users || users.usuarios.length === 0) {
		return (
			<div className="flex flex-col w-full h-full items-center justify-center py-10">
				<p className="font-bold text-xl text-muted-foreground">
					Nenhum usuário encontrado
				</p>
			</div>
		);
	}

	const totalPages = Math.ceil(users.total / users.limit);

	function getPageNumbers(
		current: number,
		total: number,
	): (number | 'ellipsis')[] {
		if (total <= 5) return Array.from({ length: total }, (_, i) => i + 1);

		if (current <= 3) return [1, 2, 3, 4, 'ellipsis', total];
		if (current >= total - 2)
			return [1, 'ellipsis', total - 3, total - 2, total - 1, total];

		return [
			1,
			'ellipsis',
			current - 1,
			current,
			current + 1,
			'ellipsis',
			total,
		];
	}

	const pageNumbers = getPageNumbers(page, totalPages);

	return (
		<div className="flex flex-col gap-4">
			{/* Mobile: cards */}
			<div className="flex flex-col gap-3 md:hidden">
				{users.usuarios.map((user: User) => (
					<UserCard key={user.id} user={user} />
				))}
			</div>

			{/* Desktop: tabela */}
			<div className="hidden md:block">
				<Table>
					<TableCaption>Listagem de usuários</TableCaption>
					<TableHeader>
						<TableRow>
							<TableHead>Nome</TableHead>
							<TableHead>Nome no sistema</TableHead>
							<TableHead>E-mail</TableHead>
							<TableHead>Ativo</TableHead>
							<TableHead>Admin</TableHead>
							<TableHead>Vota</TableHead>
						</TableRow>
					</TableHeader>
					<TableBody>
						{users.usuarios.map((user: User) => (
							<UserItem key={user.id} user={user} />
						))}
					</TableBody>
				</Table>
			</div>

			{/* Paginação */}
			{totalPages > 1 && (
				<Pagination>
					<PaginationContent>
						<PaginationItem>
							<PaginationPrevious
								href="#"
								onClick={(e) => {
									e.preventDefault();
									if (page > 1) onPageChange?.(page - 1);
								}}
								aria-disabled={page === 1}
								className={
									page === 1
										? 'pointer-events-none opacity-50'
										: 'cursor-pointer'
								}
							/>
						</PaginationItem>

						{pageNumbers.map((p) =>
							p === 'ellipsis' ? (
								<PaginationItem
									key={`ellipsis-${p === pageNumbers[1] ? 'start' : 'end'}`}
								>
									<PaginationEllipsis />
								</PaginationItem>
							) : (
								<PaginationItem key={p}>
									<PaginationLink
										href="#"
										isActive={p === page}
										onClick={(e) => {
											e.preventDefault();
											onPageChange?.(p);
										}}
										className="cursor-pointer"
									>
										{p}
									</PaginationLink>
								</PaginationItem>
							),
						)}

						<PaginationItem>
							<PaginationNext
								href="#"
								onClick={(e) => {
									e.preventDefault();
									if (page < totalPages) onPageChange?.(page + 1);
								}}
								aria-disabled={page === totalPages}
								className={
									page === totalPages
										? 'pointer-events-none opacity-50'
										: 'cursor-pointer'
								}
							/>
						</PaginationItem>
					</PaginationContent>
				</Pagination>
			)}
		</div>
	);
}
