import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from '@voting/shared';
import { Info } from 'lucide-react';
import { useConnectedUsers } from '../hooks/useConnectedUsers';

export function ConnectedUsersCard() {
	const { data: users = [], isLoading } = useConnectedUsers();

	const vereadores = users.filter((u) => !u.is_admin);
	const admins = users.filter((u) => u.is_admin);

	return (
		<Card className="w-full">
			<CardHeader>
				<CardTitle className="flex items-center gap-2">
					<Info className="h-8 w-8 rounded-full border-primary text-primary" />
					Usuários conectados
				</CardTitle>
				<CardDescription>
					{isLoading
						? 'Carregando...'
						: `${admins.length} admin(s) · ${vereadores.length} vereador(es)`}
				</CardDescription>
			</CardHeader>
			<CardContent>
				{isLoading ? (
					<p className="text-sm text-muted-foreground">
						Carregando usuários...
					</p>
				) : (
					users.map((user) => (
						<div
							key={user.user_id}
							className="flex gap-2 justify-between items-center w-full py-0.5"
						>
							<p className="text-sm">{user.username}</p>
							<span className="text-xs text-muted-foreground">
								{user.is_admin ? 'admin' : 'vereador'}
							</span>
						</div>
					))
				)}
			</CardContent>
		</Card>
	);
}
