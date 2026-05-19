import { Button, Card, CardContent, CardFooter } from '@voting/shared';
import { cn } from '@voting/shared/src/lib/utils';
import { Calendar, Info, Printer } from 'lucide-react';
import { MeetingDTO } from '@/hooks/useTodayMeetings';
import { ProjectDTO } from '@/types/meeting';

type MeetingCardProps = {
	meeting: MeetingDTO;
	selectedMeeting?: MeetingDTO;
	handleMeetingSelect: (meeting: MeetingDTO) => void;
	handlePrint: () => void;
	isPrinting?: boolean;
	projects?: ProjectDTO[];
};

export function MeetingCard({
	meeting,
	selectedMeeting,
	handleMeetingSelect,
	handlePrint,
	isPrinting,
	projects,
}: MeetingCardProps) {
	const isSelected = selectedMeeting && meeting.id === selectedMeeting.id;

	function extraClass() {
		if (isSelected)
			return 'border-primary bg-blue-50 dark:bg-blue-900 shadow-md';
		return '';
	}

	return (
		<Card
			onClick={() => handleMeetingSelect(meeting)}
			className={cn(
				'w-full min-h-24 h-fit py-4 cursor-pointer transition-all hover:shadow-md',
				extraClass(),
			)}
		>
			<CardContent>
				<div className="flex pb-2">
					<p className="text-sm font-bold">{meeting.con_desc}</p>
				</div>
				<div className="flex flex-col gap-2">
					<div className="flex items-center gap-2">
						<Calendar strokeWidth={1} size={20} />
						<p className="text-sm">{meeting.rec_numero}</p>
					</div>
					<div className="flex items-center gap-2">
						<Info strokeWidth={1} size={20} />
						<p className="text-sm">{meeting.rec_tipo_reuniao}</p>
					</div>
				</div>
			</CardContent>

			{isSelected && projects && projects.length > 0 && (
				<CardFooter>
					<Button
						className="w-full"
						variant="outline"
						onClick={(e) => {
							e.stopPropagation();
							handlePrint();
						}}
						disabled={isPrinting}
					>
						<Printer size={16} strokeWidth={1} />
						{isPrinting ? 'Gerando...' : 'Imprimir votações'}
					</Button>
				</CardFooter>
			)}
		</Card>
	);
}
