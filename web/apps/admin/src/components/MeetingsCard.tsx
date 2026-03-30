import { Link } from '@tanstack/react-router';
import {
	Button,
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from '@voting/shared';
import { CalendarDays } from 'lucide-react';
import type { MeetingDTO } from '../hooks/useTodayMeetings';
import { MeetingList } from './MeetingList';

type MeetingsCardProps = {
	meetings: MeetingDTO[];
	isLoading: boolean;
};

export function MeetingsCard({ meetings, isLoading }: MeetingsCardProps) {
	const now = new Date().toLocaleDateString('pt-BR');

	return (
		<Card className="w-full min-h-24 h-fit">
			<CardHeader>
				<CardTitle className="flex items-center gap-2">
					<CalendarDays className="h-8 w-8 border-primary text-primary" />
					Reuniões de hoje, {now}
				</CardTitle>
				<CardDescription>
					{isLoading
						? 'Carregando reuniões...'
						: meetings.length === 0
							? 'Nenhuma reunião agendada para hoje'
							: 'Existem reuniões agendadas para hoje'}
				</CardDescription>
			</CardHeader>
			{!isLoading && meetings.length > 0 && (
				<>
					<CardContent>
						<MeetingList meetings={meetings} />
					</CardContent>
					<CardFooter className="flex gap-2 justify-end items-end h-fit">
						<CardDescription>
							<Button asChild className="justify-center" variant="outline">
								<Link to="/meetings">Acessar reuniões</Link>
							</Button>
						</CardDescription>
					</CardFooter>
				</>
			)}
		</Card>
	);
}
