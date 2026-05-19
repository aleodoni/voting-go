import { LogOut } from 'lucide-react';
import { Button } from './ui/button';

interface ButtonLogoutProps {
	logout: () => void;
}

function ButtonLogout({ logout }: ButtonLogoutProps) {
	return (
		<Button variant={'outline'} onClick={logout}>
			Sair
			<LogOut className="h-5 w-5 lg:w-6 lg:h-6" />
		</Button>
	);
}

export { ButtonLogout };
