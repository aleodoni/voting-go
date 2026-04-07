import { useQueryClient } from '@tanstack/react-query';
import { createFileRoute } from '@tanstack/react-router';
import { ContainerPage, H2, SSEEvent, useSSE } from '@voting/shared';
import { useCallback, useEffect, useState } from 'react';
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

	const queryClient = useQueryClient();

	const {
		data: todayMeetings,
		isLoading: isMeetingsLoading,
		isFetching: isMeetingsFetching,
		isRefetching: isMeetingsRefetching,
	} = useTodayMeetings();

	const { data: openVotingProject } = useIsProjectVoting();

	useEffect(() => {
		if (todayMeetings?.length && !selectedMeeting) {
			setSelectedMeeting(todayMeetings[0]);
		}
	}, [todayMeetings, selectedMeeting]);

	function handleMeetingSelect(meeting: MeetingDTO) {
		setSelectedMeeting(meeting);
	}

	const handleVotacaoEvent = useCallback(
		(event: SSEEvent) => {
			if (!selectedMeeting?.id) return;

			switch (event.type) {
				case 'votacao_aberta':
				case 'votacao_fechada':
				case 'votacao_cancelada':
				case 'voto_registrado':
					queryClient.invalidateQueries({
						queryKey: ['projects-meeting', selectedMeeting.id],
					});
					break;
			}
		},
		[queryClient, selectedMeeting?.id],
	);

	useSSE({
		onConnect: () => {
			if (selectedMeeting?.id) {
				queryClient.invalidateQueries({
					queryKey: ['projects-meeting', selectedMeeting.id],
				});
			}
		},
		onEvent: handleVotacaoEvent,
	});

	const {
		data: projectsMeeting,
		isFetching,
		isRefetching,
	} = useProjectsMeeting(selectedMeeting?.id);

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
