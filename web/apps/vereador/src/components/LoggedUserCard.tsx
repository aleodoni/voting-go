import { Link } from '@tanstack/react-router';
import {
	Button,
	Card,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
	User,
} from '@voting/shared';
import { Info } from 'lucide-react';

type LoggedUserCardProps = {
	userInfo: User;
};

export function LoggedUserCard({ userInfo }: LoggedUserCardProps) {
	return (
		<Card className="w-full h-fit">
			<CardHeader>
				<CardTitle className="flex items-center gap-2">
					<Info className="h-8 w-8 rounded-full border-primary text-primary" />
					Usuário logado
				</CardTitle>
				<CardDescription>Informações do usuário logado</CardDescription>
			</CardHeader>
			<CardFooter className="flex gap-2 justify-between items-end w-full h-16">
				<div className="flex flex-col justify-center">
					<p className="font-bold">
						{userInfo?.nome_fantasia || userInfo?.nome}
					</p>
					<p className="text-muted-foreground text-sm">{userInfo?.email}</p>
				</div>
				<Button asChild className="justify-center" variant="outline">
					<Link to="/user-info">Editar</Link>
				</Button>
			</CardFooter>
		</Card>
	);
}
