import type { MeetingDTO } from '../hooks/useTodayMeetings';

type MeetingListProps = {
	meetings: MeetingDTO[];
};

export function MeetingList({ meetings }: MeetingListProps) {
	return (
		<ul className="my-4 ml-6 list-disc [&>li]:mt-0.5 text-sm text-muted-foreground">
			{meetings.map((meeting) => (
				<li key={meeting.rec_id}>
					{meeting.con_desc} - {meeting.rec_numero}
				</li>
			))}
		</ul>
	);
}
