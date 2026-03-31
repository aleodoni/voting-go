import { createFileRoute } from '@tanstack/react-router';
import { useAuth } from '@voting/shared';
import { LoggedUserCard } from '@/components/LoggedUserCard';
import { VotingCard } from '@/components/VotingCard';
import { useIsProjectVoting } from '@/hooks/useIsProjectVoting';

export const Route = createFileRoute('/')({
	component: DashboardPage,
});

function DashboardPage() {
	const { user } = useAuth();
	const { data: projectVoting } = useIsProjectVoting();

	return (
		<div className="flex flex-col justify-around flex-1 min-h-0 gap-8">
			<VotingCard projectVoting={projectVoting || null} />
			{user && <LoggedUserCard userInfo={user} />}
		</div>
	);
}
