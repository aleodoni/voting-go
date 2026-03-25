import { ConnectedUsersCard } from '@/components/ConnectedUsersCard'
import { LoggedUserCard } from '@/components/LoggedUserCard'
import { ManageUserCard } from '@/components/ManageUserCard'
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
  const { data: todayMeetings, isLoading } = useTodayMeetings()
  const { data: votingStats } = useVotingStats()
  const { user } = useAuth()

  return (
    <div className="w-full min-h-screen px-6 py-4">
      <div className="max-w-7xl mx-auto flex flex-col gap-4">
        <Header
          subtitulo="Módulo administrativo"
          logout={logout}
          setTheme={setTheme}
        />

        <div className="grid grid-cols-3 gap-4 flex-1 auto-rows-fr">
          <div className="flex col-span-3">
            <MeetingsCard meetings={todayMeetings ?? []} isLoading={isLoading} />
          </div>

          <ConnectedUsersCard />
          <VotingStatsTotalCard stats={votingStats} isLoading={!votingStats} />
          <VotingProgress stats={votingStats} isLoading={!votingStats} />

          {user && <LoggedUserCard userInfo={user} />}
          <div className="flex col-span-2">
            <ManageUserCard/>
          </div>

        </div>
      </div>
    </div>
  )
}