import { createFileRoute } from '@tanstack/react-router';
import { useAuth } from '@voting/shared';
import { ConnectedUsersCard } from '@/components/ConnectedUsersCard';
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
		<div className="grid grid-cols-1 gap-4 flex-1 sm:grid-cols-2 lg:grid-cols-3 lg:auto-rows-fr">
			{/* MeetingsCard — ocupa linha inteira em todos os tamanhos */}
			<div className="flex col-span-1 sm:col-span-2 lg:col-span-3">
				<MeetingsCard meetings={todayMeetings ?? []} isLoading={isLoading} />
			</div>

			{/* Em mobile: 1 coluna. Em sm: 2 colunas. Em lg: 3 colunas */}
			<ConnectedUsersCard />
			<VotingStatsTotalCard stats={votingStats} isLoading={!votingStats} />
			<VotingProgress stats={votingStats} isLoading={!votingStats} />

			{/* LoggedUserCard — visível apenas se houver usuário */}
			{user && <LoggedUserCard userInfo={user} />}

			{/* ManageUserCard — ocupa 2 colunas a partir de sm */}
			<div className="flex col-span-1 sm:col-span-2">
				<ManageUserCard />
			</div>
		</div>
	);
}
