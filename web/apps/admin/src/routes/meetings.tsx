import { createFileRoute } from '@tanstack/react-router';
import { ContainerPage, H2 } from '@voting/shared';
import { useEffect, useState } from 'react';
import { MeetingSelect } from '@/components/MeetingSelect';
import { ProjectsMeeting } from '@/components/MeetingSelect/ProjectsMeeting';
import { useIsProjectVoting } from '@/hooks/useIsProjectVoting';
import { useProjectsMeeting } from '@/hooks/useProjectsMeeting';
import { MeetingDTO, useTodayMeetings } from '@/hooks/useTodayMeetings';

export const Route = createFileRoute('/meetings')({
	component: Meetings,
});

function Meetings() {
	const [selectedMeeting, setSelectedMeeting] = useState<
		MeetingDTO | undefined
	>(undefined);

	// Reuniões de hoje
	const {
		data: todayMeetings,
		isLoading: isMeetingsLoading,
		isFetching: isMeetingsFetching,
		isRefetching: isMeetingsRefetching,
	} = useTodayMeetings();

	// Projetos da reunião selecionada
	const {
		data: projectsMeeting,
		isFetching,
		isRefetching,
	} = useProjectsMeeting(selectedMeeting?.id);

	// Votação aberta (hook não recebe argumentos)
	const { data: openVotingProject } = useIsProjectVoting();

	// Inicializa selectedMeeting com a primeira reunião disponível
	useEffect(() => {
		if (todayMeetings?.length && !selectedMeeting) {
			setSelectedMeeting(todayMeetings[0]);
		}
	}, [todayMeetings, selectedMeeting]);

	function handleMeetingSelect(meeting: MeetingDTO) {
		setSelectedMeeting(meeting);
	}

	return (
		<ContainerPage>
			<H2>Reuniões de hoje, {new Date().toLocaleDateString('pt-BR')}</H2>

			<div className="flex flex-col w-full gap-6 mt-6 lg:flex-row lg:gap-8 lg:items-start">
				<div className="w-full lg:w-1/3 lg:sticky lg:top-6 lg:self-start">
					<MeetingSelect
						meetings={todayMeetings ?? []}
						handleMeetingSelect={handleMeetingSelect}
						selectedMeeting={selectedMeeting}
						isMeetingsFetching={isMeetingsFetching}
						isMeetingsRefetching={isMeetingsRefetching}
						isMeetingsLoading={isMeetingsLoading}
						projects={projectsMeeting ?? []}
					/>
				</div>

				<div className="w-full lg:w-2/3">
					<ProjectsMeeting
						projects={projectsMeeting ?? []}
						isFetching={isFetching}
						isRefetching={isRefetching}
						isVotingProject={!!openVotingProject}
					/>
				</div>
			</div>
		</ContainerPage>
	);
}
