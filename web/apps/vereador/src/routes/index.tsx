import { createFileRoute } from '@tanstack/react-router';
import { useAuth } from '@voting/shared';
import { LoggedUserCard } from '@/components/LoggedUserCard';

export const Route = createFileRoute('/')({
	component: DashboardPage,
});

function DashboardPage() {
	const { user } = useAuth();

	return (
		<div className="grid grid-cols-3 gap-4 flex-1 auto-rows-fr">
			<div className="flex col-span-3"></div>
			{user && <LoggedUserCard userInfo={user} />}
		</div>
	);
}
