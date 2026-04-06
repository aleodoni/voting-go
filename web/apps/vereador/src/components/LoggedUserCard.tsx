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
			<CardHeader className="p-4 sm:p-6 pb-2 sm:pb-3">
				<CardTitle className="flex items-center gap-2 text-base sm:text-lg">
					<Info className="h-5 w-5 sm:h-6 sm:w-6 text-primary shrink-0" />
					Usuário logado
				</CardTitle>
				<CardDescription className="text-xs sm:text-sm">
					Informações do usuário logado
				</CardDescription>
			</CardHeader>

			<CardFooter className="flex gap-2 justify-between items-center w-full px-4 py-3 sm:px-6 sm:py-4 border-t">
				<div className="flex flex-col justify-center min-w-0">
					<p className="font-bold text-sm sm:text-base truncate">
						{userInfo?.nome_fantasia || userInfo?.nome}
					</p>
					<p className="text-muted-foreground text-xs sm:text-sm truncate">
						{userInfo?.email}
					</p>
				</div>
				<Button asChild className="shrink-0" variant="outline" size="sm">
					<Link to="/user-info">Editar</Link>
				</Button>
			</CardFooter>
		</Card>
	);
}
