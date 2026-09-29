import { createFileRoute } from '@tanstack/react-router';
import { LoggedUserCard, useAuth, useIsProjectVoting } from '@voting/shared';
import { VotingCard } from '@/components/VotingCard';

export const Route = createFileRoute('/')({
	component: DashboardPage,
});

function DashboardPage() {
	const { user } = useAuth();
	const { data: projectVoting } = useIsProjectVoting();

	return (
		<div className="flex flex-col flex-1 gap-4">
			<VotingCard projectVoting={projectVoting || null} />
			{user && <LoggedUserCard userInfo={user} />}
		</div>
	);
}
