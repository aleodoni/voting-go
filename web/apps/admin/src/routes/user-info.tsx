import { createFileRoute } from '@tanstack/react-router';
import {
	ContainerPage,
	H2,
	Header,
	P,
	useAuth,
	useTheme,
} from '@voting/shared';
import { FormUserInfo } from '@/components/FormUserInfo';

export const Route = createFileRoute('/user-info')({
	component: UserInfo,
});

function UserInfo() {
	const { setTheme } = useTheme();
	const { user, logout } = useAuth();

	return (
		<div className="w-full min-h-screen px-6 py-4">
			<div className="max-w-7xl mx-auto flex flex-col gap-4">
				<Header
					subtitulo="Módulo administrativo"
					logout={logout}
					setTheme={setTheme}
				/>
				<ContainerPage>
					<H2>Usuário logado</H2>
					<P>
						<b>Usuário: </b>
						{user?.nome}
					</P>
					<P>
						<b>email: </b>
						{user?.email}
					</P>
					<div className="flex w-full items-center gap-2 mt-8">
						{user && <FormUserInfo userInfo={user} />}
					</div>
				</ContainerPage>
			</div>
		</div>
	);
}
