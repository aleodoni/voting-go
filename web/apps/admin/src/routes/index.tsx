import { createFileRoute } from '@tanstack/react-router';
import { useAuth } from '@voting/shared';
import { ConnectedUsersCard } from '@/components/ConnectedUsersCard';
import { LastSyncsCard } from '@/components/LastSynchsCard';
import { LoggedUserCard } from '@/components/LoggedUserCard';
import { ManageUserCard } from '@/components/ManageUserCard';
import { MeetingsCard } from '@/components/MeetingsCard';
import { VotingProgress } from '@/components/VotingProgress';
import { VotingStatsTotalCard } from '@/components/VotingStatsTotalCard';
import { useTodayMeetings } from '@/hooks/useTodayMeetings';
import { useVotingStats } from '@/hooks/useVotingStats';

export const Route = createFileRoute('/')({
	component: DashboardPage,
});

function DashboardPage() {
	const { data: todayMeetings, isLoading } = useTodayMeetings();
	const { data: votingStats } = useVotingStats();
	const { user } = useAuth();

	return (
		<div className="flex flex-col gap-4 flex-1">
			{/* Meetings */}
			<MeetingsCard meetings={todayMeetings ?? []} isLoading={isLoading} />

			{/* Métricas */}
			<div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
				<ConnectedUsersCard />
				<VotingStatsTotalCard stats={votingStats} isLoading={!votingStats} />
				<VotingProgress stats={votingStats} isLoading={!votingStats} />
			</div>

			<LastSyncsCard />
			{/* Administração */}
			<div className="grid grid-cols-1 gap-4 lg:grid-cols-2">
				{user && <LoggedUserCard userInfo={user} />}
				<ManageUserCard />
			</div>
		</div>
	);
}
