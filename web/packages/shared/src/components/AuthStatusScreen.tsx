import { AlertCircle, ShieldAlert, UserX } from 'lucide-react';
import { Button } from './ui/button';
import {
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from './ui/card';

export type AuthStatusVariant = 'forbidden' | 'inactive' | 'failure';

const VARIANTS = {
	forbidden: { title: 'Acesso negado', Icon: ShieldAlert },
	inactive: { title: 'Usuário inativo', Icon: UserX },
	failure: { title: 'Não foi possível entrar', Icon: AlertCircle },
} as const;

type AuthStatusScreenProps = {
	variant: AuthStatusVariant;
	message: string;
	// Quem está conectado: ajuda a perceber que entrou com a conta errada.
	username?: string;
	onLogout: () => void;
	// Só para falhas transitórias (API fora do ar, por exemplo).
	onRetry?: () => void;
};

/**
 * Tela mostrada pelo AuthProvider quando o usuário não pode usar o sistema:
 * sem permissão, conta inativa ou falha ao carregar a sessão.
 */
export function AuthStatusScreen({
	variant,
	message,
	username,
	onLogout,
	onRetry,
}: AuthStatusScreenProps) {
	const { title, Icon } = VARIANTS[variant];

	return (
		<div className="flex min-h-screen items-center justify-center p-4">
			<Card className="w-full max-w-md" role="alert">
				<CardHeader>
					<CardTitle className="flex items-center gap-2 text-lg">
						<Icon
							className="h-6 w-6 shrink-0 text-destructive"
							aria-hidden="true"
						/>
						{title}
					</CardTitle>
					<CardDescription className="text-sm">{message}</CardDescription>
				</CardHeader>

				{username && (
					<CardContent>
						<p className="text-muted-foreground text-sm">
							Conectado como{' '}
							<span className="font-medium text-foreground">{username}</span>
						</p>
					</CardContent>
				)}

				<CardFooter className="flex justify-end gap-2">
					{onRetry && (
						<Button variant="outline" onClick={onRetry}>
							Tentar novamente
						</Button>
					)}
					<Button variant={onRetry ? 'default' : 'outline'} onClick={onLogout}>
						Sair
					</Button>
				</CardFooter>
			</Card>
		</div>
	);
}
