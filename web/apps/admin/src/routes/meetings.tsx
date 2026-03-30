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
		(_event: SSEEvent) => {
			queryClient.invalidateQueries({
				queryKey: ['projects-meeting', selectedMeeting?.id],
			});
		},
		[queryClient, selectedMeeting?.id],
	);

	useSSE({ onEvent: handleVotacaoEvent });

	const {
		data: projectsMeeting,
		isFetching,
		isRefetching,
	} = useProjectsMeeting(selectedMeeting?.id);

	return (
		<ContainerPage>
			<H2>Reuniões de hoje, {new Date().toLocaleDateString('pt-BR')}</H2>

			<div className="flex w-full gap-8 mt-6 items-start">
				<div className="w-1/3 sticky top-6 self-start">
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

				<div className="w-2/3">
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
