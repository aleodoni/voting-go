import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from '@voting/shared';
import { Vote } from 'lucide-react';
import { VotingStats } from '@/hooks/useVotingStats';

interface VotingStatsTotalProps {
	stats?: VotingStats;
	isLoading: boolean;
}
export function VotingStatsTotalCard({
	stats,
	isLoading,
}: VotingStatsTotalProps) {
	return (
		<Card className="w-full">
			<CardHeader>
				<CardTitle className="flex items-center gap-2">
					<Vote className="h-8 w-8 border-primary text-primary" />
					Votações de hoje
				</CardTitle>
				<CardDescription>Total de projetos a serem votados</CardDescription>
			</CardHeader>
			<CardContent className="flex items-center justify-center mt-12">
				{isLoading ? (
					<p>Carregando...</p>
				) : (
					<p className="text-6xl font-bold text-primary">
						{stats?.total_projects}
					</p>
				)}
			</CardContent>
		</Card>
	);
}
