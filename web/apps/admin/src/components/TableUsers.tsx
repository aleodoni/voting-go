import {
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
};

export function TableUsers({
	email = '',
	nome = '',
	page = 1,
	listarInativos = false,
}: TableUserProps) {
	const { data: users, isLoading } = useUsers(
		nome,
		email,
		listarInativos,
		page,
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

	return (
		<>
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
		</>
	);
}
