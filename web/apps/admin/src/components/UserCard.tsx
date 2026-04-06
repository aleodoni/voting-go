import { useNavigate, useSearch } from '@tanstack/react-router';
import { Card, CardContent, User } from '@voting/shared';
import { PermissionIcon } from './PermissionIcon';

type UserCardProps = {
	user: User;
};

export function UserCard({ user }: UserCardProps) {
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
		<Card
			className="w-full cursor-pointer hover:shadow-md transition-all border-muted"
			onClick={handleClick}
			title="Editar usuário"
		>
			<CardContent className="p-4 space-y-3">
				<div>
					<p className="font-semibold text-sm">{user.nome}</p>
					<p className="text-xs text-muted-foreground">{user.nome_fantasia}</p>
					<p className="text-xs text-muted-foreground">{user.email}</p>
				</div>

				<div className="flex items-center gap-4 pt-1 border-t">
					<div className="flex items-center gap-1.5">
						<span className="text-xs text-muted-foreground">Ativo</span>
						<PermissionIcon status={user.credencial.ativo} />
					</div>
					<div className="flex items-center gap-1.5">
						<span className="text-xs text-muted-foreground">Admin</span>
						<PermissionIcon status={user.credencial.pode_administrar} />
					</div>
					<div className="flex items-center gap-1.5">
						<span className="text-xs text-muted-foreground">Vota</span>
						<PermissionIcon status={user.credencial.pode_votar} />
					</div>
				</div>
			</CardContent>
		</Card>
	);
}
