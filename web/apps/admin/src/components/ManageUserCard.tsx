import { Button, Card, CardDescription, CardFooter, CardHeader, CardTitle } from '@voting/shared'
import { Users } from 'lucide-react';
import { Link } from '@tanstack/react-router';
export function ManageUserCard() {
	return (
		<Card className="w-full h-fit">
			<CardHeader>
				<CardTitle className="flex items-center gap-2">
					<Users className="h-8 w-8 rounded-full border-2 border-primary text-primary " />
					Gerenciar usuários
				</CardTitle>
				<CardDescription>Gerencie os usuários do sistema</CardDescription>
			</CardHeader>
			<CardFooter className="flex gap-2 justify-end items-end w-full h-16">
				<Button asChild className="justify-center" variant={'outline'}>
					<Link to="/manage-users">Gerenciar</Link>
				</Button>
			</CardFooter>
		</Card>
	);
}
