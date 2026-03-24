import { MeetingsCard } from '@/components/MeetingsCard'
import { useTodayMeetings } from '@/hooks/useTodayMeetings'
import { createFileRoute } from '@tanstack/react-router'
import { Header, useAuth, useTheme } from '@voting/shared'

export const Route = createFileRoute('/')({
  component: DashboardPage,
})

function DashboardPage() {
  const { logout } = useAuth()
  const { setTheme } = useTheme()
  const { data: todayMeetings, isLoading } = useTodayMeetings();

  return (
    <div className="flex flex-col w-full h-full">
      <Header subtitulo="Módulo administrativo" logout={logout} setTheme={setTheme}/>
      <div className="col-span-2">
					<MeetingsCard meetings={todayMeetings ?? []} isLoading={isLoading} />
				</div>
    </div>
  )

  
}