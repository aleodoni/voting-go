import { createFileRoute } from '@tanstack/react-router';
import { ContainerPage, H2, useAuth } from '@voting/shared';
import { UserInfoPage } from '@/components/UserInfoPage';

export const Route = createFileRoute('/user-info')({
	component: UserInfo,
});

function UserInfo() {
	const { user } = useAuth();

	if (!user) return null;

	return (
		<ContainerPage>
			<H2>Meu perfil</H2>
			<UserInfoPage userInfo={user} />
		</ContainerPage>
	);
}
