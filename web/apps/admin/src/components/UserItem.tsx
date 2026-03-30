import { useNavigate, useSearch } from '@tanstack/react-router';
import { TableCell, TableRow, User } from '@voting/shared';
import { PermissionIcon } from './PermissionIcon';

type UserItemProps = {
	user: User;
};

export function UserItem({ user }: UserItemProps) {
	const navigate = useNavigate();
	const search = useSearch({ from: '/manage-users' });

	function handleClick() {
		navigate({
			to: '/user/$userId',
			params: { userId: user.id },
			search: { returnSearch: search },
		});
	}

	return (
		<TableRow
			className="cursor-pointer hover:bg-muted/50 transition-colors"
			onClick={handleClick}
			title="Editar usuário"
		>
			<TableCell>{user.nome}</TableCell>
			<TableCell>{user.nome_fantasia}</TableCell>
			<TableCell>{user.email}</TableCell>
			<TableCell>
				<PermissionIcon status={user.credencial.ativo} />
			</TableCell>
			<TableCell>
				<PermissionIcon status={user.credencial.pode_administrar} />
			</TableCell>
			<TableCell>
				<PermissionIcon status={user.credencial.pode_votar} />
			</TableCell>
		</TableRow>
	);
}
