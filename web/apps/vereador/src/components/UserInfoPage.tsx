import {
	Badge,
	Card,
	CardContent,
	CardHeader,
	CardTitle,
	User,
} from '@voting/shared';
import { User as UserIcon } from 'lucide-react';
import { FormUserInfo } from './FormUserInfo';

type UserInfoPageProps = {
	userInfo: User;
};

export function UserInfoPage({ userInfo }: UserInfoPageProps) {
	return (
		<Card className="w-full h-fit mt-4 sm:mt-8">
			<CardHeader className="pb-2 sm:pb-4 p-4 sm:p-6">
				<CardTitle className="flex items-center gap-2 text-base sm:text-lg text-foreground/80">
					<UserIcon className="h-5 w-5 shrink-0" />
					Informações Pessoais
				</CardTitle>
			</CardHeader>

			<CardContent className="flex flex-col gap-4 p-4 sm:p-6 pt-0 sm:pt-0">
				{/* Dados do usuário */}
				<div className="flex flex-col gap-1.5 text-sm text-foreground/60 border-b pb-4">
					<span>Nome: {userInfo.nome}</span>
					<span>E-mail: {userInfo.email}</span>

					<div className="flex flex-wrap gap-2 mt-1">
						{userInfo.credencial.ativo ? (
							<Badge className="bg-primary text-primary-foreground">
								Usuário ativo
							</Badge>
						) : (
							<Badge variant="outline">Usuário inativo</Badge>
						)}

						{userInfo.credencial.pode_votar ? (
							<Badge className="bg-primary text-primary-foreground">
								Pode votar
							</Badge>
						) : (
							<Badge variant="outline">Sem permissão para votar</Badge>
						)}

						{userInfo.credencial.pode_administrar ? (
							<Badge className="bg-primary text-primary-foreground">
								Admin
							</Badge>
						) : (
							<Badge variant="outline" className="text-foreground/30">
								Sem permissão admin
							</Badge>
						)}
					</div>
				</div>

				{/* Formulário */}
				<FormUserInfo userInfo={userInfo} />
			</CardContent>
		</Card>
	);
}
