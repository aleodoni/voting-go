import { createFileRoute } from '@tanstack/react-router';
import { ContainerPage, H2, P, useAuth } from '@voting/shared';
import { FormUserInfo } from '@/components/FormUserInfo';

export const Route = createFileRoute('/user-info')({
	component: UserInfo,
});

function UserInfo() {
	const { user } = useAuth();

	return (
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
	);
}
