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
import { UserItem } from './UserItem';

type TableUserProps = {
	nome?: string;
	email?: string;
	page?: number;
};

export function TableUsers({
	email = '',
	nome = '',
	page = 1,
}: TableUserProps) {
	const { data: users, isLoading } = useUsers(nome, email, page);

	if (isLoading) {
		return (
			<div className="flex flex-col w-full h-full items-center justify-center">
				<p className="font-bold text-2xl text-muted-foreground">
					Carregando...
				</p>
			</div>
		);
	}

	if (users && users.usuarios.length > 0) {
		return (
			<Table>
				<TableCaption>Listagem de usuários</TableCaption>
				<TableHeader>
					<TableRow>
						<TableHead>Nome</TableHead>
						<TableHead>Nome no sistema</TableHead>
						<TableHead>e-mail</TableHead>
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
		);
	}

	return (
		<div className="flex flex-col w-full h-full items-center justify-center">
			<p className="font-bold text-2xl text-muted-foreground">
				Nenhum usuário encontrado
			</p>
		</div>
	);
}
