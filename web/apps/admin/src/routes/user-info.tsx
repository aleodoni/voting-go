import { createFileRoute } from '@tanstack/react-router';
import { Header, useAuth, useTheme } from '@voting/shared';
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
				{user && <FormUserInfo userInfo={user} />}
			</div>
		</div>
	);
}
