import { useMeetingReport } from '@/hooks/useReport';
import { MeetingDTO } from '@/hooks/useTodayMeetings';
import { ProjectDTO } from '@/types/meeting';
import { MeetingCard } from './MeetingCard';

type MeetingSelectProps = {
	handleMeetingSelect: (meeting: MeetingDTO) => void;
	selectedMeeting?: MeetingDTO;
	meetings: MeetingDTO[];
	projects?: ProjectDTO[];
	isMeetingsFetching: boolean;
	isMeetingsRefetching: boolean;
	isMeetingsLoading: boolean;
};

export function MeetingSelect({
	handleMeetingSelect,
	selectedMeeting,
	meetings,
	isMeetingsFetching,
	isMeetingsRefetching,
	isMeetingsLoading,
	projects,
}: MeetingSelectProps) {
	const reportMutation = useMeetingReport();
	const isPrinting = reportMutation.isPending;

	if ((isMeetingsFetching || isMeetingsLoading) && !isMeetingsRefetching) {
		return (
			<div className="w-full flex items-center justify-center py-10">
				<p>Carregando...</p>
			</div>
		);
	}

	return (
		<div className="flex flex-col w-full rounded-xl border bg-background shadow-sm p-4">
			<div className="pb-4 border-b">
				<p className="text-xl font-semibold">Reuniões</p>
			</div>

			<div className="flex flex-col gap-4 pt-4">
				{meetings.map((meeting) => (
					<MeetingCard
						key={meeting.id}
						meeting={meeting}
						selectedMeeting={selectedMeeting}
						handleMeetingSelect={handleMeetingSelect}
						handlePrint={() => reportMutation.mutate(meeting)}
						isPrinting={isPrinting}
						projects={projects}
					/>
				))}
			</div>
		</div>
	);
}
