import { ConnectedUsersCard } from '@/components/ConnectedUsersCard'
import { MeetingsCard } from '@/components/MeetingsCard'
import { VotingProgress } from '@/components/VotingProgress'
import { VotingStatsTotalCard } from '@/components/VotingStatsTotalCard'
import { useTodayMeetings } from '@/hooks/useTodayMeetings'
import { useVotingStats } from '@/hooks/useVotingStats'
import { createFileRoute } from '@tanstack/react-router'
import { Header, useAuth, useTheme } from '@voting/shared'

export const Route = createFileRoute('/')({
  component: DashboardPage,
})

function DashboardPage() {
  const { logout } = useAuth()
  const { setTheme } = useTheme()
  const { data: todayMeetings, isLoading } = useTodayMeetings();
  const { data: votingStats } = useVotingStats()

  return (
    <div className="flex flex-col w-full h-full gap-4">
      <Header subtitulo="Módulo administrativo" logout={logout} setTheme={setTheme}/>
      <div className="col-span-2">
					<MeetingsCard meetings={todayMeetings ?? []} isLoading={isLoading} />
				</div>
      <div className="flex flex-row h-full col-span-2 gap-4 items-start">
        <ConnectedUsersCard/>
        <VotingStatsTotalCard stats={votingStats} isLoading={!votingStats}/>
        <VotingProgress stats={votingStats} isLoading={!votingStats}/>
      </div>
    </div>
  )

  
}